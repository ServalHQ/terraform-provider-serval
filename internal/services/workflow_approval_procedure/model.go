// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflow_approval_procedure

import (
	"github.com/ServalHQ/terraform-provider-serval/internal/apijson"
	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkflowApprovalProcedureDataEnvelope struct {
	Data WorkflowApprovalProcedureModel `json:"data"`
}

type WorkflowApprovalProcedureModel struct {
	ID         types.String                                                      `tfsdk:"id" json:"id,computed"`
	WorkflowID types.String                                                      `tfsdk:"workflow_id" path:"workflow_id,required"`
	Steps      customfield.NestedObjectList[WorkflowApprovalProcedureStepsModel] `tfsdk:"steps" json:"steps,computed_optional"`
}

func (m WorkflowApprovalProcedureModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m WorkflowApprovalProcedureModel) MarshalJSONForUpdate(state WorkflowApprovalProcedureModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type WorkflowApprovalProcedureStepsModel struct {
	ID                types.String                                       `tfsdk:"id" json:"id,computed"`
	AllowSelfApproval types.Bool                                         `tfsdk:"allow_self_approval" json:"allowSelfApproval,computed_optional"`
	Approvers         *[]*WorkflowApprovalProcedureStepsApproversModel   `tfsdk:"approvers" json:"approvers,optional"`
	CustomWorkflow    *WorkflowApprovalProcedureStepsCustomWorkflowModel `tfsdk:"custom_workflow" json:"customWorkflow,optional"`
}

type WorkflowApprovalProcedureStepsApproversModel struct {
	AppOwner jsontypes.Normalized                               `tfsdk:"app_owner" json:"appOwner,optional"`
	Notify   types.Bool                                         `tfsdk:"notify" json:"notify,computed_optional"`
	Group    *WorkflowApprovalProcedureStepsApproversGroupModel `tfsdk:"group" json:"group,optional"`
	Manager  jsontypes.Normalized                               `tfsdk:"manager" json:"manager,optional"`
	User     *WorkflowApprovalProcedureStepsApproversUserModel  `tfsdk:"user" json:"user,optional"`
}

type WorkflowApprovalProcedureStepsApproversGroupModel struct {
	GroupID types.String `tfsdk:"group_id" json:"groupId,optional"`
}

type WorkflowApprovalProcedureStepsApproversUserModel struct {
	UserID types.String `tfsdk:"user_id" json:"userId,optional"`
}

type WorkflowApprovalProcedureStepsCustomWorkflowModel struct {
	WorkflowID types.String `tfsdk:"workflow_id" json:"workflowId,optional"`
}
