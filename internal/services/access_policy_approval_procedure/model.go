// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy_approval_procedure

import (
	"github.com/ServalHQ/terraform-provider-serval/internal/apijson"
	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessPolicyApprovalProcedureDataEnvelope struct {
	Data AccessPolicyApprovalProcedureModel `json:"data"`
}

type AccessPolicyApprovalProcedureModel struct {
	ID             types.String                                                          `tfsdk:"id" json:"id,computed"`
	AccessPolicyID types.String                                                          `tfsdk:"access_policy_id" json:"accessPolicyId" path:"access_policy_id,required"`
	Steps          customfield.NestedObjectList[AccessPolicyApprovalProcedureStepsModel] `tfsdk:"steps" json:"steps,computed_optional"`
}

func (m AccessPolicyApprovalProcedureModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AccessPolicyApprovalProcedureModel) MarshalJSONForUpdate(state AccessPolicyApprovalProcedureModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type AccessPolicyApprovalProcedureStepsModel struct {
	ID                types.String    `tfsdk:"id" json:"id,computed"`
	AllowSelfApproval types.Bool      `tfsdk:"allow_self_approval" json:"allowSelfApproval,computed_optional"`
	CustomWorkflowID  types.String    `tfsdk:"custom_workflow_id" json:"customWorkflowId,optional"`
	ServalGroupIDs    *[]types.String `tfsdk:"serval_group_ids" json:"servalGroupIds,optional"`
	SpecificUserIDs   *[]types.String `tfsdk:"specific_user_ids" json:"specificUserIds,optional"`
}
