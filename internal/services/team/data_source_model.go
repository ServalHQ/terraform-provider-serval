// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package team

import (
	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TeamDataDataSourceEnvelope struct {
	Data TeamDataSourceModel `json:"data,computed"`
}

type TeamDataSourceModel struct {
	ID             types.String                   `tfsdk:"id" path:"id,required"`
	CreatedAt      timetypes.RFC3339              `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	Description    types.String                   `tfsdk:"description" json:"description,computed"`
	Name           types.String                   `tfsdk:"name" json:"name,computed"`
	OrganizationID types.String                   `tfsdk:"organization_id" json:"organizationId,computed"`
	Prefix         types.String                   `tfsdk:"prefix" json:"prefix,computed"`
	UserIDs        customfield.List[types.String] `tfsdk:"user_ids" json:"userIds,computed"`
}
