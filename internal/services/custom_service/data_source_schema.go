// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_service

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*CustomServiceDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The ID of the custom service.",
				Required:    true,
			},
			"domain": schema.StringAttribute{
				Description: `(OPTIONAL) The domain for branding/logo lookup (e.g., "hr.company.com").`,
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: `The name of the custom service (e.g., "Internal HR System").`,
				Computed:    true,
			},
			"team_id": schema.StringAttribute{
				Description: "(IMMUTABLE) The ID of the team that the custom service belongs to.",
				Computed:    true,
			},
		},
	}
}

func (d *CustomServiceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *CustomServiceDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
