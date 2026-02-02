// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy_approval_procedure

import (
	"context"

	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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
				Description:   "The approval steps for the procedure.",
				Optional:      true,
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace(), listplanmodifier.UseStateForUnknown()},
				Computed:      true,
				CustomType:    customfield.NewNestedObjectListType[AccessPolicyApprovalProcedureStepsModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
							Description:   "The ID of the approval step.",
							Computed:      true,
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"allow_self_approval": schema.BoolAttribute{
							Description: "Whether the step can be approved by the requester themselves.",
							Computed:    true,
							Optional:    true,
						},
						"custom_workflow_id": schema.StringAttribute{
							Description: "A workflow ID to execute to determine the approvers for this step (or to auto-approve the step).",
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
