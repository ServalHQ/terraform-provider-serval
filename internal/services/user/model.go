// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"github.com/ServalHQ/terraform-provider-serval/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UserDataEnvelope struct {
	Data UserModel `json:"data"`
}

type UserModel struct {
	ID            types.String      `tfsdk:"id" json:"id,computed"`
	Email         types.String      `tfsdk:"email" json:"email,required"`
	AvatarURL     types.String      `tfsdk:"avatar_url" json:"avatarUrl,optional"`
	FirstName     types.String      `tfsdk:"first_name" json:"firstName,optional"`
	LastName      types.String      `tfsdk:"last_name" json:"lastName,optional"`
	Role          types.String      `tfsdk:"role" json:"role,optional"`
	CreatedAt     timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	DeactivatedAt timetypes.RFC3339 `tfsdk:"deactivated_at" json:"deactivatedAt,computed" format:"date-time"`
	Name          types.String      `tfsdk:"name" json:"name,computed"`
	Timezone      types.String      `tfsdk:"timezone" json:"timezone,computed"`
}

func (m UserModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m UserModel) MarshalJSONForUpdate(state UserModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
