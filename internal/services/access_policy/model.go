// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/serval-terraform/internal/apijson"
)

type AccessPolicyDataEnvelope struct {
	Data AccessPolicyModel `json:"data"`
}

type AccessPolicyModel struct {
	ID                           types.String `tfsdk:"id" json:"id,computed"`
	TeamID                       types.String `tfsdk:"team_id" json:"teamId,optional"`
	Description                  types.String `tfsdk:"description" json:"description,optional"`
	MaxAccessMinutes             types.Int64  `tfsdk:"max_access_minutes" json:"maxAccessMinutes,optional"`
	Name                         types.String `tfsdk:"name" json:"name,optional"`
	RequireBusinessJustification types.Bool   `tfsdk:"require_business_justification" json:"requireBusinessJustification,optional"`
}

func (m AccessPolicyModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AccessPolicyModel) MarshalJSONForUpdate(state AccessPolicyModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
