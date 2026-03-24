// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tag

import (
	"github.com/ServalHQ/terraform-provider-serval/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TagDataEnvelope struct {
	Data TagModel `json:"data"`
}

type TagModel struct {
	ID       types.String `tfsdk:"id" json:"id,computed"`
	TeamID   types.String `tfsdk:"team_id" json:"teamId,required"`
	Name     types.String `tfsdk:"name" json:"name,required"`
	Color    types.String `tfsdk:"color" json:"color,optional"`
	IconSlug types.String `tfsdk:"icon_slug" json:"iconSlug,optional"`
}

func (m TagModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m TagModel) MarshalJSONForUpdate(state TagModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
