// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_instance

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AppInstanceDataDataSourceEnvelope struct {
	Data AppInstanceDataSourceModel `json:"data,computed"`
}

type AppInstanceDataSourceModel struct {
	ID                        types.String `tfsdk:"id" path:"id,required"`
	AccessRequestsEnabled     types.Bool   `tfsdk:"access_requests_enabled" json:"accessRequestsEnabled,computed"`
	CustomServiceID           types.String `tfsdk:"custom_service_id" json:"customServiceId,computed"`
	DefaultAccessPolicyID     types.String `tfsdk:"default_access_policy_id" json:"defaultAccessPolicyId,computed"`
	ExternalServiceInstanceID types.String `tfsdk:"external_service_instance_id" json:"externalServiceInstanceId,computed"`
	Name                      types.String `tfsdk:"name" json:"name,computed"`
	Service                   types.String `tfsdk:"service" json:"service,computed"`
	TeamID                    types.String `tfsdk:"team_id" json:"teamId,computed"`
}
