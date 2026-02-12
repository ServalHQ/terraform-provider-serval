# Why Direct State & HCL Generation: OpenTofu Import Pipeline Limitations

## Background

The `configexport` service generates Terraform configuration for Serval workspaces by fetching all resources from APIs and producing `.tf` files and `terraform.tfstate`. The original pipeline delegated this to OpenTofu's `import` and `plan -generate-config-out` commands. At scale (100K+ resources), this pipeline becomes untenable. This document catalogs the specific bottlenecks discovered in the OpenTofu source code.

## Limitation 1: Global State Mutex (`SyncState`)

**File:** `opentofu/internal/states/sync.go`

OpenTofu wraps the entire in-memory state in a `SyncState` struct protected by a single `sync.RWMutex`. Every resource read or write acquires this lock:

```go
type SyncState struct {
    state *State
    lock  sync.RWMutex
}
```

During import, each resource calls `SetResourceInstanceCurrent()` which takes a **write lock**. Since all resources contend on this single mutex, state writes become serialized regardless of the parallelism setting. For 1M resources, this creates millions of lock acquisitions on a single contention point.

## Limitation 2: Default Parallelism of 10

**File:** `opentofu/internal/command/arguments/extended.go`

```go
const DefaultParallelism = 10
```

OpenTofu's DAG walker uses a semaphore (`parallelSem`) initialized to this value. Even if you override it via `-parallelism`, the global state mutex (Limitation 1) ensures the effective parallelism is much lower than the configured value. Increasing parallelism actually increases lock contention.

## Limitation 3: Intermediate State Persistence Every 20 Seconds

**File:** `opentofu/internal/backend/local/backend_apply.go`

```go
var defaultPersistInterval = 20 * time.Second
```

During `tofu apply`, a `StateHook` fires after every resource operation. Every 20 seconds, it serializes the **entire state** to JSON and writes it to disk. For a state file with 1M resources:

- JSON serialization of the full state takes seconds
- This blocks all other resource operations while it runs
- The `StatesMarshalEqual` function serializes the state **twice** for comparison before persisting
- At 1M resources, intermediate persists alone can consume more time than the actual import work

## Limitation 4: Full State DeepCopy on Every Read

**File:** `opentofu/internal/states/sync.go`

Every call to `SyncState.Module()` or `SyncState.ResourceInstanceObject()` performs a `DeepCopy()` of the returned data while holding the read lock. For complex resources with nested attributes, this creates enormous GC pressure. At 1M resources, the runtime spends a significant fraction of time in deep-copy allocation.

## Limitation 5: Per-Resource gRPC Round-Trips

**File:** `opentofu/internal/tofu/node_resource_plan_instance.go` (`managedResourceExecute`)

For each imported resource, OpenTofu makes **at minimum** these provider gRPC calls:

1. `ImportResourceState` -- asks the provider to look up the resource by ID
2. `ReadResource` -- re-fetches the full resource state
3. `ValidateResourceConfig` -- validates the generated config
4. `PlanResourceChange` -- computes the plan diff

Each call crosses a gRPC boundary (even for in-process providers, there's serialization overhead). For 1M resources, that's 4M+ gRPC round-trips minimum. The provider's `ImportResourceState` then makes an HTTP call back to the Serval API -- meaning the same data we already fetched is fetched again.

## Limitation 6: TransitiveReduction Complexity

**File:** `opentofu/internal/tofu/graph.go` (uses `dag.AcyclicGraph.TransitiveReduction`)

Before walking the resource graph, OpenTofu computes its transitive reduction with complexity **O(V * E)** where V is vertices and E is edges. For 1M resources with even modest inter-resource dependencies, this single graph operation can take minutes.

## Limitation 7: 2M+ Goroutine Overhead

**File:** `opentofu/internal/tofu/graph_walk_context.go` (`Execute`)

The DAG walker spawns a goroutine for every vertex in the graph. Each resource creates at least two vertices (the resource node + its state cleanup node). For 1M resources, that's 2M+ goroutines. While Go handles goroutines efficiently, the scheduling overhead, stack allocation (~8KB each = ~16GB), and the semaphore contention across all of them creates measurable drag.

## Limitation 8: Plan Changes Mutex

**File:** `opentofu/internal/tofu/context.go`

In addition to the state mutex, plan changes are protected by a separate `ChangesSync` with its own `sync.Mutex`. Every resource that completes planning must acquire this lock to record its change. This is a write-only mutex (not RWMutex), creating a second serialization bottleneck.

## Limitation 9: Config Generation Happens After Full Plan

**File:** `opentofu/internal/genconfig/generate_config_write.go`

When using `-generate-config-out`, HCL generation happens as a post-processing step. The `Change.MaybeWriteConfig` function is called after the plan completes, writing generated config to disk. This means OpenTofu must complete the entire plan (with all its overhead) before any config is written. There's no streaming or incremental output.

## Limitation 10: Redundant Data Flow

The fundamental architectural limitation: the `configexport` service already has all the resource data from the Serval APIs. OpenTofu's import pipeline then:

1. Receives import blocks pointing to resource IDs
2. Calls the provider's `ImportResourceState` gRPC
3. The provider calls the Serval API **again** to fetch the same data
4. Returns it through gRPC serialization
5. OpenTofu deserializes it into state
6. Validates it, plans it, deep-copies it, mutex-locks it, persists it

This round-trip through OpenTofu adds hours of overhead to re-fetch and re-process data that was already available in memory.

## The Solution: Direct Generation

Instead of routing through OpenTofu, we:

1. Capture the raw API JSON during the initial fetch (already happening)
2. Unmarshal it into the **same provider model structs** using `apijson.UnmarshalRoot` (the exact function the provider uses)
3. Serialize the model to state JSON via reflection on `tfsdk` tags (`statecodec`)
4. Generate HCL via reflection on `tfsdk` + `json` tags (`hclcodec`)
5. Assemble a valid v4 state file

This eliminates all 10 limitations above. For 1M resources, generation completes in seconds instead of hours.

### Correctness Guarantees

- **Structural:** The provider's model structs (with `tfsdk`/`json` tags) are the single source of truth, shared between the provider and the generator
- **Runtime:** We call the same `apijson.UnmarshalRoot` function as the provider's `Read` method
- **CI:** Golden-file differential tests compare direct output against known-good fixtures for all 11 resource types
- **Optional:** `tofu plan -refresh-only` can be run on the output to confirm zero drift
