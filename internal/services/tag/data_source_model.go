// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tag

import (
	"context"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/serval-go/packages/param"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TagDataDataSourceEnvelope struct {
	Data TagDataSourceModel `json:"data,computed"`
}

type TagDataSourceModel struct {
	ID        types.String                 `tfsdk:"id" path:"id,computed_optional"`
	Color     types.String                 `tfsdk:"color" json:"color,computed"`
	IconSlug  types.String                 `tfsdk:"icon_slug" json:"iconSlug,computed"`
	Name      types.String                 `tfsdk:"name" json:"name,computed"`
	TeamID    types.String                 `tfsdk:"team_id" json:"teamId,computed"`
	FindOneBy *TagFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

func (m *TagDataSourceModel) toListParams(_ context.Context) (params serval.TagListParams, diags diag.Diagnostics) {
	params = serval.TagListParams{}

	if !m.FindOneBy.TeamID.IsNull() {
		params.TeamID = param.NewOpt(m.FindOneBy.TeamID.ValueString())
	}

	return
}

type TagFindOneByDataSourceModel struct {
	TeamID types.String `tfsdk:"team_id" query:"teamId,optional"`
}
