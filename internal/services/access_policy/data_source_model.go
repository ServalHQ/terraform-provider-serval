// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy

import (
	"context"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/serval-go/packages/param"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessPolicyDataDataSourceEnvelope struct {
	Data AccessPolicyDataSourceModel `json:"data,computed"`
}

type AccessPolicyDataSourceModel struct {
	ID                           types.String                          `tfsdk:"id" path:"id,computed_optional"`
	Description                  types.String                          `tfsdk:"description" json:"description,computed"`
	MaxAccessMinutes             types.Int64                           `tfsdk:"max_access_minutes" json:"maxAccessMinutes,computed"`
	Name                         types.String                          `tfsdk:"name" json:"name,computed"`
	RecommendedAccessMinutes     types.Int64                           `tfsdk:"recommended_access_minutes" json:"recommendedAccessMinutes,computed"`
	RequireBusinessJustification types.Bool                            `tfsdk:"require_business_justification" json:"requireBusinessJustification,computed"`
	TeamID                       types.String                          `tfsdk:"team_id" json:"teamId,computed"`
	FindOneBy                    *AccessPolicyFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

func (m *AccessPolicyDataSourceModel) toListParams(_ context.Context) (params serval.AccessPolicyListParams, diags diag.Diagnostics) {
	params = serval.AccessPolicyListParams{}

	if !m.FindOneBy.TeamID.IsNull() {
		params.TeamID = param.NewOpt(m.FindOneBy.TeamID.ValueString())
	}

	return
}

type AccessPolicyFindOneByDataSourceModel struct {
	TeamID types.String `tfsdk:"team_id" query:"teamId,optional"`
}
