// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_resource_entitlement

import (
	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AppResourceEntitlementDataDataSourceEnvelope struct {
	Data AppResourceEntitlementDataSourceModel `json:"data,computed"`
}

type AppResourceEntitlementDataSourceModel struct {
	ID                 types.String                                                                      `tfsdk:"id" path:"id,required"`
	AccessPolicyID     types.String                                                                      `tfsdk:"access_policy_id" json:"accessPolicyId,computed"`
	Description        types.String                                                                      `tfsdk:"description" json:"description,computed"`
	Name               types.String                                                                      `tfsdk:"name" json:"name,computed"`
	RequestsEnabled    types.Bool                                                                        `tfsdk:"requests_enabled" json:"requestsEnabled,computed"`
	ResourceID         types.String                                                                      `tfsdk:"resource_id" json:"resourceId,computed"`
	ProvisioningMethod customfield.NestedObject[AppResourceEntitlementProvisioningMethodDataSourceModel] `tfsdk:"provisioning_method" json:"provisioningMethod,computed"`
}

type AppResourceEntitlementProvisioningMethodDataSourceModel struct {
	BuiltinWorkflow    jsontypes.Normalized                                                                                `tfsdk:"builtin_workflow" json:"builtinWorkflow,computed"`
	CustomWorkflow     customfield.NestedObject[AppResourceEntitlementProvisioningMethodCustomWorkflowDataSourceModel]     `tfsdk:"custom_workflow" json:"customWorkflow,computed"`
	LinkedEntitlements customfield.NestedObject[AppResourceEntitlementProvisioningMethodLinkedEntitlementsDataSourceModel] `tfsdk:"linked_entitlements" json:"linkedEntitlements,computed"`
	Manual             customfield.NestedObject[AppResourceEntitlementProvisioningMethodManualDataSourceModel]             `tfsdk:"manual" json:"manual,computed"`
}

type AppResourceEntitlementProvisioningMethodCustomWorkflowDataSourceModel struct {
	DeprovisionWorkflowID types.String `tfsdk:"deprovision_workflow_id" json:"deprovisionWorkflowId,computed"`
	ProvisionWorkflowID   types.String `tfsdk:"provision_workflow_id" json:"provisionWorkflowId,computed"`
}

type AppResourceEntitlementProvisioningMethodLinkedEntitlementsDataSourceModel struct {
	LinkedEntitlementIDs customfield.List[types.String] `tfsdk:"linked_entitlement_ids" json:"linkedEntitlementIds,computed"`
}

type AppResourceEntitlementProvisioningMethodManualDataSourceModel struct {
	Assignees customfield.NestedObjectList[AppResourceEntitlementProvisioningMethodManualAssigneesDataSourceModel] `tfsdk:"assignees" json:"assignees,computed"`
}

type AppResourceEntitlementProvisioningMethodManualAssigneesDataSourceModel struct {
	AssigneeID   types.String `tfsdk:"assignee_id" json:"assigneeId,computed"`
	AssigneeType types.String `tfsdk:"assignee_type" json:"assigneeType,computed"`
}
