// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package group

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/serval-terraform/internal/customfield"
)

type GroupDataDataSourceEnvelope struct {
	Data GroupDataSourceModel `json:"data,computed"`
}

type GroupDataSourceModel struct {
	ID    types.String                                            `tfsdk:"id" path:"id,required"`
	Group customfield.NestedObject[GroupGroupDataSourceModel]     `tfsdk:"group" json:"group,computed"`
	Users customfield.NestedObjectList[GroupUsersDataSourceModel] `tfsdk:"users" json:"users,computed"`
}

type GroupGroupDataSourceModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	DeletedAt timetypes.RFC3339 `tfsdk:"deleted_at" json:"deletedAt,computed" format:"date-time"`
	Name      types.String      `tfsdk:"name" json:"name,computed"`
}

type GroupUsersDataSourceModel struct {
	ID            types.String                                                  `tfsdk:"id" json:"id,computed"`
	AvatarURL     types.String                                                  `tfsdk:"avatar_url" json:"avatarUrl,computed"`
	CreatedAt     timetypes.RFC3339                                             `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	DeactivatedAt timetypes.RFC3339                                             `tfsdk:"deactivated_at" json:"deactivatedAt,computed" format:"date-time"`
	Email         types.String                                                  `tfsdk:"email" json:"email,computed"`
	FirstName     types.String                                                  `tfsdk:"first_name" json:"firstName,computed"`
	Groups        customfield.NestedObjectList[GroupUsersGroupsDataSourceModel] `tfsdk:"groups" json:"groups,computed"`
	LastName      types.String                                                  `tfsdk:"last_name" json:"lastName,computed"`
	Name          types.String                                                  `tfsdk:"name" json:"name,computed"`
	Role          types.String                                                  `tfsdk:"role" json:"role,computed"`
	Teams         customfield.NestedObjectList[GroupUsersTeamsDataSourceModel]  `tfsdk:"teams" json:"teams,computed"`
}

type GroupUsersGroupsDataSourceModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	DeletedAt timetypes.RFC3339 `tfsdk:"deleted_at" json:"deletedAt,computed" format:"date-time"`
	Name      types.String      `tfsdk:"name" json:"name,computed"`
}

type GroupUsersTeamsDataSourceModel struct {
	Role types.String                                                 `tfsdk:"role" json:"role,computed"`
	Team customfield.NestedObject[GroupUsersTeamsTeamDataSourceModel] `tfsdk:"team" json:"team,computed"`
}

type GroupUsersTeamsTeamDataSourceModel struct {
	ID          types.String      `tfsdk:"id" json:"id,computed"`
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	Description types.String      `tfsdk:"description" json:"description,computed"`
	Name        types.String      `tfsdk:"name" json:"name,computed"`
	OrgID       types.String      `tfsdk:"org_id" json:"orgId,computed"`
	Prefix      types.String      `tfsdk:"prefix" json:"prefix,computed"`
}
