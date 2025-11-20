// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflow

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*WorkflowResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The ID of the workflow.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"team_id": schema.StringAttribute{
				Description:   "The ID of the team.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"content": schema.StringAttribute{
				Description: "The content/code of the workflow.",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "A description of the workflow.",
				Optional:    true,
			},
			"execution_scope": schema.StringAttribute{
				Description: "The execution scope of the workflow.\nAvailable values: \"WORKFLOW_EXECUTION_SCOPE_UNSPECIFIED\", \"TEAM_PRIVATE\", \"TEAM_PUBLIC\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"WORKFLOW_EXECUTION_SCOPE_UNSPECIFIED",
						"TEAM_PRIVATE",
						"TEAM_PUBLIC",
					),
				},
			},
			"is_temporary": schema.BoolAttribute{
				Description: "Whether the workflow is temporary.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the workflow.",
				Optional:    true,
			},
			"parameters": schema.StringAttribute{
				Description: "The parameters schema of the workflow (JSON).",
				Optional:    true,
			},
			"require_form_confirmation": schema.BoolAttribute{
				Description: "Whether the workflow requires form confirmation.",
				Optional:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of the workflow.\nAvailable values: \"WORKFLOW_TYPE_UNSPECIFIED\", \"EXECUTABLE\", \"GUIDANCE\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"WORKFLOW_TYPE_UNSPECIFIED",
						"EXECUTABLE",
						"GUIDANCE",
					),
				},
			},
		},
	}
}

func (r *WorkflowResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *WorkflowResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
