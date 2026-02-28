// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy_approval_procedure

import (
	"context"

	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*AccessPolicyApprovalProcedureDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"access_policy_id": schema.StringAttribute{
				Description: "The ID of the access policy.",
				Required:    true,
			},
			"id": schema.StringAttribute{
				Description: "The ID of the approval procedure.",
				Required:    true,
			},
			"steps": schema.ListNestedAttribute{
				Description: "The steps in the approval procedure.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[AccessPolicyApprovalProcedureStepsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The ID of the approval step.",
							Computed:    true,
						},
					"allow_self_approval": schema.BoolAttribute{
						Description: "Whether the step can be approved by the requester themselves. Defaults to true if not set.",
							Computed:    true,
						},
					"approvers": schema.ListNestedAttribute{
						Description: "The list of approvers for this step. Exactly one of `approvers` or `custom_workflow` must be set.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[AccessPolicyApprovalProcedureStepsApproversDataSourceModel](ctx),
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
										CustomType:  customfield.NewNestedObjectType[AccessPolicyApprovalProcedureStepsApproversGroupDataSourceModel](ctx),
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
										CustomType:  customfield.NewNestedObjectType[AccessPolicyApprovalProcedureStepsApproversUserDataSourceModel](ctx),
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
							CustomType:  customfield.NewNestedObjectType[AccessPolicyApprovalProcedureStepsCustomWorkflowDataSourceModel](ctx),
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

func (d *AccessPolicyApprovalProcedureDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *AccessPolicyApprovalProcedureDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
