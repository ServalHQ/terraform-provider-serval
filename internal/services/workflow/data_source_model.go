// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflow

import (
	"context"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/serval-go/packages/param"
	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkflowDataDataSourceEnvelope struct {
	Data WorkflowDataSourceModel `json:"data,computed"`
}

type WorkflowDataSourceModel struct {
	ID                      types.String                      `tfsdk:"id" path:"id,computed_optional"`
	Content                 types.String                      `tfsdk:"content" json:"content,computed"`
	Description             types.String                      `tfsdk:"description" json:"description,computed"`
	ExecutionScope          types.String                      `tfsdk:"execution_scope" json:"executionScope,computed"`
	HasUnpublishedChanges   types.Bool                        `tfsdk:"has_unpublished_changes" json:"hasUnpublishedChanges,computed"`
	IsPublished             types.Bool                        `tfsdk:"is_published" json:"isPublished,computed"`
	Name                    types.String                      `tfsdk:"name" json:"name,computed"`
	RequireFormConfirmation types.Bool                        `tfsdk:"require_form_confirmation" json:"requireFormConfirmation,computed"`
	TeamID                  types.String                      `tfsdk:"team_id" json:"teamId,computed"`
	TagIDs                  customfield.List[types.String]    `tfsdk:"tag_ids" json:"tagIds,computed"`
	FindOneBy               *WorkflowFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

func (m *WorkflowDataSourceModel) toListParams(_ context.Context) (params serval.WorkflowListParams, diags diag.Diagnostics) {
	params = serval.WorkflowListParams{}

	if !m.FindOneBy.TeamID.IsNull() {
		params.TeamID = param.NewOpt(m.FindOneBy.TeamID.ValueString())
	}

	return
}

type WorkflowFindOneByDataSourceModel struct {
	TeamID types.String `tfsdk:"team_id" query:"teamId,optional"`
}
