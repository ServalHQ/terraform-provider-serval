// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflow

import (
	"context"

	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*WorkflowDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The ID of the workflow.",
				Computed:    true,
				Optional:    true,
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
				Description: "Whether there are unpublished changes to the workflow (computed by server).",
				Computed:    true,
			},
			"is_published": schema.BoolAttribute{
				Description: "Whether the workflow is published. Set to true to publish the workflow.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the workflow.",
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
			"tag_ids": schema.ListAttribute{
				Description: "IDs of tags associated with this workflow.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"team_id": schema.StringAttribute{
						Description: "The ID of the team.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *WorkflowDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *WorkflowDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("id"), path.MatchRoot("find_one_by")),
	}
}
