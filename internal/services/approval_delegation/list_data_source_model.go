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

type ApprovalDelegationsDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[ApprovalDelegationsItemsDataSourceModel] `json:"data,computed"`
}

type ApprovalDelegationsDataSourceModel struct {
	DelegatorUserID types.String                                                          `tfsdk:"delegator_user_id" query:"delegatorUserId,optional"`
	MaxItems        types.Int64                                                           `tfsdk:"max_items"`
	Items           customfield.NestedObjectList[ApprovalDelegationsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *ApprovalDelegationsDataSourceModel) toListParams(_ context.Context) (params serval.ApprovalDelegationListParams, diags diag.Diagnostics) {
	params = serval.ApprovalDelegationListParams{}

	if !m.DelegatorUserID.IsNull() {
		params.DelegatorUserID = param.NewOpt(m.DelegatorUserID.ValueString())
	}

	return
}

type ApprovalDelegationsItemsDataSourceModel struct {
	ID              types.String                                                              `tfsdk:"id" json:"id,computed"`
	CreatedAt       timetypes.RFC3339                                                         `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	Delegates       customfield.NestedObjectList[ApprovalDelegationsDelegatesDataSourceModel] `tfsdk:"delegates" json:"delegates,computed"`
	DelegatorUserID types.String                                                              `tfsdk:"delegator_user_id" json:"delegatorUserId,computed"`
	Description     types.String                                                              `tfsdk:"description" json:"description,computed"`
	Priority        types.Int64                                                               `tfsdk:"priority" json:"priority,computed"`
}

type ApprovalDelegationsDelegatesDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}
