// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package team

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/serval-terraform/internal/apijson"
	"github.com/stainless-sdks/serval-terraform/internal/customfield"
)

type TeamDataEnvelope struct {
	Data TeamModel `json:"data"`
}

type TeamModel struct {
	ID          types.String                                 `tfsdk:"id" json:"id,computed"`
	Description types.String                                 `tfsdk:"description" json:"description,optional,no_refresh"`
	Name        types.String                                 `tfsdk:"name" json:"name,optional,no_refresh"`
	Prefix      types.String                                 `tfsdk:"prefix" json:"prefix,optional,no_refresh"`
	UserIDs     *[]types.String                              `tfsdk:"user_ids" json:"userIds,optional,no_refresh"`
	CreatedAt   timetypes.RFC3339                            `tfsdk:"created_at" json:"createdAt,computed,no_refresh" format:"date-time"`
	OrgID       types.String                                 `tfsdk:"org_id" json:"orgId,computed,no_refresh"`
	Team        customfield.NestedObject[TeamTeamModel]      `tfsdk:"team" json:"team,computed"`
	Users       customfield.NestedObjectList[TeamUsersModel] `tfsdk:"users" json:"users,computed"`
}

func (m TeamModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m TeamModel) MarshalJSONForUpdate(state TeamModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type TeamTeamModel struct {
	ID          types.String      `tfsdk:"id" json:"id,computed"`
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	Description types.String      `tfsdk:"description" json:"description,computed"`
	Name        types.String      `tfsdk:"name" json:"name,computed"`
	OrgID       types.String      `tfsdk:"org_id" json:"orgId,computed"`
	Prefix      types.String      `tfsdk:"prefix" json:"prefix,computed"`
}

type TeamUsersModel struct {
	ID            types.String                                       `tfsdk:"id" json:"id,computed"`
	AvatarURL     types.String                                       `tfsdk:"avatar_url" json:"avatarUrl,computed"`
	CreatedAt     timetypes.RFC3339                                  `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	DeactivatedAt timetypes.RFC3339                                  `tfsdk:"deactivated_at" json:"deactivatedAt,computed" format:"date-time"`
	Email         types.String                                       `tfsdk:"email" json:"email,computed"`
	FirstName     types.String                                       `tfsdk:"first_name" json:"firstName,computed"`
	Groups        customfield.NestedObjectList[TeamUsersGroupsModel] `tfsdk:"groups" json:"groups,computed"`
	LastName      types.String                                       `tfsdk:"last_name" json:"lastName,computed"`
	Name          types.String                                       `tfsdk:"name" json:"name,computed"`
	Role          types.String                                       `tfsdk:"role" json:"role,computed"`
	Teams         customfield.NestedObjectList[TeamUsersTeamsModel]  `tfsdk:"teams" json:"teams,computed"`
}

type TeamUsersGroupsModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	DeletedAt timetypes.RFC3339 `tfsdk:"deleted_at" json:"deletedAt,computed" format:"date-time"`
	Name      types.String      `tfsdk:"name" json:"name,computed"`
}

type TeamUsersTeamsModel struct {
	Role types.String                                      `tfsdk:"role" json:"role,computed"`
	Team customfield.NestedObject[TeamUsersTeamsTeamModel] `tfsdk:"team" json:"team,computed"`
}

type TeamUsersTeamsTeamModel struct {
	ID          types.String      `tfsdk:"id" json:"id,computed"`
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	Description types.String      `tfsdk:"description" json:"description,computed"`
	Name        types.String      `tfsdk:"name" json:"name,computed"`
	OrgID       types.String      `tfsdk:"org_id" json:"orgId,computed"`
	Prefix      types.String      `tfsdk:"prefix" json:"prefix,computed"`
}
