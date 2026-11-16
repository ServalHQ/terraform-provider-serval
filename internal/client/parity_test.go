package client

import (
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/serval-go/option"
	"github.com/stretchr/testify/require"
)

// The cases below send the same calls through serval-go and through Client
// and require identical requests and errors on the wire.

type parityCase struct {
	name    string
	auth    Config
	replies []reply
	sdk     func(ctx context.Context, c *serval.Client) error
	ours    func(ctx context.Context, c *Client) error
}

var tokenReply = reply{status: http.StatusOK, body: `{"access_token":"tok-1","token_type":"bearer","expires_in":3600}`}

func parityCases() []parityCase {
	tagBody := []byte(`{"name":"prod","color":"red","teamId":"team-1"}`)
	return []parityCase{
		{
			name: "create with bearer token",
			auth: Config{BearerToken: "bearer-1"},
			sdk: func(ctx context.Context, c *serval.Client) error {
				res := new(http.Response)
				_, err := c.Tags.New(ctx, serval.TagNewParams{}, option.WithRequestBody("application/json", tagBody), option.WithResponseBodyInto(&res))
				return err
			},
			ours: func(ctx context.Context, c *Client) error {
				_, err := c.Do(ctx, Request{Method: http.MethodPost, Path: "v2/tags", Body: tagBody})
				return err
			},
		},
		{
			name:    "get with client credentials",
			auth:    Config{ClientID: "id-1", ClientSecret: "secret-1"},
			replies: []reply{tokenReply},
			sdk: func(ctx context.Context, c *serval.Client) error {
				res := new(http.Response)
				_, err := c.Tags.Get(ctx, "tag-1", option.WithResponseBodyInto(&res))
				return err
			},
			ours: func(ctx context.Context, c *Client) error {
				_, err := c.Do(ctx, Request{Method: http.MethodGet, Path: "v2/tags/tag-1"})
				return err
			},
		},
		{
			name: "update",
			auth: Config{BearerToken: "bearer-1"},
			sdk: func(ctx context.Context, c *serval.Client) error {
				res := new(http.Response)
				_, err := c.Tags.Update(ctx, "tag-1", serval.TagUpdateParams{}, option.WithRequestBody("application/json", tagBody), option.WithResponseBodyInto(&res))
				return err
			},
			ours: func(ctx context.Context, c *Client) error {
				_, err := c.Do(ctx, Request{Method: http.MethodPut, Path: "v2/tags/tag-1", Body: tagBody})
				return err
			},
		},
		{
			name: "delete",
			auth: Config{BearerToken: "bearer-1"},
			sdk: func(ctx context.Context, c *serval.Client) error {
				_, err := c.Tags.Delete(ctx, "tag-1")
				return err
			},
			ours: func(ctx context.Context, c *Client) error {
				_, err := c.Do(ctx, Request{Method: http.MethodDelete, Path: "v2/tags/tag-1"})
				return err
			},
		},
		{
			name: "list pages",
			auth: Config{BearerToken: "bearer-1"},
			replies: []reply{
				{status: http.StatusOK, body: `{"data":[{"id":"a"}],"nextPageToken":"p2"}`},
				{status: http.StatusOK, body: `{"data":[{"id":"b"}],"nextPageToken":""}`},
			},
			sdk: func(ctx context.Context, c *serval.Client) error {
				page, err := c.Workflows.List(ctx, serval.WorkflowListParams{TeamID: serval.String("team 1"), PageSize: serval.Int(1000)})
				for err == nil && page != nil && len(page.Data) > 0 {
					page, err = page.GetNextPage()
				}
				return err
			},
			ours: func(ctx context.Context, c *Client) error {
				query := map[string][]string{"teamId": {"team 1"}, "pageSize": {"1000"}}
				return c.ListPages(ctx, Request{Path: "v2/workflows", Query: query}, func(Page) (bool, error) { return true, nil })
			},
		},
		{
			name: "retries 503 then succeeds",
			auth: Config{BearerToken: "bearer-1"},
			replies: []reply{
				{status: http.StatusServiceUnavailable, headers: map[string]string{"Retry-After-Ms": "1"}, body: `{}`},
				{status: http.StatusOK, body: `{}`},
			},
			sdk: func(ctx context.Context, c *serval.Client) error {
				_, err := c.Tags.Get(ctx, "tag-1")
				return err
			},
			ours: func(ctx context.Context, c *Client) error {
				_, err := c.Do(ctx, Request{Method: http.MethodGet, Path: "v2/tags/tag-1"})
				return err
			},
		},
		{
			name: "404 is not retried",
			auth: Config{BearerToken: "bearer-1"},
			replies: []reply{
				{status: http.StatusNotFound, body: `{"code":"not_found","message":"tag not found"}`},
			},
			sdk: func(ctx context.Context, c *serval.Client) error {
				_, err := c.Tags.Get(ctx, "tag-1")
				return err
			},
			ours: func(ctx context.Context, c *Client) error {
				_, err := c.Do(ctx, Request{Method: http.MethodGet, Path: "v2/tags/tag-1"})
				return err
			},
		},
	}
}

func TestParityWithServalGo(t *testing.T) {
	ctx := context.Background()
	for _, tc := range parityCases() {
		t.Run(tc.name, func(t *testing.T) {
			sdkServer := newRecorder(t, tc.replies...)
			sdkOpts := []option.RequestOption{option.WithBaseURL(sdkServer.server.URL), option.WithMaxRetries(DefaultMaxRetries)}
			if tc.auth.BearerToken != "" {
				sdkOpts = append(sdkOpts, option.WithBearerToken(tc.auth.BearerToken))
			}
			if tc.auth.ClientID != "" {
				// Unique credentials per test: serval-go caches tokens process-wide.
				sdkOpts = append(sdkOpts, option.WithClientID(tc.auth.ClientID+t.Name()), option.WithClientSecret(tc.auth.ClientSecret))
			}
			sdkClient := serval.NewClient(sdkOpts...)
			sdkErr := tc.sdk(ctx, &sdkClient)

			oursServer := newRecorder(t, tc.replies...)
			cfg := tc.auth
			cfg.BaseURL = oursServer.server.URL
			cfg.MaxRetries = DefaultMaxRetries
			if cfg.ClientID != "" {
				cfg.ClientID += t.Name()
			}
			c, err := New(cfg)
			require.NoError(t, err)
			oursErr := tc.ours(ctx, c)

			require.NotEmpty(t, sdkServer.recorded())
			require.Equal(t, sdkServer.recorded(), oursServer.recorded())
			if sdkErr == nil {
				require.NoError(t, oursErr)
				return
			}
			require.Error(t, oursErr)
			sdkMsg := replaceHost(sdkErr.Error(), sdkServer.server.URL)
			oursMsg := replaceHost(oursErr.Error(), oursServer.server.URL)
			require.Equal(t, sdkMsg, oursMsg)
			var sdkAPIErr *serval.Error
			var oursAPIErr *APIError
			require.ErrorAs(t, sdkErr, &sdkAPIErr)
			require.ErrorAs(t, oursErr, &oursAPIErr)
			require.Equal(t, sdkAPIErr.StatusCode, oursAPIErr.StatusCode)
			body, err := io.ReadAll(oursAPIErr.Response.Body)
			require.NoError(t, err)
			require.Equal(t, string(oursAPIErr.Body), string(body))
		})
	}
}
