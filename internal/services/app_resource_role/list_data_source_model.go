// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_resource_role

import (
	"context"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/serval-go/packages/param"
	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AppResourceRolesDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[AppResourceRolesItemsDataSourceModel] `json:"data,computed"`
}

type AppResourceRolesDataSourceModel struct {
	AppInstanceID types.String                                                       `tfsdk:"app_instance_id" query:"appInstanceId,optional"`
	ResourceID    types.String                                                       `tfsdk:"resource_id" query:"resourceId,optional"`
	TeamID        types.String                                                       `tfsdk:"team_id" query:"teamId,optional"`
	MaxItems      types.Int64                                                        `tfsdk:"max_items"`
	Items         customfield.NestedObjectList[AppResourceRolesItemsDataSourceModel] `tfsdk:"items"`
}

func (m *AppResourceRolesDataSourceModel) toListParams(_ context.Context) (params serval.AppResourceRoleListParams, diags diag.Diagnostics) {
	params = serval.AppResourceRoleListParams{}

	if !m.AppInstanceID.IsNull() {
		params.AppInstanceID = param.NewOpt(m.AppInstanceID.ValueString())
	}
	if !m.ResourceID.IsNull() {
		params.ResourceID = param.NewOpt(m.ResourceID.ValueString())
	}
	if !m.TeamID.IsNull() {
		params.TeamID = param.NewOpt(m.TeamID.ValueString())
	}

	return
}

type AppResourceRolesItemsDataSourceModel struct {
	ID                 types.String                                                                `tfsdk:"id" json:"id,computed"`
	AccessPolicyID     types.String                                                                `tfsdk:"access_policy_id" json:"accessPolicyId,computed"`
	Description        types.String                                                                `tfsdk:"description" json:"description,computed"`
	ExternalData       types.String                                                                `tfsdk:"external_data" json:"externalData,computed"`
	ExternalID         types.String                                                                `tfsdk:"external_id" json:"externalId,computed"`
	Name               types.String                                                                `tfsdk:"name" json:"name,computed"`
	ProvisioningMethod customfield.NestedObject[AppResourceRolesProvisioningMethodDataSourceModel] `tfsdk:"provisioning_method" json:"provisioningMethod,computed"`
	RequestsEnabled    types.Bool                                                                  `tfsdk:"requests_enabled" json:"requestsEnabled,computed"`
	ResourceID         types.String                                                                `tfsdk:"resource_id" json:"resourceId,computed"`
}

type AppResourceRolesProvisioningMethodDataSourceModel struct {
	BuiltinWorkflow jsontypes.Normalized                                                                      `tfsdk:"builtin_workflow" json:"builtinWorkflow,computed"`
	CustomWorkflow  customfield.NestedObject[AppResourceRolesProvisioningMethodCustomWorkflowDataSourceModel] `tfsdk:"custom_workflow" json:"customWorkflow,computed"`
	LinkedRoles     customfield.NestedObject[AppResourceRolesProvisioningMethodLinkedRolesDataSourceModel]    `tfsdk:"linked_roles" json:"linkedRoles,computed"`
	Manual          customfield.NestedObject[AppResourceRolesProvisioningMethodManualDataSourceModel]         `tfsdk:"manual" json:"manual,computed"`
}

type AppResourceRolesProvisioningMethodCustomWorkflowDataSourceModel struct {
	DeprovisionWorkflowID types.String `tfsdk:"deprovision_workflow_id" json:"deprovisionWorkflowId,computed"`
	ProvisionWorkflowID   types.String `tfsdk:"provision_workflow_id" json:"provisionWorkflowId,computed"`
}

type AppResourceRolesProvisioningMethodLinkedRolesDataSourceModel struct {
	LinkedRoleIDs customfield.List[types.String] `tfsdk:"linked_role_ids" json:"linkedRoleIds,computed"`
}

type AppResourceRolesProvisioningMethodManualDataSourceModel struct {
	Assignees customfield.NestedObjectList[AppResourceRolesProvisioningMethodManualAssigneesDataSourceModel] `tfsdk:"assignees" json:"assignees,computed"`
}

type AppResourceRolesProvisioningMethodManualAssigneesDataSourceModel struct {
	AssigneeID   types.String `tfsdk:"assignee_id" json:"assigneeId,computed"`
	AssigneeType types.String `tfsdk:"assignee_type" json:"assigneeType,computed"`
}
