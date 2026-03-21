// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package approval_delegation

import (
	"context"

	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ApprovalDelegationDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The ID of the approval delegation rule.",
				Computed:    true,
				Optional:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "The timestamp when the delegation was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"delegator_user_id": schema.StringAttribute{
				Description: "The ID of the user who is delegating their approvals.",
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Description: "A description of the delegation rule.",
				Computed:    true,
			},
			"priority": schema.Int64Attribute{
				Description: "The priority of the delegation rule (lower values are higher priority).",
				Computed:    true,
			},
			"delegates": schema.ListNestedAttribute{
				Description: "The delegates who can approve on behalf of the delegator.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ApprovalDelegationDelegatesDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The ID of the delegate (user ID or group ID, depending on type).",
							Computed:    true,
						},
						"type": schema.StringAttribute{
							Description: "The type of delegate (user or group).\nAvailable values: \"APPROVAL_DELEGATE_TYPE_UNSPECIFIED\", \"APPROVAL_DELEGATE_TYPE_USER\", \"APPROVAL_DELEGATE_TYPE_GROUP\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"APPROVAL_DELEGATE_TYPE_UNSPECIFIED",
									"APPROVAL_DELEGATE_TYPE_USER",
									"APPROVAL_DELEGATE_TYPE_GROUP",
								),
							},
						},
					},
				},
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"delegator_user_id": schema.StringAttribute{
						Description: "Optional user ID to list delegations for a specific user.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *ApprovalDelegationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ApprovalDelegationDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("id"), path.MatchRoot("find_one_by")),
	}
}
