// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package team_user

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*TeamUserDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"user_id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"team_id": schema.StringAttribute{
				Required: true,
			},
			"id": schema.StringAttribute{
				Description: "Composite identifier ({team_id}:{user_id}). Alternative to team_id + user_id.",
				Computed:    true,
				Optional:    true,
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
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"user_id": schema.StringAttribute{
						Description: "Filter by user ID. If not provided, returns all users in the team.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *TeamUserDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *TeamUserDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("user_id"), path.MatchRoot("find_one_by")),
	}
}
