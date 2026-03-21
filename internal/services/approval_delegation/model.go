// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package approval_delegation

import (
	"github.com/ServalHQ/terraform-provider-serval/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ApprovalDelegationDataEnvelope struct {
	Data ApprovalDelegationModel `json:"data"`
}

type ApprovalDelegationModel struct {
	ID              types.String                         `tfsdk:"id" json:"id,computed"`
	DelegatorUserID types.String                         `tfsdk:"delegator_user_id" json:"delegatorUserId,computed_optional"`
	Delegates       *[]*ApprovalDelegationDelegatesModel `tfsdk:"delegates" json:"delegates,required"`
	Description     types.String                         `tfsdk:"description" json:"description,optional"`
	Priority        types.Int64                          `tfsdk:"priority" json:"priority,optional"`
	CreatedAt       timetypes.RFC3339                    `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
}

func (m ApprovalDelegationModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ApprovalDelegationModel) MarshalJSONForUpdate(state ApprovalDelegationModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ApprovalDelegationDelegatesModel struct {
	ID   types.String `tfsdk:"id" json:"id,optional"`
	Type types.String `tfsdk:"type" json:"type,optional"`
}
