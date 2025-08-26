// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*AccessPolicyDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The ID of the access policy.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "A description of the access policy.",
				Computed:    true,
			},
			"max_access_minutes": schema.Int64Attribute{
				Description: "The maximum number of minutes that access can be granted for.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the access policy.",
				Computed:    true,
			},
			"require_business_justification": schema.BoolAttribute{
				Description: "Whether a business justification is required when requesting access.",
				Computed:    true,
			},
			"team_id": schema.StringAttribute{
				Description: "The ID of the team that the access policy belongs to.",
				Computed:    true,
			},
		},
	}
}

func (d *AccessPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *AccessPolicyDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
