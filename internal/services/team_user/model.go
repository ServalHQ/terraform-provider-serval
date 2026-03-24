// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package team_user

import (
	"github.com/ServalHQ/terraform-provider-serval/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TeamUserDataEnvelope struct {
	Data TeamUserModel `json:"data"`
}

type TeamUserModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	Role      types.String      `tfsdk:"role" json:"role,required"`
	TeamID    types.String      `tfsdk:"team_id" json:"teamId,required"`
	UserID    types.String      `tfsdk:"user_id" json:"userId,required"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
}

func (m TeamUserModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m TeamUserModel) MarshalJSONForUpdate(state TeamUserModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
