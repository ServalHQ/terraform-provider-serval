// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy

import (
	"context"

	"github.com/ServalHQ/terraform-provider-serval/internal/client"
	flowapi "github.com/ServalHQ/terraform-provider-serval/internal/gen/proto/svflow/publicapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"google.golang.org/protobuf/proto"
)

type AccessPolicyDataDataSourceEnvelope struct {
	Data AccessPolicyDataSourceModel `json:"data,computed"`
}

type AccessPolicyDataSourceModel struct {
	ID                           types.String                          `tfsdk:"id" path:"id,computed_optional"`
	Description                  types.String                          `tfsdk:"description" json:"description,computed"`
	MaxAccessMinutes             types.Int64                           `tfsdk:"max_access_minutes" json:"maxAccessMinutes,computed"`
	Name                         types.String                          `tfsdk:"name" json:"name,computed"`
	RecommendedAccessMinutes     types.Int64                           `tfsdk:"recommended_access_minutes" json:"recommendedAccessMinutes,computed"`
	RequireBusinessJustification types.Bool                            `tfsdk:"require_business_justification" json:"requireBusinessJustification,computed"`
	TeamID                       types.String                          `tfsdk:"team_id" json:"teamId,computed"`
	FindOneBy                    *AccessPolicyFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

func (m *AccessPolicyDataSourceModel) toListParams(_ context.Context) (params proto.Message, diags diag.Diagnostics) {
	req := &flowapi.ListAccessPoliciesRequest{}
	var explicit []string

	if !m.FindOneBy.TeamID.IsNull() {
		req.TeamId = m.FindOneBy.TeamID.ValueString()
		explicit = append(explicit, "teamId")
	}

	params = client.WithExplicit(req, explicit...)
	return
}

type AccessPolicyFindOneByDataSourceModel struct {
	TeamID types.String `tfsdk:"team_id" query:"teamId,optional"`
}
