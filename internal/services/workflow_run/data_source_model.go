// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflow_run

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkflowRunDataDataSourceEnvelope struct {
	Data WorkflowRunDataSourceModel `json:"data,computed"`
}

type WorkflowRunDataSourceModel struct {
	ID                types.String `tfsdk:"id" path:"id,required"`
	CompletedAt       types.String `tfsdk:"completed_at" json:"completedAt,computed"`
	CreatedAt         types.String `tfsdk:"created_at" json:"createdAt,computed"`
	InitiatedByUserID types.String `tfsdk:"initiated_by_user_id" json:"initiatedByUserId,computed"`
	Inputs            types.String `tfsdk:"inputs" json:"inputs,computed"`
	LinkedTicketID    types.String `tfsdk:"linked_ticket_id" json:"linkedTicketId,computed"`
	Output            types.String `tfsdk:"output" json:"output,computed"`
	Status            types.String `tfsdk:"status" json:"status,computed"`
	TargetUserID      types.String `tfsdk:"target_user_id" json:"targetUserId,computed"`
	TeamID            types.String `tfsdk:"team_id" json:"teamId,computed"`
	WorkflowID        types.String `tfsdk:"workflow_id" json:"workflowId,computed"`
}
