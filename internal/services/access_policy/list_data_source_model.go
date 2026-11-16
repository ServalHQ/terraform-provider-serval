// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy

import (
	"context"

	"github.com/ServalHQ/terraform-provider-serval/internal/client"
	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	flowapi "github.com/ServalHQ/terraform-provider-serval/internal/gen/proto/svflow/publicapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"google.golang.org/protobuf/proto"
)

type AccessPoliciesDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[AccessPoliciesItemsDataSourceModel] `json:"data,computed"`
}

type AccessPoliciesDataSourceModel struct {
	TeamID   types.String                                                     `tfsdk:"team_id" query:"teamId,optional"`
	MaxItems types.Int64                                                      `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[AccessPoliciesItemsDataSourceModel] `tfsdk:"items"`
}

func (m *AccessPoliciesDataSourceModel) toListParams(_ context.Context) (params proto.Message, diags diag.Diagnostics) {
	req := &flowapi.ListAccessPoliciesRequest{}
	var explicit []string

	if !m.TeamID.IsNull() {
		req.TeamId = m.TeamID.ValueString()
		explicit = append(explicit, "teamId")
	}

	params = client.WithExplicit(req, explicit...)
	return
}

type AccessPoliciesItemsDataSourceModel struct {
	ID                           types.String `tfsdk:"id" json:"id,computed"`
	Description                  types.String `tfsdk:"description" json:"description,computed"`
	MaxAccessMinutes             types.Int64  `tfsdk:"max_access_minutes" json:"maxAccessMinutes,computed"`
	Name                         types.String `tfsdk:"name" json:"name,computed"`
	RecommendedAccessMinutes     types.Int64  `tfsdk:"recommended_access_minutes" json:"recommendedAccessMinutes,computed"`
	RequireBusinessJustification types.Bool   `tfsdk:"require_business_justification" json:"requireBusinessJustification,computed"`
	TeamID                       types.String `tfsdk:"team_id" json:"teamId,computed"`
}
