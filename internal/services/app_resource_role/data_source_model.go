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

type AppResourceRoleDataDataSourceEnvelope struct {
	Data AppResourceRoleDataSourceModel `json:"data,computed"`
}

type AppResourceRoleDataSourceModel struct {
	ID                 types.String                                                               `tfsdk:"id" path:"id,computed_optional"`
	AccessPolicyID     types.String                                                               `tfsdk:"access_policy_id" json:"accessPolicyId,computed"`
	Description        types.String                                                               `tfsdk:"description" json:"description,computed"`
	ExternalData       types.String                                                               `tfsdk:"external_data" json:"externalData,computed"`
	ExternalID         types.String                                                               `tfsdk:"external_id" json:"externalId,computed"`
	Name               types.String                                                               `tfsdk:"name" json:"name,computed"`
	RequestsEnabled    types.Bool                                                                 `tfsdk:"requests_enabled" json:"requestsEnabled,computed"`
	ResourceID         types.String                                                               `tfsdk:"resource_id" json:"resourceId,computed"`
	ProvisioningMethod customfield.NestedObject[AppResourceRoleProvisioningMethodDataSourceModel] `tfsdk:"provisioning_method" json:"provisioningMethod,computed"`
	FindOneBy          *AppResourceRoleFindOneByDataSourceModel                                   `tfsdk:"find_one_by"`
}

func (m *AppResourceRoleDataSourceModel) toListParams(_ context.Context) (params serval.AppResourceRoleListParams, diags diag.Diagnostics) {
	params = serval.AppResourceRoleListParams{}

	if !m.FindOneBy.AppInstanceID.IsNull() {
		params.AppInstanceID = param.NewOpt(m.FindOneBy.AppInstanceID.ValueString())
	}
	if !m.FindOneBy.ResourceID.IsNull() {
		params.ResourceID = param.NewOpt(m.FindOneBy.ResourceID.ValueString())
	}
	if !m.FindOneBy.TeamID.IsNull() {
		params.TeamID = param.NewOpt(m.FindOneBy.TeamID.ValueString())
	}

	return
}

type AppResourceRoleProvisioningMethodDataSourceModel struct {
	BuiltinWorkflow jsontypes.Normalized                                                                     `tfsdk:"builtin_workflow" json:"builtinWorkflow,computed"`
	CustomWorkflow  customfield.NestedObject[AppResourceRoleProvisioningMethodCustomWorkflowDataSourceModel] `tfsdk:"custom_workflow" json:"customWorkflow,computed"`
	LinkedRoles     customfield.NestedObject[AppResourceRoleProvisioningMethodLinkedRolesDataSourceModel]    `tfsdk:"linked_roles" json:"linkedRoles,computed"`
	Manual          customfield.NestedObject[AppResourceRoleProvisioningMethodManualDataSourceModel]         `tfsdk:"manual" json:"manual,computed"`
}

type AppResourceRoleProvisioningMethodCustomWorkflowDataSourceModel struct {
	DeprovisionWorkflowID types.String `tfsdk:"deprovision_workflow_id" json:"deprovisionWorkflowId,computed"`
	ProvisionWorkflowID   types.String `tfsdk:"provision_workflow_id" json:"provisionWorkflowId,computed"`
}

type AppResourceRoleProvisioningMethodLinkedRolesDataSourceModel struct {
	LinkedRoleIDs customfield.List[types.String] `tfsdk:"linked_role_ids" json:"linkedRoleIds,computed"`
}

type AppResourceRoleProvisioningMethodManualDataSourceModel struct {
	Assignees customfield.NestedObjectList[AppResourceRoleProvisioningMethodManualAssigneesDataSourceModel] `tfsdk:"assignees" json:"assignees,computed"`
}

type AppResourceRoleProvisioningMethodManualAssigneesDataSourceModel struct {
	AssigneeID   types.String `tfsdk:"assignee_id" json:"assigneeId,computed"`
	AssigneeType types.String `tfsdk:"assignee_type" json:"assigneeType,computed"`
}

type AppResourceRoleFindOneByDataSourceModel struct {
	AppInstanceID types.String `tfsdk:"app_instance_id" query:"appInstanceId,optional"`
	ResourceID    types.String `tfsdk:"resource_id" query:"resourceId,optional"`
	TeamID        types.String `tfsdk:"team_id" query:"teamId,optional"`
}
