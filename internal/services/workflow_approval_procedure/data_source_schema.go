// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflow_approval_procedure

import (
	"context"

	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*WorkflowApprovalProcedureDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The ID of the approval procedure.",
				Required:    true,
			},
			"workflow_id": schema.StringAttribute{
				Description: "The ID of the workflow.",
				Required:    true,
			},
			"steps": schema.ListNestedAttribute{
				Description: "The steps in the approval procedure.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[WorkflowApprovalProcedureStepsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The ID of the approval step.",
							Computed:    true,
						},
						"allow_self_approval": schema.BoolAttribute{
							Description: "Whether the step can be approved by the requester themselves.\n optional so server can distinguish \"not set\" from \"explicitly false\"\n (DB defaults to TRUE; proto3 defaults bool to false)",
							Computed:    true,
						},
						"approvers": schema.ListNestedAttribute{
							Description: "Exactly one of approvers or custom_workflow must be set.\n Mutual exclusivity validated server-side.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[WorkflowApprovalProcedureStepsApproversDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"app_owner": schema.StringAttribute{
										Description: "App owners as approvers. Only valid for access policy approval procedures.",
										Computed:    true,
										CustomType:  jsontypes.NormalizedType{},
									},
									"notify": schema.BoolAttribute{
										Description: "Whether to notify this approver when the step is pending.",
										Computed:    true,
									},
									"group": schema.SingleNestedAttribute{
										Description: "A Serval group as approvers.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[WorkflowApprovalProcedureStepsApproversGroupDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"group_id": schema.StringAttribute{
												Description: "The ID of the Serval group.",
												Computed:    true,
											},
										},
									},
									"manager": schema.StringAttribute{
										Description: "The requester's manager as an approver.",
										Computed:    true,
										CustomType:  jsontypes.NormalizedType{},
									},
									"user": schema.SingleNestedAttribute{
										Description: "A specific user as an approver.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[WorkflowApprovalProcedureStepsApproversUserDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"user_id": schema.StringAttribute{
												Description: "The ID of the user.",
												Computed:    true,
											},
										},
									},
								},
							},
						},
						"custom_workflow": schema.SingleNestedAttribute{
							Description: "Configuration for a custom workflow that determines approvers or auto-approves.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[WorkflowApprovalProcedureStepsCustomWorkflowDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"workflow_id": schema.StringAttribute{
									Description: "The ID of the workflow to execute.",
									Computed:    true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *WorkflowApprovalProcedureDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *WorkflowApprovalProcedureDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
