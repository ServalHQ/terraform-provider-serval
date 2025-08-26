// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy_approval_procedure

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

var _ resource.ResourceWithConfigValidators = (*AccessPolicyApprovalProcedureResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The ID of the access policy approval procedure.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"access_policy_id": schema.StringAttribute{
				Description:   "The ID of the access policy to create the approval procedure for.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"steps": schema.ListNestedAttribute{
				Description: "The approval steps for the procedure.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The ID of the approval step.",
							Optional:    true,
						},
						"allow_self_approval": schema.BoolAttribute{
							Description: "Whether the step can be approved by the requester themselves.",
							Optional:    true,
						},
						"serval_group_ids": schema.ListAttribute{
							Description: "The IDs of the Serval groups that can approve the step.",
							Optional:    true,
							ElementType: types.StringType,
						},
						"specific_user_ids": schema.ListAttribute{
							Description: "The IDs of the specific users that can approve the step.",
							Optional:    true,
							ElementType: types.StringType,
						},
						"step_type": schema.StringAttribute{
							Description: "The type of approval step.\nAvailable values: \"APPROVAL_PROCEDURE_STEP_TYPE_UNSPECIFIED\", \"SPECIFIC_USERS\", \"SERVAL_GROUPS\".",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"APPROVAL_PROCEDURE_STEP_TYPE_UNSPECIFIED",
									"SPECIFIC_USERS",
									"SERVAL_GROUPS",
								),
							},
						},
					},
				},
			},
		},
	}
}

func (r *AccessPolicyApprovalProcedureResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *AccessPolicyApprovalProcedureResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
