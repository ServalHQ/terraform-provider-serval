// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy_approval_procedure

import (
	"context"

	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
							Description: "Whether the step can be approved by the requester themselves.",
							Computed:    true,
						},
						"serval_group_ids": schema.ListAttribute{
							Description: "The IDs of the Serval groups that can approve the step.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"specific_user_ids": schema.ListAttribute{
							Description: "The IDs of the specific users that can approve the step.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"step_type": schema.StringAttribute{
							Description: "The type of approval step.\nAvailable values: \"APPROVAL_PROCEDURE_STEP_TYPE_UNSPECIFIED\", \"SPECIFIC_USERS\", \"SERVAL_GROUPS\".",
							Computed:    true,
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

func (d *AccessPolicyApprovalProcedureDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *AccessPolicyApprovalProcedureDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
