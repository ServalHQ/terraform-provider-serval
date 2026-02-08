// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_request

import (
	"context"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/serval-go/packages/param"
	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessRequestsDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[AccessRequestsItemsDataSourceModel] `json:"data,computed"`
}

type AccessRequestsDataSourceModel struct {
	TeamID   types.String                                                     `tfsdk:"team_id" query:"teamId,optional"`
	MaxItems types.Int64                                                      `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[AccessRequestsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *AccessRequestsDataSourceModel) toListParams(_ context.Context) (params serval.AccessRequestListParams, diags diag.Diagnostics) {
	params = serval.AccessRequestListParams{}

	if !m.TeamID.IsNull() {
		params.TeamID = param.NewOpt(m.TeamID.ValueString())
	}

	return
}

type AccessRequestsItemsDataSourceModel struct {
	ID                    types.String `tfsdk:"id" json:"id,computed"`
	AccessMinutes         types.Int64  `tfsdk:"access_minutes" json:"accessMinutes,computed"`
	BusinessJustification types.String `tfsdk:"business_justification" json:"businessJustification,computed"`
	CreatedAt             types.String `tfsdk:"created_at" json:"createdAt,computed"`
	ExpiresAt             types.String `tfsdk:"expires_at" json:"expiresAt,computed"`
	LinkedTicketID        types.String `tfsdk:"linked_ticket_id" json:"linkedTicketId,computed"`
	RequestedByUserID     types.String `tfsdk:"requested_by_user_id" json:"requestedByUserId,computed"`
	RequestedRoleID       types.String `tfsdk:"requested_role_id" json:"requestedRoleId,computed"`
	Status                types.String `tfsdk:"status" json:"status,computed"`
	TargetUserID          types.String `tfsdk:"target_user_id" json:"targetUserId,computed"`
	TeamID                types.String `tfsdk:"team_id" json:"teamId,computed"`
}
