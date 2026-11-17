package user

import (
	"context"
	"io"
	"net/http"

	authapi "github.com/ServalHQ/terraform-provider-serval/internal/gen/proto/svauth/public"
	"google.golang.org/protobuf/proto"

	"github.com/ServalHQ/terraform-provider-serval/internal/apijson"
	"github.com/ServalHQ/terraform-provider-serval/internal/cache"
	"github.com/ServalHQ/terraform-provider-serval/internal/client"
	"github.com/ServalHQ/terraform-provider-serval/internal/logging"
)

var Cache *cache.Store[UserModel]

func Prefetch(ctx context.Context, c *client.Client) (int, error) {
	Cache = cache.NewStore[UserModel]()
	apiCalls := 0
	var pageToken *string
	for {
		params := &authapi.ListUsersRequest{
			PageSize:           proto.Int32(1000),
			IncludeDeactivated: proto.Bool(true),
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
			Path:        "v2/users",
			Query:       query,
			Middlewares: []client.Middleware{logging.Middleware(ctx)},
		})
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
