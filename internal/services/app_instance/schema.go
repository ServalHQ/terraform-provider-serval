// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_instance

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*AppInstanceResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The ID of the app instance.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"custom_service_id": schema.StringAttribute{
				Description:   "The ID of a custom service to create the app instance for.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"service": schema.StringAttribute{
				Description:   `The service identifier (for built-in services like "github", "okta", "aws").`,
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"team_id": schema.StringAttribute{
				Description:   "The ID of the team.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"access_requests_enabled": schema.BoolAttribute{
				Description: "Whether access requests are enabled for the app instance.",
				Optional:    true,
			},
			"default_access_policy_id": schema.StringAttribute{
				Description: "The default access policy for the app instance (optional).",
				Optional:    true,
			},
			"instance_id": schema.StringAttribute{
				Description: "The instance ID of the app instance.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the app instance.",
				Optional:    true,
			},
		},
	}
}

func (r *AppInstanceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *AppInstanceResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
