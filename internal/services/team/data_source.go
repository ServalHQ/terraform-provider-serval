// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package team

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

type TeamDataSource struct {
	client *serval.Client
}

var _ datasource.DataSourceWithConfigure = (*TeamDataSource)(nil)

func NewTeamDataSource() datasource.DataSource {
	return &TeamDataSource{}
}

func (d *TeamDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_team"
}

func (d *TeamDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *TeamDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *TeamDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If ID is provided, use direct Get method
	if !data.ID.IsNull() && !data.ID.IsUnknown() {
		res := new(http.Response)
		env := TeamDataDataSourceEnvelope{*data}
		_, err := d.client.Teams.Get(
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

	// If name or prefix is provided, list all teams and filter
	if (!data.Name.IsNull() && !data.Name.IsUnknown()) || (!data.Prefix.IsNull() && !data.Prefix.IsUnknown()) {
		var filterField, filterValue string
		if !data.Name.IsNull() && !data.Name.IsUnknown() {
			filterField = "name"
			filterValue = data.Name.ValueString()
		} else {
			filterField = "prefix"
			filterValue = data.Prefix.ValueString()
		}
		
		// Fetch all pages to find the team with matching name or prefix
		cursor := ""
		
		for {
			params := serval.TeamListParams{
				Limit: serval.Int(1000), // Set high limit to fetch all teams
			}
			if cursor != "" {
				params.Cursor = serval.String(cursor)
			}

			res := new(http.Response)
			_, err := d.client.Teams.List(
				ctx,
				params,
				option.WithResponseBodyInto(&res),
				option.WithMiddleware(logging.Middleware(ctx)),
			)
			if err != nil {
				resp.Diagnostics.AddError("failed to list teams", err.Error())
				return
			}
			
			bytes, _ := io.ReadAll(res.Body)
			
			// Parse the list response - API returns {data: [...], next: "..."}
			var listResponse struct {
				Data []TeamDataSourceModel `json:"data"`
				Next *string               `json:"next"`
			}
			err = apijson.UnmarshalComputed(bytes, &listResponse)
			if err != nil {
				resp.Diagnostics.AddError("failed to deserialize list response", err.Error())
				return
			}

			// Check if we found the team in this page
			for i := range listResponse.Data {
				var match bool
				if filterField == "name" {
					match = listResponse.Data[i].Name.ValueString() == filterValue
				} else {
					match = listResponse.Data[i].Prefix.ValueString() == filterValue
				}
				
				if match {
					data = &listResponse.Data[i]
					resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
					return
				}
			}

			// Check if there are more pages
			if listResponse.Next == nil || *listResponse.Next == "" {
				break
			}
			cursor = *listResponse.Next
		}

		// Team not found after checking all pages
		resp.Diagnostics.AddError(
			"team not found",
			fmt.Sprintf("No team found with %s: %s", filterField, filterValue),
		)
		return
	}

	// Should never reach here due to validator, but handle it anyway
	resp.Diagnostics.AddError(
		"missing required field",
		"Either 'id', 'name', or 'prefix' must be specified",
	)
}
