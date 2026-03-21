// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package approval_delegation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*ApprovalDelegationResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The ID of the approval delegation.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"delegator_user_id": schema.StringAttribute{
				Description:   "The ID of the user who is delegating their approvals. When omitted, defaults to the authenticated user.",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplaceIfConfigured()},
			},
			"delegates": schema.ListNestedAttribute{
				Description: "The delegates who can approve on behalf of the delegator.",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The ID of the delegate (user ID or group ID, depending on type).",
							Optional:    true,
						},
						"type": schema.StringAttribute{
							Description: "The type of delegate (user or group).\nAvailable values: \"APPROVAL_DELEGATE_TYPE_UNSPECIFIED\", \"APPROVAL_DELEGATE_TYPE_USER\", \"APPROVAL_DELEGATE_TYPE_GROUP\".",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"APPROVAL_DELEGATE_TYPE_UNSPECIFIED",
									"APPROVAL_DELEGATE_TYPE_USER",
									"APPROVAL_DELEGATE_TYPE_GROUP",
								),
							},
						},
					},
				},
			},
			"description": schema.StringAttribute{
				Description: "A description of the delegation rule.",
				Optional:    true,
			},
			"priority": schema.Int64Attribute{
				Description: "The priority of the delegation rule (lower values are higher priority).",
				Optional:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "The timestamp when the delegation was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *ApprovalDelegationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ApprovalDelegationResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
