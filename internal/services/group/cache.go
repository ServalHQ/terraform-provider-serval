package group

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

// Cache is the global group cache, initialized by the provider.
var Cache *cache.ResourceCache[GroupModel]

// InitCache initializes the group cache. Call this from provider.Configure().
func InitCache() {
	Cache = cache.NewResourceCache[GroupModel]()
}

// GetCached retrieves a group from the cache, loading via List API if needed.
// Returns (model, found, error). If the cache fails to load, error is non-nil.
// If the group doesn't exist, found is false.
func GetCached(ctx context.Context, client *serval.Client, id string) (*GroupModel, bool, error) {
	if Cache == nil {
		return nil, false, nil // Cache not initialized, caller should fall back to API
	}

	return Cache.GetOrLoad(id, func() (map[string]*GroupModel, error) {
		return fetchAllGroups(ctx, client)
	})
}

// fetchAllGroups retrieves all groups via the List API, handling pagination.
func fetchAllGroups(ctx context.Context, client *serval.Client) (map[string]*GroupModel, error) {
	result := make(map[string]*GroupModel)
	var pageToken *string

	for {
		params := serval.GroupListParams{
			PageSize: serval.Int(5000),
		}
		if pageToken != nil {
			params.PageToken = serval.String(*pageToken)
		}

		res := new(http.Response)
		_, err := client.Groups.List(ctx, params,
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
			Data          []GroupModel `json:"data"`
			NextPageToken *string      `json:"nextPageToken,omitempty"`
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
