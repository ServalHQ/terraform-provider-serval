// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package team_user

import (
	"context"

	"github.com/ServalHQ/serval-go"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TeamUserDataDataSourceEnvelope struct {
	Data TeamUserDataSourceModel `json:"data,computed"`
}

type TeamUserDataSourceModel struct {
	TeamID    types.String      `tfsdk:"team_id" path:"team_id,required"`
	UserID    types.String      `tfsdk:"user_id" path:"user_id,required"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	Role      types.String      `tfsdk:"role" json:"role,computed"`
}

func (m *TeamUserDataSourceModel) toReadParams(_ context.Context) (params serval.TeamUserGetParams, diags diag.Diagnostics) {
	params = serval.TeamUserGetParams{
		TeamID: m.TeamID.ValueString(),
	}

	return
}
