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
	UserID    types.String                      `tfsdk:"user_id" path:"user_id,computed_optional"`
	TeamID    types.String                      `tfsdk:"team_id" path:"team_id,required"`
	ID        types.String                      `tfsdk:"id" query:"id,computed_optional"`
	CreatedAt timetypes.RFC3339                 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	Role      types.String                      `tfsdk:"role" json:"role,computed"`
	FindOneBy *TeamUserFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

func (m *TeamUserDataSourceModel) toReadParams(_ context.Context) (params serval.TeamUserGetParams, diags diag.Diagnostics) {
	params = serval.TeamUserGetParams{
		TeamID: m.TeamID.ValueString(),
	}

	if !m.ID.IsNull() {
		params.ID = param.NewOpt(m.ID.ValueString())
	}

	return
}

func (m *TeamUserDataSourceModel) toListParams(_ context.Context) (params serval.TeamUserListParams, diags diag.Diagnostics) {
	params = serval.TeamUserListParams{
		TeamID: m.TeamID.ValueString(),
	}

	if !m.FindOneBy.UserID.IsNull() {
		params.UserID = param.NewOpt(m.FindOneBy.UserID.ValueString())
	}

	return
}

type TeamUserFindOneByDataSourceModel struct {
	UserID types.String `tfsdk:"user_id" query:"userId,optional"`
}
