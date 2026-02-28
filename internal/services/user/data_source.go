// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/serval-go/option"
	"github.com/ServalHQ/terraform-provider-serval/internal/apijson"
	"github.com/ServalHQ/terraform-provider-serval/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type UserDataSource struct {
	client *serval.Client
}

var _ datasource.DataSourceWithConfigure = (*UserDataSource)(nil)

func NewUserDataSource() datasource.DataSource {
	return &UserDataSource{}
}

func (d *UserDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

func (d *UserDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *UserDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *UserDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Check if ID is provided for direct lookup
	if !data.ID.IsNull() && !data.ID.IsUnknown() {
		res := new(http.Response)
		env := UserDataDataSourceEnvelope{*data}
		_, err := d.client.Users.Get(
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
		return
	}

	// Check if email is provided for list and filter lookup
	if !data.Email.IsNull() && !data.Email.IsUnknown() {
		targetEmail := data.Email.ValueString()

		includeDeactivated := true
		if data.FindOneBy != nil && !data.FindOneBy.IncludeDeactivated.IsNull() {
			includeDeactivated = data.FindOneBy.IncludeDeactivated.ValueBool()
		}
		params := serval.UserListParams{
			PageSize:           serval.Int(1000),
			IncludeDeactivated: serval.Bool(includeDeactivated),
		}

		res := new(http.Response)
		_, err := d.client.Users.List(
			ctx,
			params,
			option.WithResponseBodyInto(&res),
			option.WithMiddleware(logging.Middleware(ctx)),
		)
		if err != nil {
			resp.Diagnostics.AddError("failed to make http request", err.Error())
			return
		}

		bytes, _ := io.ReadAll(res.Body)

		// Parse the response which contains a Data array
		var listResponse struct {
			Data []UserDataSourceModel `json:"data"`
		}
		err = apijson.UnmarshalComputed(bytes, &listResponse)
		if err != nil {
			resp.Diagnostics.AddError("failed to deserialize http response", err.Error())
			return
		}

		// Check if we found the user
		for _, user := range listResponse.Data {
			if user.Email.ValueString() == targetEmail {
				data = &user
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}

		// User not found
		resp.Diagnostics.AddError(
			"user not found",
			fmt.Sprintf("No user found with email: %s", targetEmail),
		)
		return
	}

	// This should never happen due to ExactlyOneOf validator
	resp.Diagnostics.AddError("missing required field", "Either 'id' or 'email' must be specified")
}
