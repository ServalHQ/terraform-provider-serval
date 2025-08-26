// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_resource

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessResourceDataDataSourceEnvelope struct {
	Data AccessResourceDataSourceModel `json:"data,computed"`
}

type AccessResourceDataSourceModel struct {
	ID            types.String `tfsdk:"id" path:"id,required"`
	AppInstanceID types.String `tfsdk:"app_instance_id" json:"appInstanceId,computed"`
	Description   types.String `tfsdk:"description" json:"description,computed"`
	ExternalID    types.String `tfsdk:"external_id" json:"externalId,computed"`
	Name          types.String `tfsdk:"name" json:"name,computed"`
	ResourceType  types.String `tfsdk:"resource_type" json:"resourceType,computed"`
}
