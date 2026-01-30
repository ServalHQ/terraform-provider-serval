// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tag

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TagDataDataSourceEnvelope struct {
	Data TagDataSourceModel `json:"data,computed"`
}

type TagDataSourceModel struct {
	ID       types.String `tfsdk:"id" path:"id,required"`
	Color    types.String `tfsdk:"color" json:"color,computed"`
	IconSlug types.String `tfsdk:"icon_slug" json:"iconSlug,computed"`
	Name     types.String `tfsdk:"name" json:"name,computed"`
}
