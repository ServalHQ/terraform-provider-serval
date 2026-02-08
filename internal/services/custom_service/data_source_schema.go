// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_service

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
)

var _ datasource.DataSourceWithConfigValidators = (*CustomServiceDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The ID of the custom service.",
				Computed:    true,
				Optional:    true,
			},
			"domain": schema.StringAttribute{
				Description: `The domain for branding/logo lookup (e.g., "hr.company.com").`,
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: `The name of the custom service (e.g., "Internal HR System").`,
				Computed:    true,
			},
			"team_id": schema.StringAttribute{
				Description: "The ID of the team that the custom service belongs to.",
				Computed:    true,
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"team_id": schema.StringAttribute{
						Description: "The ID of the team.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *CustomServiceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *CustomServiceDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("id"), path.MatchRoot("find_one_by")),
	}
}
