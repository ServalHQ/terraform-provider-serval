// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package guidance

import (
	"context"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/serval-go/packages/param"
	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type GuidancesDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[GuidancesItemsDataSourceModel] `json:"data,computed"`
}

type GuidancesDataSourceModel struct {
	TeamID   types.String                                                `tfsdk:"team_id" query:"teamId,optional"`
	MaxItems types.Int64                                                 `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[GuidancesItemsDataSourceModel] `tfsdk:"items"`
}

func (m *GuidancesDataSourceModel) toListParams(_ context.Context) (params serval.GuidanceListParams, diags diag.Diagnostics) {
	params = serval.GuidanceListParams{}

	if !m.TeamID.IsNull() {
		params.TeamID = param.NewOpt(m.TeamID.ValueString())
	}

	return
}

type GuidancesItemsDataSourceModel struct {
	ID                    types.String `tfsdk:"id" json:"id,computed"`
	Content               types.String `tfsdk:"content" json:"content,computed"`
	Description           types.String `tfsdk:"description" json:"description,computed"`
	HasUnpublishedChanges types.Bool   `tfsdk:"has_unpublished_changes" json:"hasUnpublishedChanges,computed"`
	IsPublished           types.Bool   `tfsdk:"is_published" json:"isPublished,computed"`
	Name                  types.String `tfsdk:"name" json:"name,computed"`
	ShouldAlwaysUse       types.Bool   `tfsdk:"should_always_use" json:"shouldAlwaysUse,computed"`
	TeamID                types.String `tfsdk:"team_id" json:"teamId,computed"`
}
