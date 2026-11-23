// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_instance

import (
	"context"

	"github.com/ServalHQ/terraform-provider-serval/internal/client"
	flowapi "github.com/ServalHQ/terraform-provider-serval/internal/gen/proto/svflow/publicapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"google.golang.org/protobuf/proto"
)

type AppInstanceDataDataSourceEnvelope struct {
	Data AppInstanceDataSourceModel `json:"data,computed"`
}

type AppInstanceDataSourceModel struct {
	ID                        types.String                         `tfsdk:"id" path:"id,computed_optional"`
	AccessRequestsEnabled     types.Bool                           `tfsdk:"access_requests_enabled" json:"accessRequestsEnabled,computed"`
	CustomServiceID           types.String                         `tfsdk:"custom_service_id" json:"customServiceId,computed"`
	DefaultAccessPolicyID     types.String                         `tfsdk:"default_access_policy_id" json:"defaultAccessPolicyId,computed"`
	ExternalServiceInstanceID types.String                         `tfsdk:"external_service_instance_id" json:"externalServiceInstanceId,computed"`
	Name                      types.String                         `tfsdk:"name" json:"name,computed"`
	Service                   types.String                         `tfsdk:"service" json:"service,computed"`
	TeamID                    types.String                         `tfsdk:"team_id" json:"teamId,computed"`
	FindOneBy                 *AppInstanceFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

func (m *AppInstanceDataSourceModel) toListParams(_ context.Context) (params proto.Message, diags diag.Diagnostics) {
	req := &flowapi.ListAppInstancesRequest{}
	var explicit []string

	if !m.FindOneBy.TeamID.IsNull() {
		req.TeamId = m.FindOneBy.TeamID.ValueString()
		explicit = append(explicit, "teamId")
	}

	params = client.WithExplicit(req, explicit...)
	return
}

type AppInstanceFindOneByDataSourceModel struct {
	TeamID types.String `tfsdk:"team_id" query:"teamId,optional"`
}
