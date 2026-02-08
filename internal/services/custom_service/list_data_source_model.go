// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_service

import (
	"context"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/serval-go/packages/param"
	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomServicesDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[CustomServicesItemsDataSourceModel] `json:"data,computed"`
}

type CustomServicesDataSourceModel struct {
	TeamID   types.String                                                     `tfsdk:"team_id" query:"teamId,optional"`
	MaxItems types.Int64                                                      `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[CustomServicesItemsDataSourceModel] `tfsdk:"items"`
}

func (m *CustomServicesDataSourceModel) toListParams(_ context.Context) (params serval.CustomServiceListParams, diags diag.Diagnostics) {
	params = serval.CustomServiceListParams{}

	if !m.TeamID.IsNull() {
		params.TeamID = param.NewOpt(m.TeamID.ValueString())
	}

	return
}

type CustomServicesItemsDataSourceModel struct {
	ID     types.String `tfsdk:"id" json:"id,computed"`
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
	Name   types.String `tfsdk:"name" json:"name,computed"`
	TeamID types.String `tfsdk:"team_id" json:"teamId,computed"`
}
