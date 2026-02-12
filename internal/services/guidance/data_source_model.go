// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package guidance

import (
	"context"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/serval-go/packages/param"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type GuidanceDataDataSourceEnvelope struct {
	Data GuidanceDataSourceModel `json:"data,computed"`
}

type GuidanceDataSourceModel struct {
	ID                    types.String                      `tfsdk:"id" path:"id,computed_optional"`
	Content               types.String                      `tfsdk:"content" json:"content,computed"`
	Description           types.String                      `tfsdk:"description" json:"description,computed"`
	HasUnpublishedChanges types.Bool                        `tfsdk:"has_unpublished_changes" json:"hasUnpublishedChanges,computed"`
	IsPublished           types.Bool                        `tfsdk:"is_published" json:"isPublished,computed"`
	Name                  types.String                      `tfsdk:"name" json:"name,computed"`
	ShouldAlwaysUse       types.Bool                        `tfsdk:"should_always_use" json:"shouldAlwaysUse,computed"`
	TeamID                types.String                      `tfsdk:"team_id" json:"teamId,computed"`
	FindOneBy             *GuidanceFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

func (m *GuidanceDataSourceModel) toListParams(_ context.Context) (params serval.GuidanceListParams, diags diag.Diagnostics) {
	params = serval.GuidanceListParams{}

	if !m.FindOneBy.TeamID.IsNull() {
		params.TeamID = param.NewOpt(m.FindOneBy.TeamID.ValueString())
	}

	return
}

type GuidanceFindOneByDataSourceModel struct {
	TeamID types.String `tfsdk:"team_id" query:"teamId,optional"`
}
