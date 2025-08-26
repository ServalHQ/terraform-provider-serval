// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/serval-terraform/internal/apijson"
	"github.com/stainless-sdks/serval-terraform/internal/customfield"
)

type AccessPolicyModel struct {
	AccessPolicyID               types.String                                    `tfsdk:"access_policy_id" path:"access_policy_id,optional"`
	TeamID                       types.String                                    `tfsdk:"team_id" json:"teamId,optional,no_refresh"`
	Description                  types.String                                    `tfsdk:"description" json:"description,optional,no_refresh"`
	MaxAccessMinutes             types.Int64                                     `tfsdk:"max_access_minutes" json:"maxAccessMinutes,optional,no_refresh"`
	Name                         types.String                                    `tfsdk:"name" json:"name,optional,no_refresh"`
	RequireBusinessJustification types.Bool                                      `tfsdk:"require_business_justification" json:"requireBusinessJustification,optional,no_refresh"`
	Data                         customfield.NestedObject[AccessPolicyDataModel] `tfsdk:"data" json:"data,computed"`
}

func (m AccessPolicyModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AccessPolicyModel) MarshalJSONForUpdate(state AccessPolicyModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type AccessPolicyDataModel struct {
	ID                           types.String `tfsdk:"id" json:"id,computed"`
	Description                  types.String `tfsdk:"description" json:"description,computed"`
	MaxAccessMinutes             types.Int64  `tfsdk:"max_access_minutes" json:"maxAccessMinutes,computed"`
	Name                         types.String `tfsdk:"name" json:"name,computed"`
	RequireBusinessJustification types.Bool   `tfsdk:"require_business_justification" json:"requireBusinessJustification,computed"`
	TeamID                       types.String `tfsdk:"team_id" json:"teamId,computed"`
}
