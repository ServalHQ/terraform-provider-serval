// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflow_approval_procedure

import (
	"context"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
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
	ID                types.String                   `tfsdk:"id" json:"id,computed"`
	AllowSelfApproval types.Bool                     `tfsdk:"allow_self_approval" json:"allowSelfApproval,computed"`
	ServalGroupIDs    customfield.List[types.String] `tfsdk:"serval_group_ids" json:"servalGroupIds,computed"`
	SpecificUserIDs   customfield.List[types.String] `tfsdk:"specific_user_ids" json:"specificUserIds,computed"`
	StepType          types.String                   `tfsdk:"step_type" json:"stepType,computed"`
}
