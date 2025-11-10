// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_resource_entitlement

import (
	"context"

	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*AppResourceEntitlementDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The ID of the entitlement.",
				Required:    true,
			},
			"access_policy_id": schema.StringAttribute{
				Description: "The default access policy for the entitlement.",
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Description: "A description of the entitlement.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the entitlement.",
				Computed:    true,
			},
			"provisioning_method": schema.StringAttribute{
				Description: "The provisioning method for the entitlement.",
				Computed:    true,
			},
			"requests_enabled": schema.BoolAttribute{
				Description: "Whether requests are enabled for the entitlement.",
				Computed:    true,
			},
			"resource_id": schema.StringAttribute{
				Description: "The ID of the resource that the entitlement belongs to.",
				Computed:    true,
			},
			"linked_entitlement_ids": schema.ListAttribute{
				Description: "The IDs of entitlements that must be provisioned before this entitlement can be provisioned (prerequisite entitlements).",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"manual_provisioning_assignees": schema.ListNestedAttribute{
				Description: "The manual provisioning assignees (users and groups) for this entitlement.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[AppResourceEntitlementManualProvisioningAssigneesDataSourceModel](ctx),
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
	}
}

func (d *AppResourceEntitlementDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *AppResourceEntitlementDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
