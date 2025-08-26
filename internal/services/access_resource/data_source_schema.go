// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_resource

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*AccessResourceDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The ID of the resource.",
				Required:    true,
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
	}
}

func (d *AccessResourceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *AccessResourceDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
