// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessPolicyDataDataSourceEnvelope struct {
	Data AccessPolicyDataSourceModel `json:"data,computed"`
}

type AccessPolicyDataSourceModel struct {
	AccessPolicyID               types.String `tfsdk:"access_policy_id" path:"access_policy_id,required"`
	Description                  types.String `tfsdk:"description" json:"description,computed"`
	ID                           types.String `tfsdk:"id" json:"id,computed"`
	MaxAccessMinutes             types.Int64  `tfsdk:"max_access_minutes" json:"maxAccessMinutes,computed"`
	Name                         types.String `tfsdk:"name" json:"name,computed"`
	RequireBusinessJustification types.Bool   `tfsdk:"require_business_justification" json:"requireBusinessJustification,computed"`
	TeamID                       types.String `tfsdk:"team_id" json:"teamId,computed"`
}
