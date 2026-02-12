// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package group

import (
	"context"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type GroupsDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[GroupsItemsDataSourceModel] `json:"data,computed"`
}

type GroupsDataSourceModel struct {
	MaxItems types.Int64                                              `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[GroupsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *GroupsDataSourceModel) toListParams(_ context.Context) (params serval.GroupListParams, diags diag.Diagnostics) {
	params = serval.GroupListParams{}

	return
}

type GroupsItemsDataSourceModel struct {
	ID             types.String                   `tfsdk:"id" json:"id,computed"`
	CreatedAt      timetypes.RFC3339              `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	DeletedAt      timetypes.RFC3339              `tfsdk:"deleted_at" json:"deletedAt,computed" format:"date-time"`
	Name           types.String                   `tfsdk:"name" json:"name,computed"`
	OrganizationID types.String                   `tfsdk:"organization_id" json:"organizationId,computed"`
	UserIDs        customfield.List[types.String] `tfsdk:"user_ids" json:"userIds,computed"`
}
