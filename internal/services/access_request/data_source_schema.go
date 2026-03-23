// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_request

import (
	"context"

	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
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
			"created_at": schema.StringAttribute{
				Description: "The timestamp when the access request was created.",
				Computed:    true,
			},
			"expires_at": schema.StringAttribute{
				Description: "The timestamp when the currently active time allocation expires. This is\n only set when the access request has been provisioned and is not yet\n concluded. If the request has been extended, this reflects the expiry of\n the latest (current) time allocation, not the original one.",
				Computed:    true,
			},
			"linked_ticket_id": schema.StringAttribute{
				Description: "The ID of the ticket that originated this access request. This always\n matches the linked_ticket_id of the first time allocation. Extensions\n may be created from a different ticket — see each time allocation's\n linked_ticket_id for that.",
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
			"time_allocations": schema.ListNestedAttribute{
				Description: "Every access request contains one or more time allocations. A time\n allocation represents a discrete grant (or pending grant) of access for a\n specific duration. The first time allocation is created with the initial\n request. When a user extends an existing access request, a new time\n allocation is appended — the previous one is invalidated (superseded) and\n the new one becomes the active or pending allocation. At most one time\n allocation is active and at most one is pending at any given time.\n\n Ordered by creation time ascending.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[AccessRequestTimeAllocationsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The unique ID of the time allocation.",
							Computed:    true,
						},
						"approved_minutes": schema.Int64Attribute{
							Description: "The number of minutes actually approved. Null while the time allocation is pending.",
							Computed:    true,
						},
						"business_justification": schema.StringAttribute{
							Description: "The business justification provided for this time allocation.",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Description: "The timestamp when the time allocation was created.",
							Computed:    true,
						},
						"invalidation_reason": schema.StringAttribute{
							Description: "Why the time allocation was invalidated. Only set when status is INVALIDATED.\nAvailable values: \"ACCESS_REQUEST_TIME_ALLOCATION_INVALIDATION_REASON_UNSPECIFIED\", \"ACCESS_REQUEST_TIME_ALLOCATION_INVALIDATION_REASON_SUPERSEDED\", \"ACCESS_REQUEST_TIME_ALLOCATION_INVALIDATION_REASON_CANCELED\", \"ACCESS_REQUEST_TIME_ALLOCATION_INVALIDATION_REASON_REJECTED\", \"ACCESS_REQUEST_TIME_ALLOCATION_INVALIDATION_REASON_CONCLUDED\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"ACCESS_REQUEST_TIME_ALLOCATION_INVALIDATION_REASON_UNSPECIFIED",
									"ACCESS_REQUEST_TIME_ALLOCATION_INVALIDATION_REASON_SUPERSEDED",
									"ACCESS_REQUEST_TIME_ALLOCATION_INVALIDATION_REASON_CANCELED",
									"ACCESS_REQUEST_TIME_ALLOCATION_INVALIDATION_REASON_REJECTED",
									"ACCESS_REQUEST_TIME_ALLOCATION_INVALIDATION_REASON_CONCLUDED",
								),
							},
						},
						"linked_ticket_id": schema.StringAttribute{
							Description: "The ID of the ticket this time allocation was created from, if any. For\n the initial allocation this is the ticket that originated the access\n request. For extensions this may be a different ticket.",
							Computed:    true,
						},
						"requested_by_user_id": schema.StringAttribute{
							Description: "The ID of the user who requested this time allocation (may differ from\n the original requester for extensions).",
							Computed:    true,
						},
						"requested_minutes": schema.Int64Attribute{
							Description: "The number of minutes of access requested in this time allocation.",
							Computed:    true,
						},
						"status": schema.StringAttribute{
							Description: "The status of this time allocation.\nAvailable values: \"ACCESS_REQUEST_TIME_ALLOCATION_STATUS_UNSPECIFIED\", \"ACCESS_REQUEST_TIME_ALLOCATION_STATUS_PENDING\", \"ACCESS_REQUEST_TIME_ALLOCATION_STATUS_ACTIVE\", \"ACCESS_REQUEST_TIME_ALLOCATION_STATUS_INVALIDATED\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"ACCESS_REQUEST_TIME_ALLOCATION_STATUS_UNSPECIFIED",
									"ACCESS_REQUEST_TIME_ALLOCATION_STATUS_PENDING",
									"ACCESS_REQUEST_TIME_ALLOCATION_STATUS_ACTIVE",
									"ACCESS_REQUEST_TIME_ALLOCATION_STATUS_INVALIDATED",
								),
							},
						},
					},
				},
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
