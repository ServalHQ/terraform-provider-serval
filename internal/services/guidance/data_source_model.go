// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package guidance

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type GuidanceDataDataSourceEnvelope struct {
	Data GuidanceDataSourceModel `json:"data,computed"`
}

type GuidanceDataSourceModel struct {
	ID                    types.String `tfsdk:"id" path:"id,required"`
	Content               types.String `tfsdk:"content" json:"content,computed"`
	Description           types.String `tfsdk:"description" json:"description,computed"`
	HasUnpublishedChanges types.Bool   `tfsdk:"has_unpublished_changes" json:"hasUnpublishedChanges,computed"`
	IsPublished           types.Bool   `tfsdk:"is_published" json:"isPublished,computed"`
	Name                  types.String `tfsdk:"name" json:"name,computed"`
	ShouldAlwaysUse       types.Bool   `tfsdk:"should_always_use" json:"shouldAlwaysUse,computed"`
	TeamID                types.String `tfsdk:"team_id" json:"teamId,computed"`
}
