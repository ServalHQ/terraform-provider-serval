// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package guidance

import (
	"context"

	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*GuidancesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"team_id": schema.StringAttribute{
				Description: "The ID of the team.",
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
				CustomType:  customfield.NewNestedObjectListType[GuidancesItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The ID of the guidance.",
							Computed:    true,
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
					},
				},
			},
		},
	}
}

func (d *GuidancesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *GuidancesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
