// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy_approval_procedure

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/serval-terraform/internal/apijson"
)

type AccessPolicyApprovalProcedureDataEnvelope struct {
	Data AccessPolicyApprovalProcedureModel `json:"data"`
}

type AccessPolicyApprovalProcedureModel struct {
	ID             types.String                                `tfsdk:"id" json:"id,computed"`
	AccessPolicyID types.String                                `tfsdk:"access_policy_id" path:"access_policy_id,required"`
	Steps          *[]*AccessPolicyApprovalProcedureStepsModel `tfsdk:"steps" json:"steps,optional"`
}

func (m AccessPolicyApprovalProcedureModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AccessPolicyApprovalProcedureModel) MarshalJSONForUpdate(state AccessPolicyApprovalProcedureModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type AccessPolicyApprovalProcedureStepsModel struct {
	ID                types.String    `tfsdk:"id" json:"id,optional"`
	AllowSelfApproval types.Bool      `tfsdk:"allow_self_approval" json:"allowSelfApproval,optional"`
	ServalGroupIDs    *[]types.String `tfsdk:"serval_group_ids" json:"servalGroupIds,optional"`
	SpecificUserIDs   *[]types.String `tfsdk:"specific_user_ids" json:"specificUserIds,optional"`
	StepType          types.String    `tfsdk:"step_type" json:"stepType,optional"`
}
