// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package team_user

import (
	"context"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/serval-go/packages/param"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TeamUserDataDataSourceEnvelope struct {
	Data TeamUserDataSourceModel `json:"data,computed"`
}

type TeamUserDataSourceModel struct {
	ID        types.String      `tfsdk:"id" path:"id,required"`
	TeamID    types.String      `tfsdk:"team_id" query:"teamId,computed_optional"`
	UserID    types.String      `tfsdk:"user_id" query:"userId,computed_optional"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	Role      types.String      `tfsdk:"role" json:"role,computed"`
}

func (m *TeamUserDataSourceModel) toReadParams(_ context.Context) (params serval.TeamUserGetParams, diags diag.Diagnostics) {
	params = serval.TeamUserGetParams{}

	if !m.TeamID.IsNull() {
		params.TeamID = param.NewOpt(m.TeamID.ValueString())
	}
	if !m.UserID.IsNull() {
		params.UserID = param.NewOpt(m.UserID.ValueString())
	}

	return
}
