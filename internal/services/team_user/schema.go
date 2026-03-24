// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package team_user

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

var _ resource.ResourceWithConfigValidators = (*TeamUserResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Composite identifier in the format {team_id}:{user_id}.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"role": schema.StringAttribute{
				Description: `Available values: "TEAM_USER_ROLE_UNSPECIFIED", "TEAM_USER_ROLE_AGENT", "TEAM_USER_ROLE_MANAGER", "TEAM_USER_ROLE_BUILDER", "TEAM_USER_ROLE_VIEWER", "TEAM_USER_ROLE_CONTRIBUTOR".`,
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"TEAM_USER_ROLE_UNSPECIFIED",
						"TEAM_USER_ROLE_AGENT",
						"TEAM_USER_ROLE_MANAGER",
						"TEAM_USER_ROLE_BUILDER",
						"TEAM_USER_ROLE_VIEWER",
						"TEAM_USER_ROLE_CONTRIBUTOR",
					),
				},
			},
			"team_id": schema.StringAttribute{
				Required: true,
			},
			"user_id": schema.StringAttribute{
				Required: true,
			},
			"created_at": schema.StringAttribute{
				Description: `A timestamp in RFC 3339 format (e.g., "2025-01-15T01:30:15Z").`,
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *TeamUserResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *TeamUserResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
