// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_resource_entitlement

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*AppResourceEntitlementDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The ID of the entitlement.",
				Required:    true,
			},
			"access_policy_id": schema.StringAttribute{
				Description: "The default access policy for the entitlement.",
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Description: "A description of the entitlement.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the entitlement.",
				Computed:    true,
			},
			"provisioning_method": schema.StringAttribute{
				Description: "The provisioning method for the entitlement.",
				Computed:    true,
			},
			"requests_enabled": schema.BoolAttribute{
				Description: "Whether requests are enabled for the entitlement.",
				Computed:    true,
			},
			"resource_id": schema.StringAttribute{
				Description: "The ID of the resource that the entitlement belongs to.",
				Computed:    true,
			},
		},
	}
}

func (d *AppResourceEntitlementDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *AppResourceEntitlementDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
