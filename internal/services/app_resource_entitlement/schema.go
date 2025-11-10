// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_resource_entitlement

import (
	"context"

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
			"name": schema.StringAttribute{
				Description: "The name of the entitlement.",
				Optional:    true,
			},
			"provisioning_method": schema.StringAttribute{
				Description: "The provisioning method for the entitlement.",
				Optional:    true,
			},
			"requests_enabled": schema.BoolAttribute{
				Description: "Whether requests are enabled for the entitlement.",
				Optional:    true,
			},
			"linked_entitlement_ids": schema.ListAttribute{
				Description: "The IDs of entitlements that must be provisioned before this entitlement can be provisioned (optional).",
				Optional:    true,
				ElementType: types.StringType,
			},
			"manual_provisioning_assignees": schema.ListNestedAttribute{
				Description: `The manual provisioning assignees (users and groups) for this entitlement (optional, only used when provisioning_method is "manual").`,
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
	}
}

func (r *AppResourceEntitlementResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *AppResourceEntitlementResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
