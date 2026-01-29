// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_resource_role

import (
	"github.com/ServalHQ/terraform-provider-serval/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AppResourceRoleDataEnvelope struct {
	Data AppResourceRoleModel `json:"data"`
}

type AppResourceRoleModel struct {
	ID                 types.String                            `tfsdk:"id" json:"id,computed"`
	ResourceID         types.String                            `tfsdk:"resource_id" json:"resourceId,required"`
	Name               types.String                            `tfsdk:"name" json:"name,required"`
	ProvisioningMethod *AppResourceRoleProvisioningMethodModel `tfsdk:"provisioning_method" json:"provisioningMethod,required"`
	AccessPolicyID     types.String                            `tfsdk:"access_policy_id" json:"accessPolicyId,optional"`
	Description        types.String                            `tfsdk:"description" json:"description,optional"`
	ExternalData       types.String                            `tfsdk:"external_data" json:"externalData,optional"`
	ExternalID         types.String                            `tfsdk:"external_id" json:"externalId,optional"`
	RequestsEnabled    types.Bool                              `tfsdk:"requests_enabled" json:"requestsEnabled,computed_optional"`
}

func (m AppResourceRoleModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AppResourceRoleModel) MarshalJSONForUpdate(state AppResourceRoleModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type AppResourceRoleProvisioningMethodModel struct {
	BuiltinWorkflow jsontypes.Normalized                                  `tfsdk:"builtin_workflow" json:"builtinWorkflow,required"`
	CustomWorkflow  *AppResourceRoleProvisioningMethodCustomWorkflowModel `tfsdk:"custom_workflow" json:"customWorkflow,required"`
	LinkedRoles     *AppResourceRoleProvisioningMethodLinkedRolesModel    `tfsdk:"linked_roles" json:"linkedRoles,required"`
	Manual          *AppResourceRoleProvisioningMethodManualModel         `tfsdk:"manual" json:"manual,required"`
}

type AppResourceRoleProvisioningMethodCustomWorkflowModel struct {
	DeprovisionWorkflowID types.String `tfsdk:"deprovision_workflow_id" json:"deprovisionWorkflowId,optional"`
	ProvisionWorkflowID   types.String `tfsdk:"provision_workflow_id" json:"provisionWorkflowId,optional"`
}

type AppResourceRoleProvisioningMethodLinkedRolesModel struct {
	LinkedRoleIDs *[]types.String `tfsdk:"linked_role_ids" json:"linkedRoleIds,optional"`
}

type AppResourceRoleProvisioningMethodManualModel struct {
	Assignees *[]*AppResourceRoleProvisioningMethodManualAssigneesModel `tfsdk:"assignees" json:"assignees,optional"`
}

type AppResourceRoleProvisioningMethodManualAssigneesModel struct {
	AssigneeID   types.String `tfsdk:"assignee_id" json:"assigneeId,optional"`
	AssigneeType types.String `tfsdk:"assignee_type" json:"assigneeType,optional"`
}
