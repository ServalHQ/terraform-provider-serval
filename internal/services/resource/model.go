// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package resource

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/serval-terraform/internal/apijson"
)

type ResourceDataEnvelope struct {
	Data ResourceModel `json:"data"`
}

type ResourceModel struct {
	ID            types.String `tfsdk:"id" json:"id,computed"`
	AppInstanceID types.String `tfsdk:"app_instance_id" json:"appInstanceId,optional"`
	Description   types.String `tfsdk:"description" json:"description,optional"`
	ExternalID    types.String `tfsdk:"external_id" json:"externalId,optional"`
	Name          types.String `tfsdk:"name" json:"name,optional"`
	ResourceType  types.String `tfsdk:"resource_type" json:"resourceType,optional"`
}

func (m ResourceModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ResourceModel) MarshalJSONForUpdate(state ResourceModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
