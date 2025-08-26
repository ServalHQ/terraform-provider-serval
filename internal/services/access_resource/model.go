// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_resource

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/serval-terraform/internal/apijson"
)

type AccessResourceDataEnvelope struct {
	Data AccessResourceModel `json:"data"`
}

type AccessResourceModel struct {
	ID            types.String `tfsdk:"id" json:"id,computed"`
	AppInstanceID types.String `tfsdk:"app_instance_id" json:"appInstanceId,optional"`
	Description   types.String `tfsdk:"description" json:"description,optional"`
	ExternalID    types.String `tfsdk:"external_id" json:"externalId,optional"`
	Name          types.String `tfsdk:"name" json:"name,optional"`
	ResourceType  types.String `tfsdk:"resource_type" json:"resourceType,optional"`
}

func (m AccessResourceModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AccessResourceModel) MarshalJSONForUpdate(state AccessResourceModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
