// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_resource

import (
	"context"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/serval-go/packages/param"
	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AppResourcesDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[AppResourcesItemsDataSourceModel] `json:"data,computed"`
}

type AppResourcesDataSourceModel struct {
	AppInstanceID types.String                                                   `tfsdk:"app_instance_id" query:"appInstanceId,optional"`
	TeamID        types.String                                                   `tfsdk:"team_id" query:"teamId,optional"`
	MaxItems      types.Int64                                                    `tfsdk:"max_items"`
	Items         customfield.NestedObjectList[AppResourcesItemsDataSourceModel] `tfsdk:"items"`
}

func (m *AppResourcesDataSourceModel) toListParams(_ context.Context) (params serval.AppResourceListParams, diags diag.Diagnostics) {
	params = serval.AppResourceListParams{}

	if !m.AppInstanceID.IsNull() {
		params.AppInstanceID = param.NewOpt(m.AppInstanceID.ValueString())
	}
	if !m.TeamID.IsNull() {
		params.TeamID = param.NewOpt(m.TeamID.ValueString())
	}

	return
}

type AppResourcesItemsDataSourceModel struct {
	ID            types.String `tfsdk:"id" json:"id,computed"`
	AppInstanceID types.String `tfsdk:"app_instance_id" json:"appInstanceId,computed"`
	Description   types.String `tfsdk:"description" json:"description,computed"`
	ExternalID    types.String `tfsdk:"external_id" json:"externalId,computed"`
	Name          types.String `tfsdk:"name" json:"name,computed"`
	ResourceType  types.String `tfsdk:"resource_type" json:"resourceType,computed"`
}
