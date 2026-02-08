// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/serval-go/packages/param"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UserDataDataSourceEnvelope struct {
	Data UserDataSourceModel `json:"data,computed"`
}

type UserDataSourceModel struct {
	ID            types.String                  `tfsdk:"id" path:"id,computed_optional"`
	AvatarURL     types.String                  `tfsdk:"avatar_url" json:"avatarUrl,computed"`
	CreatedAt     timetypes.RFC3339             `tfsdk:"created_at" json:"createdAt,computed" format:"date-time"`
	DeactivatedAt timetypes.RFC3339             `tfsdk:"deactivated_at" json:"deactivatedAt,computed" format:"date-time"`
	Email         types.String                  `tfsdk:"email" json:"email,computed"`
	FirstName     types.String                  `tfsdk:"first_name" json:"firstName,computed"`
	LastName      types.String                  `tfsdk:"last_name" json:"lastName,computed"`
	Name          types.String                  `tfsdk:"name" json:"name,computed"`
	Role          types.String                  `tfsdk:"role" json:"role,computed"`
	Timezone      types.String                  `tfsdk:"timezone" json:"timezone,computed"`
	FindOneBy     *UserFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

func (m *UserDataSourceModel) toListParams(_ context.Context) (params serval.UserListParams, diags diag.Diagnostics) {
	params = serval.UserListParams{}

	if !m.FindOneBy.IncludeDeactivated.IsNull() {
		params.IncludeDeactivated = param.NewOpt(m.FindOneBy.IncludeDeactivated.ValueBool())
	}

	return
}

type UserFindOneByDataSourceModel struct {
	IncludeDeactivated types.Bool `tfsdk:"include_deactivated" query:"includeDeactivated,optional"`
}
