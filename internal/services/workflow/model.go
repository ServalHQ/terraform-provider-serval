// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflow

import (
	"github.com/ServalHQ/terraform-provider-serval/internal/apijson"
	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkflowDataEnvelope struct {
	Data WorkflowModel `json:"data"`
}

type WorkflowModel struct {
	ID                      types.String                   `tfsdk:"id" json:"id,computed"`
	TeamID                  types.String                   `tfsdk:"team_id" json:"teamId,required"`
	Content                 types.String                   `tfsdk:"content" json:"content,required"`
	Name                    types.String                   `tfsdk:"name" json:"name,required"`
	Description             types.String                   `tfsdk:"description" json:"description,optional"`
	ExecutionScope          types.String                   `tfsdk:"execution_scope" json:"executionScope,optional"`
	IsPublished             types.Bool                     `tfsdk:"is_published" json:"isPublished,computed_optional"`
	RequireFormConfirmation types.Bool                     `tfsdk:"require_form_confirmation" json:"requireFormConfirmation,computed_optional"`
	HasUnpublishedChanges   types.Bool                     `tfsdk:"has_unpublished_changes" json:"hasUnpublishedChanges,computed"`
	TagIDs                  customfield.List[types.String] `tfsdk:"tag_ids" json:"tagIds,computed"`
}

func (m WorkflowModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m WorkflowModel) MarshalJSONForUpdate(state WorkflowModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
