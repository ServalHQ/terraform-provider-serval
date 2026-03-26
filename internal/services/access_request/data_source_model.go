// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_request

import (
	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessRequestDataDataSourceEnvelope struct {
	Data AccessRequestDataSourceModel `json:"data,computed"`
}

type AccessRequestDataSourceModel struct {
	ID              types.String                                                              `tfsdk:"id" path:"id,required"`
	CreatedAt       types.String                                                              `tfsdk:"created_at" json:"createdAt,computed"`
	ExpiresAt       types.String                                                              `tfsdk:"expires_at" json:"expiresAt,computed"`
	LinkedTicketID  types.String                                                              `tfsdk:"linked_ticket_id" json:"linkedTicketId,computed"`
	RequestedRoleID types.String                                                              `tfsdk:"requested_role_id" json:"requestedRoleId,computed"`
	Status          types.String                                                              `tfsdk:"status" json:"status,computed"`
	TargetUserID    types.String                                                              `tfsdk:"target_user_id" json:"targetUserId,computed"`
	TeamID          types.String                                                              `tfsdk:"team_id" json:"teamId,computed"`
	TimeAllocations customfield.NestedObjectList[AccessRequestTimeAllocationsDataSourceModel] `tfsdk:"time_allocations" json:"timeAllocations,computed"`
}

type AccessRequestTimeAllocationsDataSourceModel struct {
	ID                    types.String `tfsdk:"id" json:"id,computed"`
	ApprovalRequestID     types.String `tfsdk:"approval_request_id" json:"approvalRequestId,computed"`
	ApprovedMinutes       types.Int64  `tfsdk:"approved_minutes" json:"approvedMinutes,computed"`
	BusinessJustification types.String `tfsdk:"business_justification" json:"businessJustification,computed"`
	CreatedAt             types.String `tfsdk:"created_at" json:"createdAt,computed"`
	InvalidationReason    types.String `tfsdk:"invalidation_reason" json:"invalidationReason,computed"`
	LinkedTicketID        types.String `tfsdk:"linked_ticket_id" json:"linkedTicketId,computed"`
	RequestedByUserID     types.String `tfsdk:"requested_by_user_id" json:"requestedByUserId,computed"`
	RequestedMinutes      types.Int64  `tfsdk:"requested_minutes" json:"requestedMinutes,computed"`
	Status                types.String `tfsdk:"status" json:"status,computed"`
}
