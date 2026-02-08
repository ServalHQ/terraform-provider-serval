// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_instance

import (
	"context"

	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*AppInstancesDataSource)(nil)

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
				CustomType:  customfield.NewNestedObjectListType[AppInstancesItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The ID of the app instance.",
							Computed:    true,
						},
						"access_requests_enabled": schema.BoolAttribute{
							Description: "Whether access requests are enabled for the app instance.",
							Computed:    true,
						},
						"custom_service_id": schema.StringAttribute{
							Description: "**Option: custom_service_id** — The ID of the custom service (for custom apps).",
							Computed:    true,
						},
						"default_access_policy_id": schema.StringAttribute{
							Description: "The default access policy for the app instance.",
							Computed:    true,
						},
						"external_service_instance_id": schema.StringAttribute{
							Description: "The external service instance ID (e.g., GitHub org name, Okta domain, AWS account ID).",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the app instance.",
							Computed:    true,
						},
						"service": schema.StringAttribute{
							Description: `**Option: service** — The service identifier (for built-in services like "github", "okta", "aws").`,
							Computed:    true,
						},
						"team_id": schema.StringAttribute{
							Description: "The ID of the Serval team that the app instance belongs to.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *AppInstancesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *AppInstancesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
