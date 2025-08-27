// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/serval-terraform/internal/apijson"
	"github.com/stainless-sdks/serval-terraform/internal/customfield"
)

type UserDataEnvelope struct {
	Data UserModel `json:"data"`
}

type UserModel struct {
	ID            types.String                                  `tfsdk:"id" json:"id,computed"`
	TeamIDs       *[]types.String                               `tfsdk:"team_ids" json:"teamIds,optional,no_refresh"`
	AvatarURL     types.String                                  `tfsdk:"avatar_url" json:"avatarUrl,optional"`
	Email         types.String                                  `tfsdk:"email" json:"email,optional"`
	FirstName     types.String                                  `tfsdk:"first_name" json:"firstName,optional"`
	LastName      types.String                                  `tfsdk:"last_name" json:"lastName,optional"`
	Role          types.String                                  `tfsdk:"role" json:"role,optional"`
	CreatedAt     timetypes.RFC3339                             `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	DeactivatedAt timetypes.RFC3339                             `tfsdk:"deactivated_at" json:"deactivatedAt,computed" format:"date-time"`
	Name          types.String                                  `tfsdk:"name" json:"name,computed"`
	Groups        customfield.NestedObjectList[UserGroupsModel] `tfsdk:"groups" json:"groups,computed"`
	Teams         customfield.NestedObjectList[UserTeamsModel]  `tfsdk:"teams" json:"teams,computed"`
}

func (m UserModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m UserModel) MarshalJSONForUpdate(state UserModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type UserGroupsModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	DeletedAt timetypes.RFC3339 `tfsdk:"deleted_at" json:"deletedAt,computed" format:"date-time"`
	Name      types.String      `tfsdk:"name" json:"name,computed"`
}

type UserTeamsModel struct {
	Role types.String                                 `tfsdk:"role" json:"role,computed"`
	Team customfield.NestedObject[UserTeamsTeamModel] `tfsdk:"team" json:"team,computed"`
}

type UserTeamsTeamModel struct {
	ID          types.String      `tfsdk:"id" json:"id,computed"`
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	Description types.String      `tfsdk:"description" json:"description,computed"`
	Name        types.String      `tfsdk:"name" json:"name,computed"`
	OrgID       types.String      `tfsdk:"org_id" json:"orgId,computed"`
	Prefix      types.String      `tfsdk:"prefix" json:"prefix,computed"`
}
