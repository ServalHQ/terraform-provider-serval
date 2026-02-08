// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package guidance

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
)

var _ datasource.DataSourceWithConfigValidators = (*GuidanceDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The ID of the guidance.",
				Computed:    true,
				Optional:    true,
			},
			"content": schema.StringAttribute{
				Description: "The content of the guidance.",
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Description: "A description of the guidance.",
				Computed:    true,
			},
			"has_unpublished_changes": schema.BoolAttribute{
				Description: "Whether there are unpublished changes to the guidance (computed by server).",
				Computed:    true,
			},
			"is_published": schema.BoolAttribute{
				Description: "Whether the guidance is published. Set to true to publish the guidance.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the guidance.",
				Computed:    true,
			},
			"should_always_use": schema.BoolAttribute{
				Description: "Whether this guidance should always be used (skipping LLM selection).",
				Computed:    true,
			},
			"team_id": schema.StringAttribute{
				Description: "The ID of the team that the guidance belongs to.",
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

func (d *GuidanceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *GuidanceDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("id"), path.MatchRoot("find_one_by")),
	}
}
