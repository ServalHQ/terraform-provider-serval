// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_resource

import (
	"context"

	flowapi "github.com/ServalHQ/terraform-provider-serval/internal/gen/proto/svflow/publicapi"
	"google.golang.org/protobuf/proto"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AppResourceDataDataSourceEnvelope struct {
	Data AppResourceDataSourceModel `json:"data,computed"`
}

type AppResourceDataSourceModel struct {
	ID            types.String                         `tfsdk:"id" path:"id,computed_optional"`
	AppInstanceID types.String                         `tfsdk:"app_instance_id" json:"appInstanceId,computed"`
	Description   types.String                         `tfsdk:"description" json:"description,computed"`
	ExternalID    types.String                         `tfsdk:"external_id" json:"externalId,computed"`
	Name          types.String                         `tfsdk:"name" json:"name,computed"`
	ResourceType  types.String                         `tfsdk:"resource_type" json:"resourceType,computed"`
	FindOneBy     *AppResourceFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

func (m *AppResourceDataSourceModel) toListParams(_ context.Context) (params proto.Message, diags diag.Diagnostics) {
	req := &flowapi.ListAppResourcesRequest{}

	if !m.FindOneBy.AppInstanceID.IsNull() {
		req.AppInstanceId = proto.String(m.FindOneBy.AppInstanceID.ValueString())
	}
	if !m.FindOneBy.TeamID.IsNull() {
		req.TeamId = proto.String(m.FindOneBy.TeamID.ValueString())
	}

	params = req
	return
}

type AppResourceFindOneByDataSourceModel struct {
	AppInstanceID types.String `tfsdk:"app_instance_id" query:"appInstanceId,optional"`
	TeamID        types.String `tfsdk:"team_id" query:"teamId,optional"`
}
