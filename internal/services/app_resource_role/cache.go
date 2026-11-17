package app_resource_role

import (
	"context"
	"fmt"
	"io"
	"net/http"

	flowapi "github.com/ServalHQ/terraform-provider-serval/internal/gen/proto/svflow/publicapi"
	"google.golang.org/protobuf/proto"

	"github.com/ServalHQ/terraform-provider-serval/internal/apijson"
	"github.com/ServalHQ/terraform-provider-serval/internal/cache"
	"github.com/ServalHQ/terraform-provider-serval/internal/client"
	"github.com/ServalHQ/terraform-provider-serval/internal/logging"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var Cache *cache.Store[AppResourceRoleModel]

func Prefetch(ctx context.Context, c *client.Client, teamIDs []string) (int, error) {
	Cache = cache.NewStore[AppResourceRoleModel]()
	apiCalls := 0
	for _, teamID := range teamIDs {
		var pageToken *string
		for {
			params := &flowapi.ListAppResourceRolesRequest{
				TeamId:   proto.String(teamID),
				PageSize: proto.Int32(5000),
			}
			if pageToken != nil {
				params.PageToken = proto.String(*pageToken)
			}
			query, err := client.QueryFromMessage(params)
			if err != nil {
				return apiCalls, err
			}
			res, err := c.Do(ctx, client.Request{
				Method:      http.MethodGet,
				Path:        "v2/app-resource-roles",
				Query:       query,
				Middlewares: []client.Middleware{logging.Middleware(ctx)},
			})
			apiCalls++
			if err != nil {
				if cache.IsServerError(err) {
					tflog.Warn(ctx, fmt.Sprintf("prefetch: skipping app_resource_roles for team %s due to server error: %s", teamID, err))
					break
				}
				return apiCalls, err
			}
			bytes, err := io.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				return apiCalls, err
			}
			var page struct {
				Data          []AppResourceRoleModel `json:"data"`
				NextPageToken *string                `json:"nextPageToken,omitempty"`
			}
			if err := apijson.Unmarshal(bytes, &page); err != nil {
				return apiCalls, err
			}
			for i := range page.Data {
				item := page.Data[i]
				Cache.Put(item.ID.ValueString(), &item)
			}
			if page.NextPageToken == nil || *page.NextPageToken == "" {
				break
			}
			pageToken = page.NextPageToken
		}
	}
	return apiCalls, nil
}

func TryRead(id string) (*AppResourceRoleModel, bool, error) {
	return cache.TryRead(Cache, id)
}
