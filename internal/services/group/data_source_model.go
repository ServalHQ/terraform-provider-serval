// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package group

import (
	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type GroupDataDataSourceEnvelope struct {
	Data GroupDataSourceModel `json:"data,computed"`
}

type GroupDataSourceModel struct {
	ID             types.String                   `tfsdk:"id" path:"id,required"`
	CreatedAt      timetypes.RFC3339              `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	DeletedAt      timetypes.RFC3339              `tfsdk:"deleted_at" json:"deletedAt,computed" format:"date-time"`
	Name           types.String                   `tfsdk:"name" json:"name,computed"`
	OrganizationID types.String                   `tfsdk:"organization_id" json:"organizationId,computed"`
	UserIDs        customfield.List[types.String] `tfsdk:"user_ids" json:"userIds,computed"`
}
