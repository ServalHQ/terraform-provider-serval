// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package approval_delegation

import (
	"context"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/serval-go/packages/param"
	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ApprovalDelegationDataDataSourceEnvelope struct {
	Data ApprovalDelegationDataSourceModel `json:"data,computed"`
}

type ApprovalDelegationDataSourceModel struct {
	ID              types.String                                                             `tfsdk:"id" path:"id,computed_optional"`
	CreatedAt       timetypes.RFC3339                                                        `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	DelegatorUserID types.String                                                             `tfsdk:"delegator_user_id" json:"delegatorUserId,computed"`
	Description     types.String                                                             `tfsdk:"description" json:"description,computed"`
	Priority        types.Int64                                                              `tfsdk:"priority" json:"priority,computed"`
	Delegates       customfield.NestedObjectList[ApprovalDelegationDelegatesDataSourceModel] `tfsdk:"delegates" json:"delegates,computed"`
	FindOneBy       *ApprovalDelegationFindOneByDataSourceModel                              `tfsdk:"find_one_by"`
}

func (m *ApprovalDelegationDataSourceModel) toListParams(_ context.Context) (params serval.ApprovalDelegationListParams, diags diag.Diagnostics) {
	params = serval.ApprovalDelegationListParams{}

	if !m.FindOneBy.DelegatorUserID.IsNull() {
		params.DelegatorUserID = param.NewOpt(m.FindOneBy.DelegatorUserID.ValueString())
	}

	return
}

type ApprovalDelegationDelegatesDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type ApprovalDelegationFindOneByDataSourceModel struct {
	DelegatorUserID types.String `tfsdk:"delegator_user_id" query:"delegatorUserId,optional"`
}
