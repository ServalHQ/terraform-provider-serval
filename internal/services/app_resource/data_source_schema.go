// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_resource

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
)

var _ datasource.DataSourceWithConfigValidators = (*AppResourceDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The ID of the resource.",
				Computed:    true,
				Optional:    true,
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
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"app_instance_id": schema.StringAttribute{
						Description: "Filter by app instance ID. At least one of team_id or app_instance_id must be provided.",
						Optional:    true,
					},
					"team_id": schema.StringAttribute{
						Description: "Filter by team ID. At least one of team_id or app_instance_id must be provided.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *AppResourceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *AppResourceDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("id"), path.MatchRoot("find_one_by")),
	}
}
