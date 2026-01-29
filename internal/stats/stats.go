// Package stats provides global statistics tracking for debugging.
// It tracks API requests and cache performance during Terraform operations.
package stats

import (
	"context"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Stats holds global counters for API and cache performance.
type Stats struct {
	APIRequests int64
	CacheHits   int64
	CacheMisses int64
	CacheLoads  int64

	// Endpoint call counts (protected by mutex)
	endpointMu    sync.Mutex
	endpointCalls map[string]int
}

// Global is the global stats instance.
var Global = &Stats{
	endpointCalls: make(map[string]int),
}

// RecordAPIRequest increments the API request counter and tracks the endpoint.
func (s *Stats) RecordAPIRequest(endpoint string) {
	atomic.AddInt64(&s.APIRequests, 1)

	s.endpointMu.Lock()
	s.endpointCalls[endpoint]++
	s.endpointMu.Unlock()
}

// GetEndpointCalls returns a copy of the endpoint call counts.
func (s *Stats) GetEndpointCalls() map[string]int {
	s.endpointMu.Lock()
	defer s.endpointMu.Unlock()

	result := make(map[string]int, len(s.endpointCalls))
	for k, v := range s.endpointCalls {
		result[k] = v
	}
	return result
}

// RecordCacheHit increments the cache hit counter.
func (s *Stats) RecordCacheHit() {
	atomic.AddInt64(&s.CacheHits, 1)
}

// RecordCacheMiss increments the cache miss counter.
func (s *Stats) RecordCacheMiss() {
	atomic.AddInt64(&s.CacheMisses, 1)
}

// RecordCacheLoad increments the cache load counter.
func (s *Stats) RecordCacheLoad() {
	atomic.AddInt64(&s.CacheLoads, 1)
}

// Snapshot returns current stats values (thread-safe).
func (s *Stats) Snapshot() (apiReqs, loads, hits, misses int64) {
	return atomic.LoadInt64(&s.APIRequests),
		atomic.LoadInt64(&s.CacheLoads),
		atomic.LoadInt64(&s.CacheHits),
		atomic.LoadInt64(&s.CacheMisses)
}

// HitRate returns the cache hit rate as a percentage.
func (s *Stats) HitRate() float64 {
	hits := atomic.LoadInt64(&s.CacheHits)
	misses := atomic.LoadInt64(&s.CacheMisses)
	total := hits + misses
	if total == 0 {
		return 0
	}
	return float64(hits) / float64(total) * 100
}

// logging state
var (
	logMu      sync.Mutex
	logStarted bool
)

// StartPeriodicLogging starts a background goroutine that logs aggregate stats
// via tflog every 5 seconds. Call this once during provider Configure.
func StartPeriodicLogging(ctx context.Context) {
	logMu.Lock()
	defer logMu.Unlock()

	if logStarted {
		return
	}
	logStarted = true

	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			apiReqs, loads, hits, misses := Global.Snapshot()
			hitRate := Global.HitRate()
			endpoints := Global.GetEndpointCalls()

			// Build endpoint summary (top 10 by call count)
			type endpointCount struct {
				endpoint string
				count    int
			}
			var sorted []endpointCount
			for ep, count := range endpoints {
				sorted = append(sorted, endpointCount{ep, count})
			}
			sort.Slice(sorted, func(i, j int) bool {
				return sorted[i].count > sorted[j].count
			})

			// Build map for logging (top 10)
			endpointMap := make(map[string]interface{})
			for i, ec := range sorted {
				if i >= 10 {
					break
				}
				endpointMap[ec.endpoint] = ec.count
			}

			tflog.Info(ctx, "[STATS] Provider performance",
				map[string]interface{}{
					"api_requests":   apiReqs,
					"cache_loads":    loads,
					"cache_hits":     hits,
					"cache_misses":   misses,
					"cache_hit_rate": hitRate,
				})

			if len(endpointMap) > 0 {
				tflog.Info(ctx, "[STATS] Endpoint calls (top 10)", endpointMap)
			}
		}
	}()
}
