// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflow_approval_procedure

import (
	"context"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkflowApprovalProcedureDataDataSourceEnvelope struct {
	Data WorkflowApprovalProcedureDataSourceModel `json:"data,computed"`
}

type WorkflowApprovalProcedureDataSourceModel struct {
	ID         types.String                                                                `tfsdk:"id" path:"id,required"`
	WorkflowID types.String                                                                `tfsdk:"workflow_id" path:"workflow_id,required"`
	Steps      customfield.NestedObjectList[WorkflowApprovalProcedureStepsDataSourceModel] `tfsdk:"steps" json:"steps,computed"`
}

func (m *WorkflowApprovalProcedureDataSourceModel) toReadParams(_ context.Context) (params serval.WorkflowApprovalProcedureGetParams, diags diag.Diagnostics) {
	params = serval.WorkflowApprovalProcedureGetParams{
		WorkflowID: m.WorkflowID.ValueString(),
	}

	return
}

type WorkflowApprovalProcedureStepsDataSourceModel struct {
	ID                types.String                                                                          `tfsdk:"id" json:"id,computed"`
	AllowSelfApproval types.Bool                                                                            `tfsdk:"allow_self_approval" json:"allowSelfApproval,computed"`
	Approvers         customfield.NestedObjectList[WorkflowApprovalProcedureStepsApproversDataSourceModel]  `tfsdk:"approvers" json:"approvers,computed"`
	CustomWorkflow    customfield.NestedObject[WorkflowApprovalProcedureStepsCustomWorkflowDataSourceModel] `tfsdk:"custom_workflow" json:"customWorkflow,computed"`
}

type WorkflowApprovalProcedureStepsApproversDataSourceModel struct {
	AppOwner jsontypes.Normalized                                                                  `tfsdk:"app_owner" json:"appOwner,computed"`
	Notify   types.Bool                                                                            `tfsdk:"notify" json:"notify,computed"`
	Group    customfield.NestedObject[WorkflowApprovalProcedureStepsApproversGroupDataSourceModel] `tfsdk:"group" json:"group,computed"`
	Manager  jsontypes.Normalized                                                                  `tfsdk:"manager" json:"manager,computed"`
	User     customfield.NestedObject[WorkflowApprovalProcedureStepsApproversUserDataSourceModel]  `tfsdk:"user" json:"user,computed"`
}

type WorkflowApprovalProcedureStepsApproversGroupDataSourceModel struct {
	GroupID types.String `tfsdk:"group_id" json:"groupId,computed"`
}

type WorkflowApprovalProcedureStepsApproversUserDataSourceModel struct {
	UserID types.String `tfsdk:"user_id" json:"userId,computed"`
}

type WorkflowApprovalProcedureStepsCustomWorkflowDataSourceModel struct {
	WorkflowID types.String `tfsdk:"workflow_id" json:"workflowId,computed"`
}
