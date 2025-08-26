// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/serval-go"
	"github.com/stainless-sdks/serval-go/packages/param"
)

type AccessPolicyDataDataSourceEnvelope struct {
	Data AccessPolicyDataSourceModel `json:"data,computed"`
}

type AccessPolicyDataSourceModel struct {
	ID                           types.String `tfsdk:"id" path:"id,required"`
	AccessPolicyID               types.String `tfsdk:"access_policy_id" query:"accessPolicyId,optional"`
	Description                  types.String `tfsdk:"description" json:"description,computed"`
	MaxAccessMinutes             types.Int64  `tfsdk:"max_access_minutes" json:"maxAccessMinutes,computed"`
	Name                         types.String `tfsdk:"name" json:"name,computed"`
	RequireBusinessJustification types.Bool   `tfsdk:"require_business_justification" json:"requireBusinessJustification,computed"`
	TeamID                       types.String `tfsdk:"team_id" json:"teamId,computed"`
}

func (m *AccessPolicyDataSourceModel) toReadParams(_ context.Context) (params serval.AccessPolicyGetParams, diags diag.Diagnostics) {
	params = serval.AccessPolicyGetParams{}

	if !m.AccessPolicyID.IsNull() {
		params.AccessPolicyID = param.NewOpt(m.AccessPolicyID.ValueString())
	}

	return
}
