// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflow_approval_procedure

import (
	"github.com/ServalHQ/terraform-provider-serval/internal/apijson"
	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
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
	ID                types.String    `tfsdk:"id" json:"id,computed"`
	AllowSelfApproval types.Bool      `tfsdk:"allow_self_approval" json:"allowSelfApproval,optional"`
	CustomWorkflowID  types.String    `tfsdk:"custom_workflow_id" json:"customWorkflowId,optional"`
	ServalGroupIDs    *[]types.String `tfsdk:"serval_group_ids" json:"servalGroupIds,optional"`
	SpecificUserIDs   *[]types.String `tfsdk:"specific_user_ids" json:"specificUserIds,optional"`
}
