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
			"team_id": schema.StringAttribute{
				Description:   "The ID of the team.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"custom_service_id": schema.StringAttribute{
				Description:   "**Option: custom_service_id** — The ID of a custom service to create the app instance for.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"service": schema.StringAttribute{
				Description:   `**Option: service** — The service identifier (for built-in services like "github", "okta", "aws").`,
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"external_service_instance_id": schema.StringAttribute{
				Description: "The external service instance ID (e.g., GitHub org name, Okta domain, AWS account ID).",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the app instance.",
				Required:    true,
			},
			"default_access_policy_id": schema.StringAttribute{
				Description: "The default access policy for the app instance (optional).",
				Optional:    true,
			},
			"access_requests_enabled": schema.BoolAttribute{
				Description: "Whether access requests are enabled for the app instance.",
				Computed:    true,
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
