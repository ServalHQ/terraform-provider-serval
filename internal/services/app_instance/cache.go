package app_instance

import (
	"context"
	"io"
	"net/http"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/serval-go/option"
	"github.com/ServalHQ/terraform-provider-serval/internal/apijson"
	"github.com/ServalHQ/terraform-provider-serval/internal/cache"
	"github.com/ServalHQ/terraform-provider-serval/internal/logging"
)

// Cache is the global app instance cache, initialized by the provider.
var Cache *cache.ResourceCache[AppInstanceModel]

// InitCache initializes the app instance cache. Call this from provider.Configure().
func InitCache() {
	Cache = cache.NewResourceCache[AppInstanceModel]()
}

// GetCached retrieves an app instance from the cache, loading via List API if needed.
func GetCached(
	ctx context.Context,
	client *serval.Client,
	id string,
) (*AppInstanceModel, bool, error) {
	if Cache == nil {
		return nil, false, nil
	}

	return Cache.GetOrLoad(id, func() (map[string]*AppInstanceModel, error) {
		return fetchAllAppInstances(ctx, client)
	})
}

func fetchAllAppInstances(
	ctx context.Context,
	client *serval.Client,
) (map[string]*AppInstanceModel, error) {
	result := make(map[string]*AppInstanceModel)
	var pageToken *string

	for {
		params := serval.AppInstanceListParams{
			PageSize: serval.Int(1000),
		}
		if pageToken != nil {
			params.PageToken = serval.String(*pageToken)
		}

		res := new(http.Response)
		_, err := client.AppInstances.List(ctx, params,
			option.WithResponseBodyInto(&res),
			option.WithMiddleware(logging.Middleware(ctx)),
		)
		if err != nil {
			return nil, err
		}

		bytes, err := io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return nil, err
		}

		var page struct {
			Data          []AppInstanceModel `json:"data"`
			NextPageToken *string            `json:"nextPageToken,omitempty"`
		}
		if err := apijson.Unmarshal(bytes, &page); err != nil {
			return nil, err
		}

		for i := range page.Data {
			item := page.Data[i]
			result[item.ID.ValueString()] = &item
		}

		if page.NextPageToken == nil || *page.NextPageToken == "" {
			break
		}
		pageToken = page.NextPageToken
	}

	return result, nil
}
