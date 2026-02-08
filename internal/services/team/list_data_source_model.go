// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package team

import (
	"context"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TeamsDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[TeamsItemsDataSourceModel] `json:"data,computed"`
}

type TeamsDataSourceModel struct {
	MaxItems types.Int64                                             `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[TeamsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *TeamsDataSourceModel) toListParams(_ context.Context) (params serval.TeamListParams, diags diag.Diagnostics) {
	params = serval.TeamListParams{}

	return
}

type TeamsItemsDataSourceModel struct {
	ID             types.String      `tfsdk:"id" json:"id,computed"`
	CreatedAt      timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	Description    types.String      `tfsdk:"description" json:"description,computed"`
	Name           types.String      `tfsdk:"name" json:"name,computed"`
	OrganizationID types.String      `tfsdk:"organization_id" json:"organizationId,computed"`
	Prefix         types.String      `tfsdk:"prefix" json:"prefix,computed"`
}
