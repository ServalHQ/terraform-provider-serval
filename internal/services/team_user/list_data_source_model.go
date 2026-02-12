// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package team_user

import (
	"context"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/serval-go/packages/param"
	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TeamUsersDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[TeamUsersItemsDataSourceModel] `json:"data,computed"`
}

type TeamUsersDataSourceModel struct {
	TeamID   types.String                                                `tfsdk:"team_id" path:"team_id,required"`
	UserID   types.String                                                `tfsdk:"user_id" query:"userId,optional"`
	MaxItems types.Int64                                                 `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[TeamUsersItemsDataSourceModel] `tfsdk:"items"`
}

func (m *TeamUsersDataSourceModel) toListParams(_ context.Context) (params serval.TeamUserListParams, diags diag.Diagnostics) {
	params = serval.TeamUserListParams{}

	if !m.UserID.IsNull() {
		params.UserID = param.NewOpt(m.UserID.ValueString())
	}

	return
}

type TeamUsersItemsDataSourceModel struct {
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	Role      types.String      `tfsdk:"role" json:"role,computed"`
	TeamID    types.String      `tfsdk:"team_id" json:"teamId,computed"`
	UserID    types.String      `tfsdk:"user_id" json:"userId,computed"`
}
