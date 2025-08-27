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
	ID             types.String                                 `tfsdk:"id" json:"id,computed"`
	Description    types.String                                 `tfsdk:"description" json:"description,optional"`
	Name           types.String                                 `tfsdk:"name" json:"name,optional"`
	Prefix         types.String                                 `tfsdk:"prefix" json:"prefix,optional"`
	UserIDs        *[]types.String                              `tfsdk:"user_ids" json:"userIds,optional,no_refresh"`
	CreatedAt      timetypes.RFC3339                            `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	OrganizationID types.String                                 `tfsdk:"organization_id" json:"organizationId,computed"`
	Users          customfield.NestedObjectList[TeamUsersModel] `tfsdk:"users" json:"users,computed"`
}

func (m TeamModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m TeamModel) MarshalJSONForUpdate(state TeamModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type TeamUsersModel struct {
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
