// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package guidance

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*GuidanceDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The ID of the guidance.",
				Required:    true,
			},
			"content": schema.StringAttribute{
				Description: "The content of the guidance.",
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Description: "A description of the guidance.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the guidance.",
				Computed:    true,
			},
			"team_id": schema.StringAttribute{
				Description: "The ID of the team that the guidance belongs to.",
				Computed:    true,
			},
		},
	}
}

func (d *GuidanceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *GuidanceDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
