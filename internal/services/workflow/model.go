// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflow

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/serval-terraform/internal/apijson"
)

type WorkflowDataEnvelope struct {
	Data WorkflowModel `json:"data"`
}

type WorkflowModel struct {
	ID                      types.String `tfsdk:"id" json:"id,computed"`
	TeamID                  types.String `tfsdk:"team_id" json:"teamId,optional"`
	Content                 types.String `tfsdk:"content" json:"content,optional"`
	Description             types.String `tfsdk:"description" json:"description,optional"`
	ExecutionScope          types.String `tfsdk:"execution_scope" json:"executionScope,optional"`
	IsTemporary             types.Bool   `tfsdk:"is_temporary" json:"isTemporary,optional"`
	Name                    types.String `tfsdk:"name" json:"name,optional"`
	Parameters              types.String `tfsdk:"parameters" json:"parameters,optional"`
	RequireFormConfirmation types.Bool   `tfsdk:"require_form_confirmation" json:"requireFormConfirmation,optional"`
	Type                    types.String `tfsdk:"type" json:"type,optional"`
}

func (m WorkflowModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m WorkflowModel) MarshalJSONForUpdate(state WorkflowModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
