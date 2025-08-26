// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package entitlement

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EntitlementDataDataSourceEnvelope struct {
	Data EntitlementDataSourceModel `json:"data,computed"`
}

type EntitlementDataSourceModel struct {
	ID                 types.String `tfsdk:"id" path:"id,required"`
	AccessPolicyID     types.String `tfsdk:"access_policy_id" json:"accessPolicyId,computed"`
	Description        types.String `tfsdk:"description" json:"description,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
	ProvisioningMethod types.String `tfsdk:"provisioning_method" json:"provisioningMethod,computed"`
	RequestsEnabled    types.Bool   `tfsdk:"requests_enabled" json:"requestsEnabled,computed"`
	ResourceID         types.String `tfsdk:"resource_id" json:"resourceId,computed"`
}
