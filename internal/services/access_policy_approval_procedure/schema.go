// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy_approval_procedure

import (
	"context"

	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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
							Description: "Whether the step can be approved by the requester themselves.\n optional so server can distinguish \"not set\" from \"explicitly false\"\n (DB defaults to TRUE; proto3 defaults bool to false)",
							Computed:    true,
							Optional:    true,
						},
						"approvers": schema.ListNestedAttribute{
							Description: "Exactly one of approvers or custom_workflow must be set.\n Mutual exclusivity validated server-side.",
							Optional:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"app_owner": schema.StringAttribute{
										Description: "App owners as approvers. Only valid for access policy approval procedures.",
										Optional:    true,
										CustomType:  jsontypes.NormalizedType{},
									},
									"notify": schema.BoolAttribute{
										Description: "Whether to notify this approver when the step is pending.",
										Computed:    true,
										Optional:    true,
									},
									"group": schema.SingleNestedAttribute{
										Description: "A Serval group as approvers.",
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"group_id": schema.StringAttribute{
												Description: "The ID of the Serval group.",
												Optional:    true,
											},
										},
									},
									"manager": schema.StringAttribute{
										Description: "The requester's manager as an approver.",
										Optional:    true,
										CustomType:  jsontypes.NormalizedType{},
									},
									"user": schema.SingleNestedAttribute{
										Description: "A specific user as an approver.",
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"user_id": schema.StringAttribute{
												Description: "The ID of the user.",
												Optional:    true,
											},
										},
									},
								},
							},
						},
						"custom_workflow": schema.SingleNestedAttribute{
							Description: "Configuration for a custom workflow that determines approvers or auto-approves.",
							Optional:    true,
							Attributes: map[string]schema.Attribute{
								"workflow_id": schema.StringAttribute{
									Description: "The ID of the workflow to execute.",
									Optional:    true,
								},
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
