// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_request

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*AccessRequestDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The ID of the access request.",
				Required:    true,
			},
			"access_minutes": schema.Int64Attribute{
				Description: "The number of minutes of access requested.",
				Computed:    true,
			},
			"business_justification": schema.StringAttribute{
				Description: "The business justification provided for the request.",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "The timestamp when the access request was created.",
				Computed:    true,
			},
			"expires_at": schema.StringAttribute{
				Description: "The timestamp when the access expires (if approved and active).",
				Computed:    true,
			},
			"linked_ticket_id": schema.StringAttribute{
				Description: "The ID of the linked ticket, if any.",
				Computed:    true,
			},
			"requested_by_user_id": schema.StringAttribute{
				Description: "The ID of the user who requested access.",
				Computed:    true,
			},
			"requested_role_id": schema.StringAttribute{
				Description: "The ID of the requested role.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "The status of the access request.\nAvailable values: \"ACCESS_REQUEST_STATUS_UNSPECIFIED\", \"ACCESS_REQUEST_STATUS_PENDING\", \"ACCESS_REQUEST_STATUS_APPROVED\", \"ACCESS_REQUEST_STATUS_DENIED\", \"ACCESS_REQUEST_STATUS_EXPIRED\", \"ACCESS_REQUEST_STATUS_REVOKED\", \"ACCESS_REQUEST_STATUS_CANCELED\", \"ACCESS_REQUEST_STATUS_FAILED\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"ACCESS_REQUEST_STATUS_UNSPECIFIED",
						"ACCESS_REQUEST_STATUS_PENDING",
						"ACCESS_REQUEST_STATUS_APPROVED",
						"ACCESS_REQUEST_STATUS_DENIED",
						"ACCESS_REQUEST_STATUS_EXPIRED",
						"ACCESS_REQUEST_STATUS_REVOKED",
						"ACCESS_REQUEST_STATUS_CANCELED",
						"ACCESS_REQUEST_STATUS_FAILED",
					),
				},
			},
			"target_user_id": schema.StringAttribute{
				Description: "The ID of the target user for whom access was requested.",
				Computed:    true,
			},
			"team_id": schema.StringAttribute{
				Description: "The ID of the team that the access request belongs to.",
				Computed:    true,
			},
		},
	}
}

func (d *AccessRequestDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *AccessRequestDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
