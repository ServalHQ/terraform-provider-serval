// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_service

import (
	"context"
	"fmt"

	"github.com/ServalHQ/terraform-provider-serval/internal/apijson"
	"github.com/ServalHQ/terraform-provider-serval/internal/client"
	"github.com/ServalHQ/terraform-provider-serval/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type CustomServicesDataSource struct {
	client *client.Client
}

var _ datasource.DataSourceWithConfigure = (*CustomServicesDataSource)(nil)

func NewCustomServicesDataSource() datasource.DataSource {
	return &CustomServicesDataSource{}
}

func (d *CustomServicesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_custom_services"
}

func (d *CustomServicesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *CustomServicesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *CustomServicesDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params, diags := data.toListParams(ctx)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	env := CustomServicesDataListDataSourceEnvelope{}
	maxItems := int(data.MaxItems.ValueInt64())
	acc := []attr.Value{}
	if maxItems <= 0 {
		maxItems = 1000
	}
	err := d.client.EachPage(ctx, "v2/custom-services", params, func(raw []byte) (bool, error) {
		if err := apijson.UnmarshalComputed(raw, &env); err != nil {
			return false, err
		}
		acc = append(acc, env.Data.Elements()...)
		return len(acc) < maxItems, nil
	})
	if err != nil {
		resp.Diagnostics.AddError(client.ListErrorSummary(err), err.Error())
		return
	}

	acc = acc[:min(len(acc), maxItems)]
	result, diags := customfield.NewObjectListFromAttributes[CustomServicesItemsDataSourceModel](ctx, acc)
	resp.Diagnostics.Append(diags...)
	data.Items = result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
