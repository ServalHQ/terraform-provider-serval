// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_service

import (
	"context"

	"github.com/ServalHQ/terraform-provider-serval/internal/client"
	flowapi "github.com/ServalHQ/terraform-provider-serval/internal/gen/proto/svflow/publicapi"
	"google.golang.org/protobuf/proto"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomServiceDataDataSourceEnvelope struct {
	Data CustomServiceDataSourceModel `json:"data,computed"`
}

type CustomServiceDataSourceModel struct {
	ID        types.String                           `tfsdk:"id" path:"id,computed_optional"`
	Domain    types.String                           `tfsdk:"domain" json:"domain,computed"`
	Name      types.String                           `tfsdk:"name" json:"name,computed"`
	TeamID    types.String                           `tfsdk:"team_id" json:"teamId,computed"`
	FindOneBy *CustomServiceFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

func (m *CustomServiceDataSourceModel) toListParams(_ context.Context) (params proto.Message, diags diag.Diagnostics) {
	req := &flowapi.ListCustomServicesRequest{}
	var explicit []string

	if !m.FindOneBy.TeamID.IsNull() {
		req.TeamId = m.FindOneBy.TeamID.ValueString()
		explicit = append(explicit, "teamId")
	}

	params = client.WithExplicit(req, explicit...)
	return
}

type CustomServiceFindOneByDataSourceModel struct {
	TeamID types.String `tfsdk:"team_id" query:"teamId,optional"`
}
