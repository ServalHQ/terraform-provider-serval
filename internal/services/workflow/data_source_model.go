// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflow

import (
	"context"

	"github.com/ServalHQ/terraform-provider-serval/internal/client"
	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	flowapi "github.com/ServalHQ/terraform-provider-serval/internal/gen/proto/svflow/publicapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"google.golang.org/protobuf/proto"
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

func (m *WorkflowDataSourceModel) toListParams(_ context.Context) (params proto.Message, diags diag.Diagnostics) {
	req := &flowapi.ListWorkflowsRequest{}
	var explicit []string

	if !m.FindOneBy.TeamID.IsNull() {
		req.TeamId = m.FindOneBy.TeamID.ValueString()
		explicit = append(explicit, "teamId")
	}

	params = client.WithExplicit(req, explicit...)
	return
}

type WorkflowFindOneByDataSourceModel struct {
	TeamID types.String `tfsdk:"team_id" query:"teamId,optional"`
}
