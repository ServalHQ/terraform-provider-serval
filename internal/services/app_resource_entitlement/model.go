// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_resource_entitlement

import (
	"github.com/ServalHQ/terraform-provider-serval/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AppResourceEntitlementDataEnvelope struct {
	Data AppResourceEntitlementModel `json:"data"`
}

type AppResourceEntitlementModel struct {
	ID                 types.String                                   `tfsdk:"id" json:"id,computed"`
	ResourceID         types.String                                   `tfsdk:"resource_id" json:"resourceId,optional"`
	AccessPolicyID     types.String                                   `tfsdk:"access_policy_id" json:"accessPolicyId,optional"`
	Description        types.String                                   `tfsdk:"description" json:"description,optional"`
	ExternalData       types.String                                   `tfsdk:"external_data" json:"externalData,optional"`
	ExternalID         types.String                                   `tfsdk:"external_id" json:"externalId,optional"`
	Name               types.String                                   `tfsdk:"name" json:"name,optional"`
	RequestsEnabled    types.Bool                                     `tfsdk:"requests_enabled" json:"requestsEnabled,optional"`
	ProvisioningMethod *AppResourceEntitlementProvisioningMethodModel `tfsdk:"provisioning_method" json:"provisioningMethod,optional"`
}

func (m AppResourceEntitlementModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AppResourceEntitlementModel) MarshalJSONForUpdate(state AppResourceEntitlementModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type AppResourceEntitlementProvisioningMethodModel struct {
	BuiltinWorkflow    jsontypes.Normalized                                             `tfsdk:"builtin_workflow" json:"builtinWorkflow,optional"`
	CustomWorkflow     *AppResourceEntitlementProvisioningMethodCustomWorkflowModel     `tfsdk:"custom_workflow" json:"customWorkflow,optional"`
	LinkedEntitlements *AppResourceEntitlementProvisioningMethodLinkedEntitlementsModel `tfsdk:"linked_entitlements" json:"linkedEntitlements,optional"`
	Manual             *AppResourceEntitlementProvisioningMethodManualModel             `tfsdk:"manual" json:"manual,optional"`
}

type AppResourceEntitlementProvisioningMethodCustomWorkflowModel struct {
	DeprovisionWorkflowID types.String `tfsdk:"deprovision_workflow_id" json:"deprovisionWorkflowId,optional"`
	ProvisionWorkflowID   types.String `tfsdk:"provision_workflow_id" json:"provisionWorkflowId,optional"`
}

type AppResourceEntitlementProvisioningMethodLinkedEntitlementsModel struct {
	LinkedEntitlementIDs *[]types.String `tfsdk:"linked_entitlement_ids" json:"linkedEntitlementIds,optional"`
}

type AppResourceEntitlementProvisioningMethodManualModel struct {
	Assignees *[]*AppResourceEntitlementProvisioningMethodManualAssigneesModel `tfsdk:"assignees" json:"assignees,optional"`
}

type AppResourceEntitlementProvisioningMethodManualAssigneesModel struct {
	AssigneeID   types.String `tfsdk:"assignee_id" json:"assigneeId,optional"`
	AssigneeType types.String `tfsdk:"assignee_type" json:"assigneeType,optional"`
}
