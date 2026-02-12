// Package export provides a public API for generating Terraform state files and
// HCL configuration directly from raw Serval API JSON responses.
//
// This package is intended to be imported by external services (e.g., svmeta)
// that need to generate Terraform configuration without going through the
// OpenTofu import pipeline.
package export

import (
	"encoding/json"

	"github.com/ServalHQ/terraform-provider-serval/internal/directgen"
)

// Resource represents a single resource to generate state and HCL for.
type Resource struct {
	Type    string          // Terraform resource type (e.g., "serval_user")
	Name    string          // Terraform resource name (e.g., "alice_smith")
	RawJSON json.RawMessage // Raw API JSON response (the full {"data": {...}} envelope)
}

// Result contains the generated state file and HCL content.
type Result struct {
	StateJSON json.RawMessage // Complete terraform.tfstate content
	HCL       string          // Generated HCL resource blocks
}

// Options configures the generation process.
type Options struct {
	// ProviderAddress is the full provider address for the state file.
	// e.g., "registry.opentofu.org/servalhq/serval"
	ProviderAddress string

	// TerraformVersion is the version string for the state file.
	TerraformVersion string

	// Lineage is the unique identifier for the state lineage.
	Lineage string
}

// Generate produces a Terraform state file and HCL configuration from raw API data.
// It uses the provider's model structs to unmarshal API JSON into the correct types,
// then serializes them to state JSON and HCL.
//
// Each resource's RawJSON should be the complete API response envelope (e.g., {"data": {...}}).
func Generate(resources []Resource, opts Options) (*Result, error) {
	// Convert public types to internal types
	internalResources := make([]directgen.Resource, len(resources))
	for i, r := range resources {
		internalResources[i] = directgen.Resource{
			Type:    r.Type,
			Name:    r.Name,
			RawJSON: r.RawJSON,
		}
	}

	internalOpts := directgen.GenerateOptions{
		ProviderAddress:  opts.ProviderAddress,
		TerraformVersion: opts.TerraformVersion,
		Lineage:          opts.Lineage,
	}

	result, err := directgen.Generate(internalResources, internalOpts)
	if err != nil {
		return nil, err
	}

	return &Result{
		StateJSON: result.StateJSON,
		HCL:       result.HCL,
	}, nil
}

// GenerateForSingleResource is a convenience method that generates state attributes
// and HCL for a single resource. Useful for testing and incremental generation.
func GenerateForSingleResource(resourceType, resourceName string, rawJSON json.RawMessage) (stateAttrs json.RawMessage, hcl string, err error) {
	return directgen.GenerateForSingleResource(resourceType, resourceName, rawJSON, directgen.DefaultOptions())
}

// Resource type constants for convenience.
const (
	ResourceTypeUser                          = directgen.ResourceTypeUser
	ResourceTypeGroup                         = directgen.ResourceTypeGroup
	ResourceTypeTeam                          = directgen.ResourceTypeTeam
	ResourceTypeWorkflow                      = directgen.ResourceTypeWorkflow
	ResourceTypeWorkflowApprovalProcedure     = directgen.ResourceTypeWorkflowApprovalProcedure
	ResourceTypeGuidance                      = directgen.ResourceTypeGuidance
	ResourceTypeAccessPolicy                  = directgen.ResourceTypeAccessPolicy
	ResourceTypeAccessPolicyApprovalProcedure = directgen.ResourceTypeAccessPolicyApprovalProcedure
	ResourceTypeAppInstance                   = directgen.ResourceTypeAppInstance
	ResourceTypeAppResource                   = directgen.ResourceTypeAppResource
	ResourceTypeAppResourceRole               = directgen.ResourceTypeAppResourceRole
	ResourceTypeCustomService                 = directgen.ResourceTypeCustomService
)

// SupportedResourceTypes returns all resource types that the generator supports.
func SupportedResourceTypes() []string {
	return directgen.SupportedResourceTypes()
}
