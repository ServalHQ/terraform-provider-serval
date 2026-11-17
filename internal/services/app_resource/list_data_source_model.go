// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_resource

import (
	"context"

	flowapi "github.com/ServalHQ/terraform-provider-serval/internal/gen/proto/svflow/publicapi"
	"google.golang.org/protobuf/proto"

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

func (m *AppResourcesDataSourceModel) toListParams(_ context.Context) (params proto.Message, diags diag.Diagnostics) {
	req := &flowapi.ListAppResourcesRequest{}

	if !m.AppInstanceID.IsNull() {
		req.AppInstanceId = proto.String(m.AppInstanceID.ValueString())
	}
	if !m.TeamID.IsNull() {
		req.TeamId = proto.String(m.TeamID.ValueString())
	}

	params = req
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
