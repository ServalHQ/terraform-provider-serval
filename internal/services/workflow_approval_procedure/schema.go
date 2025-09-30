// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflow_approval_procedure

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*WorkflowApprovalProcedureResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The ID of the workflow approval procedure.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"workflow_id": schema.StringAttribute{
				Description:   "The ID of the workflow to create the approval procedure for.",
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
						"custom_workflow_id": schema.StringAttribute{
							Description: "If step_type is CUSTOM_WORKFLOW, this is the workflow ID to execute",
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

func (r *WorkflowApprovalProcedureResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *WorkflowApprovalProcedureResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
