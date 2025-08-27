// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package group

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/serval-terraform/internal/apijson"
	"github.com/stainless-sdks/serval-terraform/internal/customfield"
)

type GroupDataEnvelope struct {
	Data GroupModel `json:"data"`
}

type GroupModel struct {
	ID             types.String                                  `tfsdk:"id" json:"id,computed"`
	Name           types.String                                  `tfsdk:"name" json:"name,optional"`
	UserIDs        *[]types.String                               `tfsdk:"user_ids" json:"userIds,optional,no_refresh"`
	CreatedAt      timetypes.RFC3339                             `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	DeletedAt      timetypes.RFC3339                             `tfsdk:"deleted_at" json:"deletedAt,computed" format:"date-time"`
	OrganizationID types.String                                  `tfsdk:"organization_id" json:"organizationId,computed"`
	Users          customfield.NestedObjectList[GroupUsersModel] `tfsdk:"users" json:"users,computed"`
}

func (m GroupModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m GroupModel) MarshalJSONForUpdate(state GroupModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type GroupUsersModel struct {
	ID            types.String      `tfsdk:"id" json:"id,computed"`
	AvatarURL     types.String      `tfsdk:"avatar_url" json:"avatarUrl,computed"`
	CreatedAt     timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	DeactivatedAt timetypes.RFC3339 `tfsdk:"deactivated_at" json:"deactivatedAt,computed" format:"date-time"`
	Email         types.String      `tfsdk:"email" json:"email,computed"`
	FirstName     types.String      `tfsdk:"first_name" json:"firstName,computed"`
	LastName      types.String      `tfsdk:"last_name" json:"lastName,computed"`
	Name          types.String      `tfsdk:"name" json:"name,computed"`
	Role          types.String      `tfsdk:"role" json:"role,computed"`
}
