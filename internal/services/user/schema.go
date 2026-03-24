// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*UserResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"email": schema.StringAttribute{
				Required: true,
			},
			"auth_method": schema.StringAttribute{
				Description: "Specifies the authentication method for a user. If unset, the org default applies.\n Set to MAGIC_LINK to allow the user to bypass SSO (e.g. guest or break-glass accounts).\nAvailable values: \"USER_AUTH_METHOD_UNSPECIFIED\", \"USER_AUTH_METHOD_MAGIC_LINK\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("USER_AUTH_METHOD_UNSPECIFIED", "USER_AUTH_METHOD_MAGIC_LINK"),
				},
			},
			"avatar_url": schema.StringAttribute{
				Optional: true,
			},
			"first_name": schema.StringAttribute{
				Optional: true,
			},
			"last_name": schema.StringAttribute{
				Optional: true,
			},
			"role": schema.StringAttribute{
				Description: `Available values: "USER_ROLE_UNSPECIFIED", "USER_ROLE_ORG_MEMBER", "USER_ROLE_ORG_ADMIN", "USER_ROLE_ORG_GUEST".`,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"USER_ROLE_UNSPECIFIED",
						"USER_ROLE_ORG_MEMBER",
						"USER_ROLE_ORG_ADMIN",
						"USER_ROLE_ORG_GUEST",
					),
				},
			},
			"created_at": schema.StringAttribute{
				Description:   `A timestamp in RFC 3339 format (e.g., "2025-01-15T01:30:15Z").`,
				Computed:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"deactivated_at": schema.StringAttribute{
				Description:   `A timestamp in RFC 3339 format (e.g., "2025-01-15T01:30:15Z").`,
				Computed:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"timezone": schema.StringAttribute{
				Description: `IANA timezone, e.g., "America/New_York"`,
				Computed:    true,
			},
		},
	}
}

func (r *UserResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *UserResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
