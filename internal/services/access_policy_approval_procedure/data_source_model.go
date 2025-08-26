// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy_approval_procedure

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/serval-go"
	"github.com/stainless-sdks/serval-terraform/internal/customfield"
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
	ID                types.String                   `tfsdk:"id" json:"id,computed"`
	AllowSelfApproval types.Bool                     `tfsdk:"allow_self_approval" json:"allowSelfApproval,computed"`
	ServalGroupIDs    customfield.List[types.String] `tfsdk:"serval_group_ids" json:"servalGroupIds,computed"`
	SpecificUserIDs   customfield.List[types.String] `tfsdk:"specific_user_ids" json:"specificUserIds,computed"`
	StepType          types.String                   `tfsdk:"step_type" json:"stepType,computed"`
}
