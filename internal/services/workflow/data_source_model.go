// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflow

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkflowDataDataSourceEnvelope struct {
	Data WorkflowDataSourceModel `json:"data,computed"`
}

type WorkflowDataSourceModel struct {
	ID                      types.String `tfsdk:"id" path:"id,required"`
	Content                 types.String `tfsdk:"content" json:"content,computed"`
	Description             types.String `tfsdk:"description" json:"description,computed"`
	ExecutionScope          types.String `tfsdk:"execution_scope" json:"executionScope,computed"`
	IsTemporary             types.Bool   `tfsdk:"is_temporary" json:"isTemporary,computed"`
	Name                    types.String `tfsdk:"name" json:"name,computed"`
	Parameters              types.String `tfsdk:"parameters" json:"parameters,computed"`
	RequireFormConfirmation types.Bool   `tfsdk:"require_form_confirmation" json:"requireFormConfirmation,computed"`
	TeamID                  types.String `tfsdk:"team_id" json:"teamId,computed"`
	Type                    types.String `tfsdk:"type" json:"type,computed"`
}
