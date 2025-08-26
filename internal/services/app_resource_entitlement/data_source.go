// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_resource_entitlement

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/stainless-sdks/serval-go"
	"github.com/stainless-sdks/serval-go/option"
	"github.com/stainless-sdks/serval-terraform/internal/apijson"
	"github.com/stainless-sdks/serval-terraform/internal/logging"
)

type AppResourceEntitlementDataSource struct {
	client *serval.Client
}

var _ datasource.DataSourceWithConfigure = (*AppResourceEntitlementDataSource)(nil)

func NewAppResourceEntitlementDataSource() datasource.DataSource {
	return &AppResourceEntitlementDataSource{}
}

func (d *AppResourceEntitlementDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_app_resource_entitlement"
}

func (d *AppResourceEntitlementDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*serval.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *serval.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *AppResourceEntitlementDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *AppResourceEntitlementDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := AppResourceEntitlementDataDataSourceEnvelope{*data}
	_, err := d.client.AppResourceEntitlements.Get(
		ctx,
		data.ID.ValueString(),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.UnmarshalComputed(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Data

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
