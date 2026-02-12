// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy

import (
	"context"

	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*AccessPoliciesDataSource)(nil)

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
				CustomType:  customfield.NewNestedObjectListType[AccessPoliciesItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The ID of the access policy.",
							Computed:    true,
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
						"recommended_access_minutes": schema.Int64Attribute{
							Description: "The recommended duration in minutes for access requests (optional).",
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
				},
			},
		},
	}
}

func (d *AccessPoliciesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *AccessPoliciesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
