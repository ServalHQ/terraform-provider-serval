// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_service

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomServiceDataDataSourceEnvelope struct {
	Data CustomServiceDataSourceModel `json:"data,computed"`
}

type CustomServiceDataSourceModel struct {
	ID     types.String `tfsdk:"id" path:"id,required"`
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
	Name   types.String `tfsdk:"name" json:"name,computed"`
	TeamID types.String `tfsdk:"team_id" json:"teamId,computed"`
}
