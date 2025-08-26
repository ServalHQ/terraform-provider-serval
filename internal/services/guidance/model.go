// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package guidance

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/serval-terraform/internal/apijson"
)

type GuidanceDataEnvelope struct {
	Data GuidanceModel `json:"data"`
}

type GuidanceModel struct {
	ID          types.String `tfsdk:"id" json:"id,computed"`
	TeamID      types.String `tfsdk:"team_id" json:"teamId,optional"`
	Content     types.String `tfsdk:"content" json:"content,optional"`
	Description types.String `tfsdk:"description" json:"description,optional"`
	Name        types.String `tfsdk:"name" json:"name,optional"`
}

func (m GuidanceModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m GuidanceModel) MarshalJSONForUpdate(state GuidanceModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
