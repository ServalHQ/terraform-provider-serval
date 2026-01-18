// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tag

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*TagDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The ID of the tag.",
				Required:    true,
			},
			"color": schema.StringAttribute{
				Description: "The color of the tag (CSS color string).",
				Computed:    true,
			},
			"icon_slug": schema.StringAttribute{
				Description: "The icon slug for the tag.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the tag.",
				Computed:    true,
			},
		},
	}
}

func (d *TagDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *TagDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
