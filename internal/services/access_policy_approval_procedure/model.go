// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy_approval_procedure

import (
	"github.com/ServalHQ/terraform-provider-serval/internal/apijson"
	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessPolicyApprovalProcedureDataEnvelope struct {
	Data AccessPolicyApprovalProcedureModel `json:"data"`
}

type AccessPolicyApprovalProcedureModel struct {
	ID             types.String                                                          `tfsdk:"id" json:"id,computed"`
	AccessPolicyID types.String                                                          `tfsdk:"access_policy_id" path:"access_policy_id,required"`
	Steps          customfield.NestedObjectList[AccessPolicyApprovalProcedureStepsModel] `tfsdk:"steps" json:"steps,computed_optional"`
}

func (m AccessPolicyApprovalProcedureModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AccessPolicyApprovalProcedureModel) MarshalJSONForUpdate(state AccessPolicyApprovalProcedureModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type AccessPolicyApprovalProcedureStepsModel struct {
	ID                types.String                                           `tfsdk:"id" json:"id,computed"`
	AllowSelfApproval types.Bool                                             `tfsdk:"allow_self_approval" json:"allowSelfApproval,computed_optional"`
	Approvers         *[]*AccessPolicyApprovalProcedureStepsApproversModel   `tfsdk:"approvers" json:"approvers,optional"`
	CustomWorkflow    *AccessPolicyApprovalProcedureStepsCustomWorkflowModel `tfsdk:"custom_workflow" json:"customWorkflow,optional"`
}

type AccessPolicyApprovalProcedureStepsApproversModel struct {
	AppOwner jsontypes.Normalized                                   `tfsdk:"app_owner" json:"appOwner,optional"`
	Notify   types.Bool                                             `tfsdk:"notify" json:"notify,computed_optional"`
	Group    *AccessPolicyApprovalProcedureStepsApproversGroupModel `tfsdk:"group" json:"group,optional"`
	Manager  jsontypes.Normalized                                   `tfsdk:"manager" json:"manager,optional"`
	User     *AccessPolicyApprovalProcedureStepsApproversUserModel  `tfsdk:"user" json:"user,optional"`
}

type AccessPolicyApprovalProcedureStepsApproversGroupModel struct {
	GroupID types.String `tfsdk:"group_id" json:"groupId,optional"`
}

type AccessPolicyApprovalProcedureStepsApproversUserModel struct {
	UserID types.String `tfsdk:"user_id" json:"userId,optional"`
}

type AccessPolicyApprovalProcedureStepsCustomWorkflowModel struct {
	WorkflowID types.String `tfsdk:"workflow_id" json:"workflowId,optional"`
}
