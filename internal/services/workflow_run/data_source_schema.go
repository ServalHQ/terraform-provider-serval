// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflow_run

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*WorkflowRunDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The ID of the workflow run.",
				Required:    true,
			},
			"completed_at": schema.StringAttribute{
				Description: "The timestamp when the workflow run completed (if applicable).",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "The timestamp when the workflow run was created.",
				Computed:    true,
			},
			"initiated_by_user_id": schema.StringAttribute{
				Description: "The ID of the user who initiated the workflow run.",
				Computed:    true,
			},
			"inputs": schema.StringAttribute{
				Description: "The inputs provided to the workflow (JSON string).",
				Computed:    true,
			},
			"linked_ticket_id": schema.StringAttribute{
				Description: "The ID of the linked ticket, if any.",
				Computed:    true,
			},
			"output": schema.StringAttribute{
				Description: "The output of the workflow run (JSON string, available when completed or failed).",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "The status of the workflow run.\nAvailable values: \"WORKFLOW_RUN_STATUS_UNSPECIFIED\", \"WORKFLOW_RUN_STATUS_PENDING\", \"WORKFLOW_RUN_STATUS_RUNNING\", \"WORKFLOW_RUN_STATUS_COMPLETED\", \"WORKFLOW_RUN_STATUS_FAILED\", \"WORKFLOW_RUN_STATUS_DENIED\", \"WORKFLOW_RUN_STATUS_CANCELED\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"WORKFLOW_RUN_STATUS_UNSPECIFIED",
						"WORKFLOW_RUN_STATUS_PENDING",
						"WORKFLOW_RUN_STATUS_RUNNING",
						"WORKFLOW_RUN_STATUS_COMPLETED",
						"WORKFLOW_RUN_STATUS_FAILED",
						"WORKFLOW_RUN_STATUS_DENIED",
						"WORKFLOW_RUN_STATUS_CANCELED",
					),
				},
			},
			"target_user_id": schema.StringAttribute{
				Description: "The ID of the target user for the workflow run (if different from initiator).",
				Computed:    true,
			},
			"team_id": schema.StringAttribute{
				Description: "The ID of the team that the workflow belongs to.",
				Computed:    true,
			},
			"workflow_id": schema.StringAttribute{
				Description: "The ID of the workflow that was run.",
				Computed:    true,
			},
		},
	}
}

func (d *WorkflowRunDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *WorkflowRunDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
