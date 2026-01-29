package team

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

// Cache is the global team cache, initialized by the provider.
var Cache *cache.ResourceCache[TeamModel]

// InitCache initializes the team cache. Call this from provider.Configure().
func InitCache() {
	Cache = cache.NewResourceCache[TeamModel]()
}

// GetCached retrieves a team from the cache, loading via List API if needed.
func GetCached(
	ctx context.Context,
	client *serval.Client,
	id string,
) (*TeamModel, bool, error) {
	if Cache == nil {
		return nil, false, nil
	}

	return Cache.GetOrLoad(id, func() (map[string]*TeamModel, error) {
		return fetchAllTeams(ctx, client)
	})
}

func fetchAllTeams(
	ctx context.Context,
	client *serval.Client,
) (map[string]*TeamModel, error) {
	result := make(map[string]*TeamModel)
	var pageToken *string

	for {
		params := serval.TeamListParams{
			PageSize: serval.Int(1000),
		}
		if pageToken != nil {
			params.PageToken = serval.String(*pageToken)
		}

		res := new(http.Response)
		_, err := client.Teams.List(ctx, params,
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
			Data          []TeamModel `json:"data"`
			NextPageToken *string     `json:"nextPageToken,omitempty"`
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
