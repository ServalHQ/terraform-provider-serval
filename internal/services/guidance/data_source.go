// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package guidance

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/ServalHQ/terraform-provider-serval/internal/apijson"
	"github.com/ServalHQ/terraform-provider-serval/internal/client"
	"github.com/ServalHQ/terraform-provider-serval/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type GuidanceDataSource struct {
	client *client.Client
}

var _ datasource.DataSourceWithConfigure = (*GuidanceDataSource)(nil)

func NewGuidanceDataSource() datasource.DataSource {
	return &GuidanceDataSource{}
}

func (d *GuidanceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_guidance"
}

func (d *GuidanceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	c, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = c
}

func (d *GuidanceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *GuidanceDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.FindOneBy != nil {
		params, diags := data.toListParams(ctx)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		env := GuidancesDataListDataSourceEnvelope{}
		page, err := d.client.FirstPage(ctx, "v2/guidances", params)
		if err != nil {
			resp.Diagnostics.AddError("failed to make http request", err.Error())
			return
		}

		bytes := page.Raw
		err = apijson.UnmarshalComputed(bytes, &env)
		if err != nil {
			resp.Diagnostics.AddError("failed to unmarshal http request", err.Error())
			return
		}

		if count := len(env.Data.Elements()); count != 1 {
			resp.Diagnostics.AddError("failed to find exactly one result", fmt.Sprint(count)+" found")
			return
		}
		ts, diags := env.Data.AsStructSliceT(ctx)
		resp.Diagnostics.Append(diags...)
		data.ID = ts[0].ID
	}

	env := GuidanceDataDataSourceEnvelope{*data}
	res, err := d.client.DoID(ctx, http.MethodGet, "v2/guidances", data.ID.ValueString(), nil, logging.Middleware(ctx))
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
