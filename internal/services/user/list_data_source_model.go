// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/serval-go/packages/param"
	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UsersDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[UsersItemsDataSourceModel] `json:"data,computed"`
}

type UsersDataSourceModel struct {
	IncludeDeactivated types.Bool                                              `tfsdk:"include_deactivated" query:"includeDeactivated,optional"`
	MaxItems           types.Int64                                             `tfsdk:"max_items"`
	Items              customfield.NestedObjectList[UsersItemsDataSourceModel] `tfsdk:"items"`
}

func (m *UsersDataSourceModel) toListParams(_ context.Context) (params serval.UserListParams, diags diag.Diagnostics) {
	params = serval.UserListParams{}

	if !m.IncludeDeactivated.IsNull() {
		params.IncludeDeactivated = param.NewOpt(m.IncludeDeactivated.ValueBool())
	}

	return
}

type UsersItemsDataSourceModel struct {
	ID            types.String      `tfsdk:"id" json:"id,computed"`
	AvatarURL     types.String      `tfsdk:"avatar_url" json:"avatarUrl,computed"`
	CreatedAt     timetypes.RFC3339 `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	DeactivatedAt timetypes.RFC3339 `tfsdk:"deactivated_at" json:"deactivatedAt,computed" format:"date-time"`
	Email         types.String      `tfsdk:"email" json:"email,computed"`
	FirstName     types.String      `tfsdk:"first_name" json:"firstName,computed"`
	LastName      types.String      `tfsdk:"last_name" json:"lastName,computed"`
	Name          types.String      `tfsdk:"name" json:"name,computed"`
	Role          types.String      `tfsdk:"role" json:"role,computed"`
	Timezone      types.String      `tfsdk:"timezone" json:"timezone,computed"`
}
