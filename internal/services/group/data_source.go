// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package group

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/ServalHQ/terraform-provider-serval/internal/apijson"
	"github.com/ServalHQ/terraform-provider-serval/internal/client"
	authapi "github.com/ServalHQ/terraform-provider-serval/internal/gen/proto/svauth/public"
	"github.com/ServalHQ/terraform-provider-serval/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"google.golang.org/protobuf/proto"
)

type GroupDataSource struct {
	client *client.Client
}

var _ datasource.DataSourceWithConfigure = (*GroupDataSource)(nil)

func NewGroupDataSource() datasource.DataSource {
	return &GroupDataSource{}
}

func (d *GroupDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_group"
}

func (d *GroupDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *GroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *GroupDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If ID is provided, use direct Get method
	if !data.ID.IsNull() && !data.ID.IsUnknown() {
		env := GroupDataDataSourceEnvelope{*data}
		res, err := d.client.DoID(ctx, http.MethodGet, "v2/groups", data.ID.ValueString(), nil, logging.Middleware(ctx))
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
		return
	}

	// If name is provided, list all groups and filter by name
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		targetName := data.Name.ValueString()

		params := &authapi.ListGroupsRequest{
			PageSize: proto.Int32(1000),
		}

		page, err := d.client.FirstPage(ctx, "v2/groups", params)
		if err != nil {
			resp.Diagnostics.AddError("failed to list groups", err.Error())
			return
		}

		bytes := page.Raw

		// Parse the list response - API returns {data: [...]}
		var listResponse struct {
			Data []GroupDataSourceModel `json:"data"`
		}
		err = apijson.UnmarshalComputed(bytes, &listResponse)
		if err != nil {
			resp.Diagnostics.AddError("failed to deserialize list response", err.Error())
			return
		}

		// Check if we found the group
		for i := range listResponse.Data {
			if listResponse.Data[i].Name.ValueString() == targetName {
				data = &listResponse.Data[i]
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}

		// Group not found
		resp.Diagnostics.AddError(
			"group not found",
			fmt.Sprintf("No group found with name: %s", targetName),
		)
		return
	}

	// Should never reach here due to validator, but handle it anyway
	resp.Diagnostics.AddError(
		"missing required field",
		"Either 'id' or 'name' must be specified",
	)
}
