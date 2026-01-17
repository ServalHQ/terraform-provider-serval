// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflow

import (
	"context"

	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*WorkflowDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The ID of the workflow.",
				Required:    true,
			},
			"content": schema.StringAttribute{
				Description: "The content/code of the workflow.",
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Description: "A description of the workflow.",
				Computed:    true,
			},
			"execution_scope": schema.StringAttribute{
				Description: "The execution scope of the workflow.\nAvailable values: \"WORKFLOW_EXECUTION_SCOPE_UNSPECIFIED\", \"TEAM_PRIVATE\", \"TEAM_PUBLIC\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"WORKFLOW_EXECUTION_SCOPE_UNSPECIFIED",
						"TEAM_PRIVATE",
						"TEAM_PUBLIC",
					),
				},
			},
			"has_unpublished_changes": schema.BoolAttribute{
				Description: "Whether there are unpublished changes to the workflow.",
				Computed:    true,
			},
			"is_published": schema.BoolAttribute{
				Description: "Whether the workflow has been published at least once.",
				Computed:    true,
			},
			"is_temporary": schema.BoolAttribute{
				Description: "Whether the workflow is temporary.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the workflow.",
				Computed:    true,
			},
			"parameters": schema.StringAttribute{
				Description: "The parameters schema of the workflow (JSON).",
				Computed:    true,
			},
			"require_form_confirmation": schema.BoolAttribute{
				Description: "Whether the workflow requires form confirmation.",
				Computed:    true,
			},
			"team_id": schema.StringAttribute{
				Description: "The ID of the team that the workflow belongs to.",
				Computed:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of the workflow.\nAvailable values: \"WORKFLOW_TYPE_UNSPECIFIED\", \"EXECUTABLE\", \"GUIDANCE\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"WORKFLOW_TYPE_UNSPECIFIED",
						"EXECUTABLE",
						"GUIDANCE",
					),
				},
			},
			"tag_ids": schema.ListAttribute{
				Description: "IDs of tags associated with this workflow.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
		},
	}
}

func (d *WorkflowDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *WorkflowDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
