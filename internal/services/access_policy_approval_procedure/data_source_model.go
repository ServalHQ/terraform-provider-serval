// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy_approval_procedure

import (
	"context"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessPolicyApprovalProcedureDataDataSourceEnvelope struct {
	Data AccessPolicyApprovalProcedureDataSourceModel `json:"data,computed"`
}

type AccessPolicyApprovalProcedureDataSourceModel struct {
	AccessPolicyID types.String                                                                    `tfsdk:"access_policy_id" path:"access_policy_id,required"`
	ID             types.String                                                                    `tfsdk:"id" path:"id,required"`
	Steps          customfield.NestedObjectList[AccessPolicyApprovalProcedureStepsDataSourceModel] `tfsdk:"steps" json:"steps,computed"`
}

func (m *AccessPolicyApprovalProcedureDataSourceModel) toReadParams(_ context.Context) (params serval.AccessPolicyApprovalProcedureGetParams, diags diag.Diagnostics) {
	params = serval.AccessPolicyApprovalProcedureGetParams{
		AccessPolicyID: m.AccessPolicyID.ValueString(),
	}

	return
}

type AccessPolicyApprovalProcedureStepsDataSourceModel struct {
	ID                types.String                                                                              `tfsdk:"id" json:"id,computed"`
	AllowSelfApproval types.Bool                                                                                `tfsdk:"allow_self_approval" json:"allowSelfApproval,computed"`
	Approvers         customfield.NestedObjectList[AccessPolicyApprovalProcedureStepsApproversDataSourceModel]  `tfsdk:"approvers" json:"approvers,computed"`
	CustomWorkflow    customfield.NestedObject[AccessPolicyApprovalProcedureStepsCustomWorkflowDataSourceModel] `tfsdk:"custom_workflow" json:"customWorkflow,computed"`
}

type AccessPolicyApprovalProcedureStepsApproversDataSourceModel struct {
	AppOwner jsontypes.Normalized                                                                      `tfsdk:"app_owner" json:"appOwner,computed"`
	Notify   types.Bool                                                                                `tfsdk:"notify" json:"notify,computed"`
	Group    customfield.NestedObject[AccessPolicyApprovalProcedureStepsApproversGroupDataSourceModel] `tfsdk:"group" json:"group,computed"`
	Manager  jsontypes.Normalized                                                                      `tfsdk:"manager" json:"manager,computed"`
	User     customfield.NestedObject[AccessPolicyApprovalProcedureStepsApproversUserDataSourceModel]  `tfsdk:"user" json:"user,computed"`
}

type AccessPolicyApprovalProcedureStepsApproversGroupDataSourceModel struct {
	GroupID types.String `tfsdk:"group_id" json:"groupId,computed"`
}

type AccessPolicyApprovalProcedureStepsApproversUserDataSourceModel struct {
	UserID types.String `tfsdk:"user_id" json:"userId,computed"`
}

type AccessPolicyApprovalProcedureStepsCustomWorkflowDataSourceModel struct {
	WorkflowID types.String `tfsdk:"workflow_id" json:"workflowId,computed"`
}
