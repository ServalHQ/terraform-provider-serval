// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy

import (
	"context"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/serval-go/packages/param"
	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessPoliciesDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[AccessPoliciesItemsDataSourceModel] `json:"data,computed"`
}

type AccessPoliciesDataSourceModel struct {
	TeamID   types.String                                                     `tfsdk:"team_id" query:"teamId,optional"`
	MaxItems types.Int64                                                      `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[AccessPoliciesItemsDataSourceModel] `tfsdk:"items"`
}

func (m *AccessPoliciesDataSourceModel) toListParams(_ context.Context) (params serval.AccessPolicyListParams, diags diag.Diagnostics) {
	params = serval.AccessPolicyListParams{}

	if !m.TeamID.IsNull() {
		params.TeamID = param.NewOpt(m.TeamID.ValueString())
	}

	return
}

type AccessPoliciesItemsDataSourceModel struct {
	ID                           types.String `tfsdk:"id" json:"id,computed"`
	Description                  types.String `tfsdk:"description" json:"description,computed"`
	MaxAccessMinutes             types.Int64  `tfsdk:"max_access_minutes" json:"maxAccessMinutes,computed"`
	Name                         types.String `tfsdk:"name" json:"name,computed"`
	RecommendedAccessMinutes     types.Int64  `tfsdk:"recommended_access_minutes" json:"recommendedAccessMinutes,computed"`
	RequireBusinessJustification types.Bool   `tfsdk:"require_business_justification" json:"requireBusinessJustification,computed"`
	TeamID                       types.String `tfsdk:"team_id" json:"teamId,computed"`
}
