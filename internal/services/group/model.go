// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package group

import (
	"github.com/ServalHQ/terraform-provider-serval/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type GroupDataEnvelope struct {
	Data GroupModel `json:"data"`
}

type GroupModel struct {
	ID             types.String      `tfsdk:"id" json:"id,computed"`
	Name           types.String      `tfsdk:"name" json:"name,required"`
	UserIDs        *[]types.String   `tfsdk:"user_ids" json:"userIds,optional"`
	CreatedAt      timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	DeletedAt      timetypes.RFC3339 `tfsdk:"deleted_at" json:"deletedAt,computed" format:"date-time"`
	OrganizationID types.String      `tfsdk:"organization_id" json:"organizationId,computed"`
}

func (m GroupModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m GroupModel) MarshalJSONForUpdate(state GroupModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

// normalizeState prevents spurious diffs by aligning what the API returns
// with what Terraform expects when a field is unset in the HCL config.
// Empty user_ids lists are normalized to null because the API returns []
// for groups with no members, but Terraform stores null when user_ids is
// omitted from config.
func (m *GroupModel) NormalizeState() {
	m.normalizeState()
}

func (m *GroupModel) normalizeState() {
	if m.UserIDs != nil && len(*m.UserIDs) == 0 {
		m.UserIDs = nil
	}
}
