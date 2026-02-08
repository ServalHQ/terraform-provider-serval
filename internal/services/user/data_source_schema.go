// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*UserDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional: true,
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
				Optional: true,
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
	}
}

func (d *UserDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *UserDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("email"),
		),
	}
}
