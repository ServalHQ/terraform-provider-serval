// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package group

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*GroupResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"user_ids": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"created_at": schema.StringAttribute{
				Description:   `A timestamp in RFC 3339 format (e.g., "2017-01-15T01:30:15.01Z").`,
				Computed:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"deleted_at": schema.StringAttribute{
				Description:   `A timestamp in RFC 3339 format (e.g., "2017-01-15T01:30:15.01Z").`,
				Computed:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{useNullableStateForUnknown()},
			},
			"organization_id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

// useNullableStateForUnknownModifier copies the prior state value to the plan
// even when the state value is null. This differs from
// stringplanmodifier.UseStateForUnknown() which skips null state values.
//
// This is needed for computed-only fields like deleted_at where null is a
// valid server-computed value ("not deleted") that should be preserved across
// plans. Without this, UseStateForUnknown leaves the plan as unknown when
// state is null, producing a spurious diff on every update.
type useNullableStateForUnknownModifier struct{}

func useNullableStateForUnknown() planmodifier.String {
	return useNullableStateForUnknownModifier{}
}

func (m useNullableStateForUnknownModifier) Description(_ context.Context) string {
	return "Copies the prior state value (including null) to the plan when the plan value is unknown."
}

func (m useNullableStateForUnknownModifier) MarkdownDescription(_ context.Context) string {
	return "Copies the prior state value (including null) to the plan when the plan value is unknown."
}

func (m useNullableStateForUnknownModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if !req.PlanValue.IsUnknown() {
		return
	}
	if req.ConfigValue.IsUnknown() {
		return
	}
	if req.StateValue.IsUnknown() {
		return
	}
	resp.PlanValue = req.StateValue
}

func (r *GroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *GroupResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
