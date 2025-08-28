// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_instance

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/serval-terraform/internal/apijson"
)

type AppInstanceDataEnvelope struct {
	Data AppInstanceModel `json:"data"`
}

type AppInstanceModel struct {
	ID                    types.String `tfsdk:"id" json:"id,computed"`
	Service               types.String `tfsdk:"service" json:"service,optional"`
	TeamID                types.String `tfsdk:"team_id" json:"teamId,optional"`
	AccessRequestsEnabled types.Bool   `tfsdk:"access_requests_enabled" json:"accessRequestsEnabled,optional"`
	DefaultAccessPolicyID types.String `tfsdk:"default_access_policy_id" json:"defaultAccessPolicyId,optional"`
	InstanceID            types.String `tfsdk:"instance_id" json:"instanceId,optional"`
	Name                  types.String `tfsdk:"name" json:"name,optional"`
}

func (m AppInstanceModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AppInstanceModel) MarshalJSONForUpdate(state AppInstanceModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
