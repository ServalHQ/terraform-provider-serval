package app_instance

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/serval-go/option"
	"github.com/ServalHQ/terraform-provider-serval/internal/apijson"
	"github.com/ServalHQ/terraform-provider-serval/internal/cache"
	"github.com/ServalHQ/terraform-provider-serval/internal/logging"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var Cache *cache.Store[AppInstanceModel]

func Prefetch(ctx context.Context, client *serval.Client, teamIDs []string) (int, error) {
	Cache = cache.NewStore[AppInstanceModel]()
	apiCalls := 0
	for _, teamID := range teamIDs {
		var pageToken *string
		for {
			params := serval.AppInstanceListParams{
				PageSize: serval.Int(1000),
				TeamID:   serval.String(teamID),
			}
			if pageToken != nil {
				params.PageToken = serval.String(*pageToken)
			}
			res := new(http.Response)
			_, err := client.AppInstances.List(ctx, params,
				option.WithResponseBodyInto(&res),
				option.WithMiddleware(logging.Middleware(ctx)),
			)
			apiCalls++
			if err != nil {
				if cache.IsServerError(err) {
					tflog.Warn(ctx, fmt.Sprintf("prefetch: skipping app_instances for team %s due to server error: %s", teamID, err))
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
				Data          []AppInstanceModel `json:"data"`
				NextPageToken *string            `json:"nextPageToken,omitempty"`
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

func TryRead(id string) (*AppInstanceModel, bool, error) {
	return cache.TryRead(Cache, id)
}
