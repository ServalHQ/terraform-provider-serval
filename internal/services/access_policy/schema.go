// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*AccessPolicyResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The ID of the access policy.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"team_id": schema.StringAttribute{
				Description:   "The ID of the team.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "The name of the access policy.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "A description of the access policy.",
				Optional:    true,
			},
			"max_access_minutes": schema.Int64Attribute{
				Description: "The maximum number of minutes that access can be granted for (optional).",
				Optional:    true,
			},
			"recommended_access_minutes": schema.Int64Attribute{
				Description: "The recommended duration in minutes for access requests (optional).",
				Optional:    true,
			},
			"require_business_justification": schema.BoolAttribute{
				Description: "Whether a business justification is required when requesting access (optional).",
				Computed:    true,
				Optional:    true,
			},
		},
	}
}

func (r *AccessPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *AccessPolicyResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
