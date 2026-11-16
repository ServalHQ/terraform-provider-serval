// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tag

import (
	"context"

	"github.com/ServalHQ/terraform-provider-serval/internal/client"
	flowapi "github.com/ServalHQ/terraform-provider-serval/internal/gen/proto/svflow/publicapi"
	"google.golang.org/protobuf/proto"

	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TagsDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[TagsItemsDataSourceModel] `json:"data,computed"`
}

type TagsDataSourceModel struct {
	TeamID   types.String                                           `tfsdk:"team_id" query:"teamId,optional"`
	MaxItems types.Int64                                            `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[TagsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *TagsDataSourceModel) toListParams(_ context.Context) (params proto.Message, diags diag.Diagnostics) {
	req := &flowapi.ListTagsRequest{}
	var explicit []string

	if !m.TeamID.IsNull() {
		req.TeamId = m.TeamID.ValueString()
		explicit = append(explicit, "teamId")
	}

	params = client.WithExplicit(req, explicit...)
	return
}

type TagsItemsDataSourceModel struct {
	ID       types.String `tfsdk:"id" json:"id,computed"`
	Color    types.String `tfsdk:"color" json:"color,computed"`
	IconSlug types.String `tfsdk:"icon_slug" json:"iconSlug,computed"`
	Name     types.String `tfsdk:"name" json:"name,computed"`
	TeamID   types.String `tfsdk:"team_id" json:"teamId,computed"`
}
