// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_resource_role

import (
	"context"

	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*AppResourceRoleDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The ID of the role.",
				Required:    true,
			},
			"access_policy_id": schema.StringAttribute{
				Description: "The default access policy for the role.",
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Description: "A description of the role.",
				Computed:    true,
			},
			"external_data": schema.StringAttribute{
				Description: "Data from the external system as a JSON string (computed by server).",
				Computed:    true,
			},
			"external_id": schema.StringAttribute{
				Description: "The external ID of the role in the external system (optional).",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the role.",
				Computed:    true,
			},
			"requests_enabled": schema.BoolAttribute{
				Description: "Whether requests are enabled for the role.",
				Computed:    true,
			},
			"resource_id": schema.StringAttribute{
				Description: "The ID of the resource that the role belongs to.",
				Computed:    true,
			},
			"provisioning_method": schema.SingleNestedAttribute{
				Description: "Provisioning configuration. **Exactly one method should be set.**",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[AppResourceRoleProvisioningMethodDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"builtin_workflow": schema.StringAttribute{
						Description: "Provisioning is handled by the service's builtin workflow integration.",
						Computed:    true,
						CustomType:  jsontypes.NormalizedType{},
					},
					"custom_workflow": schema.SingleNestedAttribute{
						Description: "Provisioning is handled by custom workflows for provision + deprovision.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[AppResourceRoleProvisioningMethodCustomWorkflowDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"deprovision_workflow_id": schema.StringAttribute{
								Description: "The workflow ID to deprovision access.",
								Computed:    true,
							},
							"provision_workflow_id": schema.StringAttribute{
								Description: "The workflow ID to provision access.",
								Computed:    true,
							},
						},
					},
					"linked_roles": schema.SingleNestedAttribute{
						Description: "Provisioning depends on prerequisite roles being provisioned first.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[AppResourceRoleProvisioningMethodLinkedRolesDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"linked_role_ids": schema.ListAttribute{
								Description: "The IDs of prerequisite roles.",
								Computed:    true,
								CustomType:  customfield.NewListType[types.String](ctx),
								ElementType: types.StringType,
							},
						},
					},
					"manual": schema.SingleNestedAttribute{
						Description: "Provisioning is handled manually by assigned users/groups.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[AppResourceRoleProvisioningMethodManualDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"assignees": schema.ListNestedAttribute{
								Description: "Users and groups that should be assigned/notified for manual provisioning.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectListType[AppResourceRoleProvisioningMethodManualAssigneesDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"assignee_id": schema.StringAttribute{
											Description: "The ID of the user or group.",
											Computed:    true,
										},
										"assignee_type": schema.StringAttribute{
											Description: "The type of assignee.\nAvailable values: \"MANUAL_PROVISIONING_ASSIGNEE_TYPE_UNSPECIFIED\", \"MANUAL_PROVISIONING_ASSIGNEE_TYPE_USER\", \"MANUAL_PROVISIONING_ASSIGNEE_TYPE_GROUP\".",
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"MANUAL_PROVISIONING_ASSIGNEE_TYPE_UNSPECIFIED",
													"MANUAL_PROVISIONING_ASSIGNEE_TYPE_USER",
													"MANUAL_PROVISIONING_ASSIGNEE_TYPE_GROUP",
												),
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *AppResourceRoleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *AppResourceRoleDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
