// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package guidance

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*GuidanceResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The ID of the guidance.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"team_id": schema.StringAttribute{
				Description:   "The ID of the team.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"content": schema.StringAttribute{
				Description: "The content of the guidance (optional).",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "A description of the guidance.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the guidance.",
				Optional:    true,
			},
			"should_always_use": schema.BoolAttribute{
				Description: "Whether this guidance should always be used (optional, defaults to false).",
				Optional:    true,
			},
			"has_unpublished_changes": schema.BoolAttribute{
				Description:   "Whether there are unpublished changes to the guidance.",
				Computed:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"is_published": schema.BoolAttribute{
				Description:   "Whether the guidance has been published at least once.",
				Computed:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (r *GuidanceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *GuidanceResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
