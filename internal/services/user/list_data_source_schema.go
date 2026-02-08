// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"

	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*UsersDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"include_deactivated": schema.BoolAttribute{
				Description: "Whether to include deactivated users in the response.",
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
				CustomType:  customfield.NewNestedObjectListType[UsersItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"avatar_url": schema.StringAttribute{
							Computed: true,
						},
						"created_at": schema.StringAttribute{
							Description: `A timestamp in RFC 3339 format (e.g., "2017-01-15T01:30:15.01Z").`,
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"deactivated_at": schema.StringAttribute{
							Description: `A timestamp in RFC 3339 format (e.g., "2017-01-15T01:30:15.01Z").`,
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"email": schema.StringAttribute{
							Computed: true,
						},
						"first_name": schema.StringAttribute{
							Computed: true,
						},
						"last_name": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"role": schema.StringAttribute{
							Description: `Available values: "USER_ROLE_UNSPECIFIED", "USER_ROLE_ORG_MEMBER", "USER_ROLE_ORG_ADMIN", "USER_ROLE_ORG_GUEST".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"USER_ROLE_UNSPECIFIED",
									"USER_ROLE_ORG_MEMBER",
									"USER_ROLE_ORG_ADMIN",
									"USER_ROLE_ORG_GUEST",
								),
							},
						},
						"timezone": schema.StringAttribute{
							Description: `IANA timezone, e.g., "America/New_York"`,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *UsersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *UsersDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
