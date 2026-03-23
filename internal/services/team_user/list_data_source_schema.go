// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package team_user

import (
	"context"

	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*TeamUsersDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"team_id": schema.StringAttribute{
				Description: "Filter by team ID. Required when using /v2/teams/{team_id}/users path.",
				Required:    true,
			},
			"user_id": schema.StringAttribute{
				Description: "Filter by user ID. If not provided, returns all users in the team.",
				Optional:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[TeamUsersItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Composite identifier in the format {team_id}:{user_id}.",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Description: `A timestamp in RFC 3339 format (e.g., "2025-01-15T01:30:15Z").`,
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"role": schema.StringAttribute{
							Description: `Available values: "TEAM_USER_ROLE_UNSPECIFIED", "TEAM_USER_ROLE_AGENT", "TEAM_USER_ROLE_MANAGER", "TEAM_USER_ROLE_BUILDER", "TEAM_USER_ROLE_VIEWER", "TEAM_USER_ROLE_CONTRIBUTOR".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"TEAM_USER_ROLE_UNSPECIFIED",
									"TEAM_USER_ROLE_AGENT",
									"TEAM_USER_ROLE_MANAGER",
									"TEAM_USER_ROLE_BUILDER",
									"TEAM_USER_ROLE_VIEWER",
									"TEAM_USER_ROLE_CONTRIBUTOR",
								),
							},
						},
						"team_id": schema.StringAttribute{
							Computed: true,
						},
						"user_id": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *TeamUsersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *TeamUsersDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
