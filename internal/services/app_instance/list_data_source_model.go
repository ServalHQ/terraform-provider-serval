// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_instance

import (
	"context"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/serval-go/packages/param"
	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AppInstancesDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[AppInstancesItemsDataSourceModel] `json:"data,computed"`
}

type AppInstancesDataSourceModel struct {
	TeamID   types.String                                                   `tfsdk:"team_id" query:"teamId,optional"`
	MaxItems types.Int64                                                    `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[AppInstancesItemsDataSourceModel] `tfsdk:"items"`
}

func (m *AppInstancesDataSourceModel) toListParams(_ context.Context) (params serval.AppInstanceListParams, diags diag.Diagnostics) {
	params = serval.AppInstanceListParams{}

	if !m.TeamID.IsNull() {
		params.TeamID = param.NewOpt(m.TeamID.ValueString())
	}

	return
}

type AppInstancesItemsDataSourceModel struct {
	ID                        types.String `tfsdk:"id" json:"id,computed"`
	AccessRequestsEnabled     types.Bool   `tfsdk:"access_requests_enabled" json:"accessRequestsEnabled,computed"`
	CustomServiceID           types.String `tfsdk:"custom_service_id" json:"customServiceId,computed"`
	DefaultAccessPolicyID     types.String `tfsdk:"default_access_policy_id" json:"defaultAccessPolicyId,computed"`
	ExternalServiceInstanceID types.String `tfsdk:"external_service_instance_id" json:"externalServiceInstanceId,computed"`
	Name                      types.String `tfsdk:"name" json:"name,computed"`
	Service                   types.String `tfsdk:"service" json:"service,computed"`
	TeamID                    types.String `tfsdk:"team_id" json:"teamId,computed"`
}
