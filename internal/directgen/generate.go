package directgen

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/ServalHQ/terraform-provider-serval/internal/apijson"
	"github.com/ServalHQ/terraform-provider-serval/internal/hclcodec"
	"github.com/ServalHQ/terraform-provider-serval/internal/statecodec"
)

// Resource represents a single resource to generate state and HCL for.
type Resource struct {
	Type    string          // Terraform resource type (e.g., "serval_user")
	Name    string          // Terraform resource name (e.g., "alice_smith")
	RawJSON json.RawMessage // Raw API JSON response (the full {"data": {...}} envelope)
}

// GenerateResult contains the generated state file and HCL content.
type GenerateResult struct {
	StateJSON json.RawMessage // Complete terraform.tfstate content
	HCL       string          // Generated HCL resource blocks
}

// stateFile is the top-level structure of a Terraform state file (v4).
type stateFile struct {
	Version          int             `json:"version"`
	TerraformVersion string          `json:"terraform_version"`
	Serial           uint64          `json:"serial"`
	Lineage          string          `json:"lineage"`
	Outputs          map[string]any  `json:"outputs"`
	Resources        []stateResource `json:"resources"`
}

// stateResource represents a resource in the state file.
type stateResource struct {
	Module    string          `json:"module,omitempty"`
	Mode      string          `json:"mode"`
	Type      string          `json:"type"`
	Name      string          `json:"name"`
	Provider  string          `json:"provider"`
	Instances []stateInstance `json:"instances"`
}

// stateInstance represents a single instance of a resource.
type stateInstance struct {
	SchemaVersion       uint64          `json:"schema_version"`
	Attributes          json.RawMessage `json:"attributes"`
	SensitiveAttributes json.RawMessage `json:"sensitive_attributes"`
}

// GenerateOptions configures the generation process.
type GenerateOptions struct {
	// ProviderAddress is the full provider address for the state file.
	// e.g., "registry.opentofu.org/servalhq/serval"
	ProviderAddress string

	// TerraformVersion is the version string for the state file.
	TerraformVersion string

	// Lineage is the unique identifier for the state lineage.
	// If empty, a placeholder will be used.
	Lineage string
}

// DefaultOptions returns sensible defaults for generation.
func DefaultOptions() GenerateOptions {
	return GenerateOptions{
		ProviderAddress:  "registry.opentofu.org/servalhq/serval",
		TerraformVersion: "1.9.0",
		Lineage:          "direct-gen",
	}
}

// Generate produces a Terraform state file and HCL configuration from raw API data.
// It uses the provider's model structs (via the model registry) to unmarshal API JSON
// into the correct types, then serializes them to state JSON and HCL.
func Generate(resources []Resource, opts GenerateOptions) (*GenerateResult, error) {
	if opts.ProviderAddress == "" {
		opts = DefaultOptions()
	}

	providerRef := fmt.Sprintf(`provider[%q]`, opts.ProviderAddress)

	var stateResources []stateResource
	var hclBlocks []string

	// Sort resources by type then name for deterministic output
	sorted := make([]Resource, len(resources))
	copy(sorted, resources)
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Type != sorted[j].Type {
			return sorted[i].Type < sorted[j].Type
		}
		return sorted[i].Name < sorted[j].Name
	})

	for _, res := range sorted {
		entry, ok := GetRegistryEntry(res.Type)
		if !ok {
			return nil, fmt.Errorf("directgen: unsupported resource type %q", res.Type)
		}

		// Step 1: Unmarshal raw API JSON into the provider model struct
		envelope := entry.Constructor()
		if err := apijson.UnmarshalRoot(res.RawJSON, envelope); err != nil {
			return nil, fmt.Errorf("directgen: unmarshal %s.%s: %w", res.Type, res.Name, err)
		}

		// Step 2: Extract the inner model from the envelope
		model := entry.Extractor(envelope)

		// Step 3: Serialize to state attributes JSON
		attrs, err := statecodec.SerializeToStateAttributes(model)
		if err != nil {
			return nil, fmt.Errorf("directgen: serialize state %s.%s: %w", res.Type, res.Name, err)
		}

		// Step 4: Generate HCL resource block
		hcl, err := hclcodec.GenerateResourceBlock(res.Type, res.Name, model)
		if err != nil {
			return nil, fmt.Errorf("directgen: generate HCL %s.%s: %w", res.Type, res.Name, err)
		}

		// Step 5: Build state resource entry
		stateRes := stateResource{
			Mode:     "managed",
			Type:     res.Type,
			Name:     res.Name,
			Provider: providerRef,
			Instances: []stateInstance{
				{
					SchemaVersion:       entry.SchemaVersion,
					Attributes:          attrs,
					SensitiveAttributes: json.RawMessage("[]"),
				},
			},
		}

		stateResources = append(stateResources, stateRes)
		hclBlocks = append(hclBlocks, hcl)
	}

	// Assemble state file
	state := stateFile{
		Version:          4,
		TerraformVersion: opts.TerraformVersion,
		Serial:           1,
		Lineage:          opts.Lineage,
		Outputs:          map[string]any{},
		Resources:        stateResources,
	}

	stateJSON, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("directgen: marshal state: %w", err)
	}

	return &GenerateResult{
		StateJSON: stateJSON,
		HCL:       strings.Join(hclBlocks, "\n"),
	}, nil
}

// GenerateForSingleResource is a convenience method for generating state and HCL
// for a single resource. Useful for testing.
func GenerateForSingleResource(resourceType, resourceName string, rawJSON json.RawMessage, opts GenerateOptions) (stateAttrs json.RawMessage, hcl string, err error) {
	entry, ok := GetRegistryEntry(resourceType)
	if !ok {
		return nil, "", fmt.Errorf("directgen: unsupported resource type %q", resourceType)
	}

	envelope := entry.Constructor()
	if err := apijson.UnmarshalRoot(rawJSON, envelope); err != nil {
		return nil, "", fmt.Errorf("directgen: unmarshal: %w", err)
	}

	model := entry.Extractor(envelope)

	attrs, err := statecodec.SerializeToStateAttributes(model)
	if err != nil {
		return nil, "", fmt.Errorf("directgen: serialize state: %w", err)
	}

	hcl, err = hclcodec.GenerateResourceBlock(resourceType, resourceName, model)
	if err != nil {
		return nil, "", fmt.Errorf("directgen: generate HCL: %w", err)
	}

	return attrs, hcl, nil
}
