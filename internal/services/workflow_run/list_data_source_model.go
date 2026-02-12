// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflow_run

import (
	"context"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/serval-go/packages/param"
	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkflowRunsDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[WorkflowRunsItemsDataSourceModel] `json:"data,computed"`
}

type WorkflowRunsDataSourceModel struct {
	TeamID   types.String                                                   `tfsdk:"team_id" query:"teamId,optional"`
	MaxItems types.Int64                                                    `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[WorkflowRunsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *WorkflowRunsDataSourceModel) toListParams(_ context.Context) (params serval.WorkflowRunListParams, diags diag.Diagnostics) {
	params = serval.WorkflowRunListParams{}

	if !m.TeamID.IsNull() {
		params.TeamID = param.NewOpt(m.TeamID.ValueString())
	}

	return
}

type WorkflowRunsItemsDataSourceModel struct {
	ID                types.String `tfsdk:"id" json:"id,computed"`
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
