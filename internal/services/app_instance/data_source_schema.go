// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_instance

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*AppInstanceDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The ID of the app instance.",
				Required:    true,
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
	}
}

func (d *AppInstanceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *AppInstanceDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
