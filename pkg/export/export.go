// Package export provides a public API for generating Terraform state files and
// HCL configuration directly from raw Serval API JSON responses.
//
// This package re-exports types and functions from internal/directgen using type
// aliases, following the same pattern as svflow/pkg/serviceconfig/export.go.
// External services (e.g., svmeta) can import this package to use the provider's
// model structs and reflection-based codecs without accessing internal packages.
package export

import (
	"encoding/json"

	"github.com/ServalHQ/terraform-provider-serval/internal/directgen"
)

// Resource is re-exported from internal/directgen.
type Resource = directgen.Resource

// GenerateResult is re-exported from internal/directgen.
type GenerateResult = directgen.GenerateResult

// GenerateOptions is re-exported from internal/directgen.
type GenerateOptions = directgen.GenerateOptions

// Generate is re-exported from internal/directgen.
var Generate = directgen.Generate

// GenerateForSingleResource is a convenience method that generates state attributes
// and HCL for a single resource using default options.
func GenerateForSingleResource(resourceType, resourceName string, rawJSON json.RawMessage) (stateAttrs json.RawMessage, hcl string, err error) {
	return directgen.GenerateForSingleResource(resourceType, resourceName, rawJSON, directgen.DefaultOptions())
}

// DefaultOptions is re-exported from internal/directgen.
var DefaultOptions = directgen.DefaultOptions

// SupportedResourceTypes is re-exported from internal/directgen.
var SupportedResourceTypes = directgen.SupportedResourceTypes

// Resource type constants re-exported from internal/directgen.
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
