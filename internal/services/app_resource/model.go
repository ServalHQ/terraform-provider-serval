// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_resource

import (
	"github.com/ServalHQ/terraform-provider-serval/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AppResourceDataEnvelope struct {
	Data AppResourceModel `json:"data"`
}

type AppResourceModel struct {
	ID            types.String `tfsdk:"id" json:"id,computed"`
	AppInstanceID types.String `tfsdk:"app_instance_id" json:"appInstanceId,required"`
	Description   types.String `tfsdk:"description" json:"description,optional"`
	ExternalID    types.String `tfsdk:"external_id" json:"externalId,optional"`
	Name          types.String `tfsdk:"name" json:"name,optional"`
	ResourceType  types.String `tfsdk:"resource_type" json:"resourceType,optional"`
}

func (m AppResourceModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AppResourceModel) MarshalJSONForUpdate(state AppResourceModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
