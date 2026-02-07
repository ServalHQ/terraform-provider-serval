package user

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

var Cache *cache.Store[UserModel]

func Prefetch(ctx context.Context, client *serval.Client) (int, error) {
	Cache = cache.NewStore[UserModel]()
	apiCalls := 0
	var pageToken *string
	for {
		params := serval.UserListParams{
			PageSize:           serval.Int(1000),
			IncludeDeactivated: serval.Bool(true),
		}
		if pageToken != nil {
			params.PageToken = serval.String(*pageToken)
		}
		res := new(http.Response)
		_, err := client.Users.List(ctx, params,
			option.WithResponseBodyInto(&res),
			option.WithMiddleware(logging.Middleware(ctx)),
		)
		apiCalls++
		if err != nil {
			return apiCalls, err
		}
		bytes, err := io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return apiCalls, err
		}
		var page struct {
			Data          []UserModel `json:"data"`
			NextPageToken *string     `json:"nextPageToken,omitempty"`
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
	return apiCalls, nil
}

func TryRead(id string) (*UserModel, bool, error) {
	return cache.TryRead(Cache, id)
}
