// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_resource_entitlement

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*AppResourceEntitlementResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The ID of the entitlement.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"resource_id": schema.StringAttribute{
				Description:   "The ID of the resource.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"access_policy_id": schema.StringAttribute{
				Description: "The default access policy for the entitlement (optional).",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "A description of the entitlement.",
				Optional:    true,
			},
			"external_data": schema.StringAttribute{
				Description: "Data from the external system as a JSON string (optional).",
				Optional:    true,
			},
			"external_id": schema.StringAttribute{
				Description: "The external ID of the entitlement in the external system (optional).",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the entitlement.",
				Optional:    true,
			},
			"requests_enabled": schema.BoolAttribute{
				Description: "Whether requests are enabled for the entitlement.",
				Optional:    true,
			},
			"provisioning_method": schema.SingleNestedAttribute{
				Description: "Provisioning configuration. Exactly one method should be set.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"builtin_workflow": schema.StringAttribute{
						Description: "**Option: builtin_workflow**",
						Optional:    true,
						CustomType:  jsontypes.NormalizedType{},
					},
					"custom_workflow": schema.SingleNestedAttribute{
						Description: "**Option: custom_workflow**",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"deprovision_workflow_id": schema.StringAttribute{
								Description: "The workflow ID to deprovision access.",
								Optional:    true,
							},
							"provision_workflow_id": schema.StringAttribute{
								Description: "The workflow ID to provision access.",
								Optional:    true,
							},
						},
					},
					"linked_entitlements": schema.SingleNestedAttribute{
						Description: "**Option: linked_entitlements**",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"linked_entitlement_ids": schema.ListAttribute{
								Description: "The IDs of prerequisite entitlements.",
								Optional:    true,
								ElementType: types.StringType,
							},
						},
					},
					"manual": schema.SingleNestedAttribute{
						Description: "**Option: manual**",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"assignees": schema.ListNestedAttribute{
								Description: "Users and groups that should be assigned/notified for manual provisioning.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"assignee_id": schema.StringAttribute{
											Description: "The ID of the user or group.",
											Optional:    true,
										},
										"assignee_type": schema.StringAttribute{
											Description: "The type of assignee.\nAvailable values: \"MANUAL_PROVISIONING_ASSIGNEE_TYPE_UNSPECIFIED\", \"MANUAL_PROVISIONING_ASSIGNEE_TYPE_USER\", \"MANUAL_PROVISIONING_ASSIGNEE_TYPE_GROUP\".",
											Optional:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"MANUAL_PROVISIONING_ASSIGNEE_TYPE_UNSPECIFIED",
													"MANUAL_PROVISIONING_ASSIGNEE_TYPE_USER",
													"MANUAL_PROVISIONING_ASSIGNEE_TYPE_GROUP",
												),
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *AppResourceEntitlementResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *AppResourceEntitlementResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
