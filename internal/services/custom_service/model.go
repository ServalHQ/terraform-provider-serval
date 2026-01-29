// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_service

import (
	"github.com/ServalHQ/terraform-provider-serval/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomServiceDataEnvelope struct {
	Data CustomServiceModel `json:"data"`
}

type CustomServiceModel struct {
	ID     types.String `tfsdk:"id" json:"id,computed"`
	TeamID types.String `tfsdk:"team_id" json:"teamId,required"`
	Name   types.String `tfsdk:"name" json:"name,required"`
	Domain types.String `tfsdk:"domain" json:"domain,optional"`
}

func (m CustomServiceModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CustomServiceModel) MarshalJSONForUpdate(state CustomServiceModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
