// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tag

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
)

var _ datasource.DataSourceWithConfigValidators = (*TagDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The ID of the tag.",
				Computed:    true,
				Optional:    true,
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
			"team_id": schema.StringAttribute{
				Description: "The ID of the team that owns this tag.",
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

func (d *TagDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *TagDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("id"), path.MatchRoot("find_one_by")),
	}
}
