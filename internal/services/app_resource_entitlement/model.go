// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_resource_entitlement

import (
	"github.com/ServalHQ/terraform-provider-serval/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AppResourceEntitlementDataEnvelope struct {
	Data AppResourceEntitlementModel `json:"data"`
}

type AppResourceEntitlementModel struct {
	ID                          types.String                                               `tfsdk:"id" json:"id,computed"`
	ResourceID                  types.String                                               `tfsdk:"resource_id" json:"resourceId,optional"`
	AccessPolicyID              types.String                                               `tfsdk:"access_policy_id" json:"accessPolicyId,optional"`
	Description                 types.String                                               `tfsdk:"description" json:"description,optional"`
	Name                        types.String                                               `tfsdk:"name" json:"name,optional"`
	ProvisioningMethod          types.String                                               `tfsdk:"provisioning_method" json:"provisioningMethod,optional"`
	RequestsEnabled             types.Bool                                                 `tfsdk:"requests_enabled" json:"requestsEnabled,optional"`
	LinkedEntitlementIDs        *[]types.String                                            `tfsdk:"linked_entitlement_ids" json:"linkedEntitlementIds,optional"`
	ManualProvisioningAssignees *[]*AppResourceEntitlementManualProvisioningAssigneesModel `tfsdk:"manual_provisioning_assignees" json:"manualProvisioningAssignees,optional"`
}

func (m AppResourceEntitlementModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AppResourceEntitlementModel) MarshalJSONForUpdate(state AppResourceEntitlementModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type AppResourceEntitlementManualProvisioningAssigneesModel struct {
	AssigneeID   types.String `tfsdk:"assignee_id" json:"assigneeId,optional"`
	AssigneeType types.String `tfsdk:"assignee_type" json:"assigneeType,optional"`
}
