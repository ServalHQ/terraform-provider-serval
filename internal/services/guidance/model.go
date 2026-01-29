// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package guidance

import (
	"github.com/ServalHQ/terraform-provider-serval/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type GuidanceDataEnvelope struct {
	Data GuidanceModel `json:"data"`
}

type GuidanceModel struct {
	ID                    types.String `tfsdk:"id" json:"id,computed"`
	TeamID                types.String `tfsdk:"team_id" json:"teamId,optional"`
	Content               types.String `tfsdk:"content" json:"content,optional"`
	Description           types.String `tfsdk:"description" json:"description,optional"`
	Name                  types.String `tfsdk:"name" json:"name,optional"`
	IsPublished           types.Bool   `tfsdk:"is_published" json:"isPublished,computed_optional"`
	ShouldAlwaysUse       types.Bool   `tfsdk:"should_always_use" json:"shouldAlwaysUse,computed_optional"`
	HasUnpublishedChanges types.Bool   `tfsdk:"has_unpublished_changes" json:"hasUnpublishedChanges,computed"`
}

func (m GuidanceModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m GuidanceModel) MarshalJSONForUpdate(state GuidanceModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
