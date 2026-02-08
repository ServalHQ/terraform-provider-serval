// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_resource

import (
	"context"

	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*AppResourcesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"app_instance_id": schema.StringAttribute{
				Description: "Filter by app instance ID. At least one of team_id or app_instance_id must be provided.",
				Optional:    true,
			},
			"team_id": schema.StringAttribute{
				Description: "Filter by team ID. At least one of team_id or app_instance_id must be provided.",
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
				CustomType:  customfield.NewNestedObjectListType[AppResourcesItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The ID of the resource.",
							Computed:    true,
						},
						"app_instance_id": schema.StringAttribute{
							Description: "The ID of the app instance that the resource belongs to.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "A description of the resource.",
							Computed:    true,
						},
						"external_id": schema.StringAttribute{
							Description: "The external ID of the resource.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the resource.",
							Computed:    true,
						},
						"resource_type": schema.StringAttribute{
							Description: "The type of the resource.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *AppResourcesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *AppResourcesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
