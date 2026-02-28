// Package directgen provides direct generation of Terraform state files and HCL
// configuration from raw API JSON, bypassing OpenTofu's import pipeline.
package directgen

import (
	"context"

	"github.com/ServalHQ/terraform-provider-serval/internal/services/access_policy"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/access_policy_approval_procedure"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/app_instance"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/app_resource"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/app_resource_role"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/custom_service"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/group"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/guidance"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/team"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/user"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/workflow"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/workflow_approval_procedure"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Resource type constants matching the Serval Terraform provider.
const (
	ResourceTypeUser                          = "serval_user"
	ResourceTypeGroup                         = "serval_group"
	ResourceTypeTeam                          = "serval_team"
	ResourceTypeWorkflow                      = "serval_workflow"
	ResourceTypeWorkflowApprovalProcedure     = "serval_workflow_approval_procedure"
	ResourceTypeGuidance                      = "serval_guidance"
	ResourceTypeAccessPolicy                  = "serval_access_policy"
	ResourceTypeAccessPolicyApprovalProcedure = "serval_access_policy_approval_procedure"
	ResourceTypeAppInstance                   = "serval_app_instance"
	ResourceTypeAppResource                   = "serval_app_resource"
	ResourceTypeAppResourceRole               = "serval_app_resource_role"
	ResourceTypeCustomService                 = "serval_custom_service"
)

// ModelConstructor returns a new zero-value data envelope for unmarshaling.
type ModelConstructor func() any

// ModelExtractor extracts the inner model from a data envelope.
type ModelExtractor func(envelope any) any

// SchemaVersion returns the Terraform schema version for the resource type.
// Currently all Serval resources use schema version 0.
const defaultSchemaVersion = 0

// ResourceRegistryEntry defines how to construct and extract a model for a resource type.
type ResourceRegistryEntry struct {
	Constructor   ModelConstructor
	Extractor     ModelExtractor
	SchemaVersion uint64
	SchemaFunc    func(context.Context) schema.Schema
}

// resourceRegistry maps Terraform resource types to their model constructors and extractors.
var resourceRegistry = map[string]ResourceRegistryEntry{
	ResourceTypeUser: {
		Constructor:   func() any { return &user.UserDataEnvelope{} },
		Extractor:     func(e any) any { return e.(*user.UserDataEnvelope).Data },
		SchemaVersion: defaultSchemaVersion,
		SchemaFunc:    user.ResourceSchema,
	},
	ResourceTypeGroup: {
		Constructor: func() any { return &group.GroupDataEnvelope{} },
		Extractor: func(e any) any {
			data := e.(*group.GroupDataEnvelope).Data
			data.NormalizeState()
			return data
		},
		SchemaVersion: defaultSchemaVersion,
		SchemaFunc:    group.ResourceSchema,
	},
	ResourceTypeTeam: {
		Constructor:   func() any { return &team.TeamDataEnvelope{} },
		Extractor:     func(e any) any { return e.(*team.TeamDataEnvelope).Data },
		SchemaVersion: defaultSchemaVersion,
		SchemaFunc:    team.ResourceSchema,
	},
	ResourceTypeWorkflow: {
		Constructor:   func() any { return &workflow.WorkflowDataEnvelope{} },
		Extractor:     func(e any) any { return e.(*workflow.WorkflowDataEnvelope).Data },
		SchemaVersion: defaultSchemaVersion,
		SchemaFunc:    workflow.ResourceSchema,
	},
	ResourceTypeWorkflowApprovalProcedure: {
		Constructor:   func() any { return &workflow_approval_procedure.WorkflowApprovalProcedureDataEnvelope{} },
		Extractor:     func(e any) any { return e.(*workflow_approval_procedure.WorkflowApprovalProcedureDataEnvelope).Data },
		SchemaVersion: defaultSchemaVersion,
		SchemaFunc:    workflow_approval_procedure.ResourceSchema,
	},
	ResourceTypeGuidance: {
		Constructor:   func() any { return &guidance.GuidanceDataEnvelope{} },
		Extractor:     func(e any) any { return e.(*guidance.GuidanceDataEnvelope).Data },
		SchemaVersion: defaultSchemaVersion,
		SchemaFunc:    guidance.ResourceSchema,
	},
	ResourceTypeAccessPolicy: {
		Constructor:   func() any { return &access_policy.AccessPolicyDataEnvelope{} },
		Extractor:     func(e any) any { return e.(*access_policy.AccessPolicyDataEnvelope).Data },
		SchemaVersion: defaultSchemaVersion,
		SchemaFunc:    access_policy.ResourceSchema,
	},
	ResourceTypeAccessPolicyApprovalProcedure: {
		Constructor: func() any { return &access_policy_approval_procedure.AccessPolicyApprovalProcedureDataEnvelope{} },
		Extractor: func(e any) any {
			return e.(*access_policy_approval_procedure.AccessPolicyApprovalProcedureDataEnvelope).Data
		},
		SchemaVersion: defaultSchemaVersion,
		SchemaFunc:    access_policy_approval_procedure.ResourceSchema,
	},
	ResourceTypeAppInstance: {
		Constructor:   func() any { return &app_instance.AppInstanceDataEnvelope{} },
		Extractor:     func(e any) any { return e.(*app_instance.AppInstanceDataEnvelope).Data },
		SchemaVersion: defaultSchemaVersion,
		SchemaFunc:    app_instance.ResourceSchema,
	},
	ResourceTypeAppResource: {
		Constructor:   func() any { return &app_resource.AppResourceDataEnvelope{} },
		Extractor:     func(e any) any { return e.(*app_resource.AppResourceDataEnvelope).Data },
		SchemaVersion: defaultSchemaVersion,
		SchemaFunc:    app_resource.ResourceSchema,
	},
	ResourceTypeAppResourceRole: {
		Constructor:   func() any { return &app_resource_role.AppResourceRoleDataEnvelope{} },
		Extractor:     func(e any) any { return e.(*app_resource_role.AppResourceRoleDataEnvelope).Data },
		SchemaVersion: defaultSchemaVersion,
		SchemaFunc:    app_resource_role.ResourceSchema,
	},
	ResourceTypeCustomService: {
		Constructor:   func() any { return &custom_service.CustomServiceDataEnvelope{} },
		Extractor:     func(e any) any { return e.(*custom_service.CustomServiceDataEnvelope).Data },
		SchemaVersion: defaultSchemaVersion,
		SchemaFunc:    custom_service.ResourceSchema,
	},
}

// GetRegistryEntry returns the registry entry for a resource type.
func GetRegistryEntry(resourceType string) (ResourceRegistryEntry, bool) {
	entry, ok := resourceRegistry[resourceType]
	return entry, ok
}

// SupportedResourceTypes returns all registered resource types.
func SupportedResourceTypes() []string {
	types := make([]string, 0, len(resourceRegistry))
	for k := range resourceRegistry {
		types = append(types, k)
	}
	return types
}
