package workflow

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

// Cache is the global workflow cache, initialized by the provider.
var Cache *cache.ResourceCache[WorkflowModel]

// InitCache initializes the workflow cache. Call this from provider.Configure().
func InitCache() {
	Cache = cache.NewResourceCache[WorkflowModel]()
}

// GetCached retrieves a workflow from the cache, loading via List API if needed.
func GetCached(
	ctx context.Context,
	client *serval.Client,
	id string,
) (*WorkflowModel, bool, error) {
	if Cache == nil {
		return nil, false, nil
	}

	return Cache.GetOrLoad(id, func() (map[string]*WorkflowModel, error) {
		return fetchAllWorkflows(ctx, client)
	})
}

func fetchAllWorkflows(
	ctx context.Context,
	client *serval.Client,
) (map[string]*WorkflowModel, error) {
	result := make(map[string]*WorkflowModel)
	var pageToken *string

	for {
		params := serval.WorkflowListParams{
			PageSize:         serval.Int(1000),
			IncludeTemporary: serval.Bool(true),
		}
		if pageToken != nil {
			params.PageToken = serval.String(*pageToken)
		}

		res := new(http.Response)
		_, err := client.Workflows.List(ctx, params,
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
			Data          []WorkflowModel `json:"data"`
			NextPageToken *string         `json:"nextPageToken,omitempty"`
		}
		if err := apijson.UnmarshalComputed(bytes, &page); err != nil {
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
