package client

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func replaceHost(s, baseURL string) string {
	return strings.ReplaceAll(s, baseURL, "http://host")
}

func TestNewDefaultsBaseURL(t *testing.T) {
	c, err := New(Config{})
	require.NoError(t, err)
	require.Equal(t, DefaultBaseURL, c.baseURL.String())
}

func TestBaseURLPathIsKept(t *testing.T) {
	rec := newRecorder(t)
	c, err := New(Config{BaseURL: rec.server.URL + "/prefix", BearerToken: "b"})
	require.NoError(t, err)
	res, err := c.Do(context.Background(), Request{Method: http.MethodGet, Path: "v2/tags/t"})
	require.NoError(t, err)
	res.Body.Close()
	require.Equal(t, "/prefix/v2/tags/t", rec.recorded()[0].Path)
}

func TestOAuthTokenIsCachedPerClient(t *testing.T) {
	rec := newRecorder(t, tokenReply, reply{status: 200, body: `{}`}, reply{status: 200, body: `{}`}, tokenReply)
	ctx := context.Background()
	cfg := Config{BaseURL: rec.server.URL, ClientID: "id", ClientSecret: "secret"}

	first, err := New(cfg)
	require.NoError(t, err)
	for range 2 {
		res, err := first.Do(ctx, Request{Method: http.MethodGet, Path: "v2/tags/t"})
		require.NoError(t, err)
		res.Body.Close()
	}
	second, err := New(cfg)
	require.NoError(t, err)
	_, err = second.tokens.token(ctx, http.DefaultClient, second.baseURL)
	require.NoError(t, err)

	var paths []string
	for _, r := range rec.recorded() {
		paths = append(paths, r.Method+" "+r.Path)
	}
	require.Equal(t, []string{"POST /v2/auth/token", "GET /v2/tags/t", "GET /v2/tags/t", "POST /v2/auth/token"}, paths)
	require.Equal(t, "Basic aWQ6c2VjcmV0", rec.recorded()[0].Authorization)
	require.Equal(t, "grant_type=client_credentials", rec.recorded()[0].Body)
}

func TestOAuthTokenFailure(t *testing.T) {
	rec := newRecorder(t, reply{status: http.StatusUnauthorized, body: `{}`})
	c, err := New(Config{BaseURL: rec.server.URL, ClientID: "id", ClientSecret: "secret"})
	require.NoError(t, err)
	_, err = c.Do(context.Background(), Request{Method: http.MethodGet, Path: "v2/tags/t"})
	require.ErrorContains(t, err, "OAuth2 token request returned status 401")
	require.Len(t, rec.recorded(), 1)
}

func TestRetriesStopAtMax(t *testing.T) {
	busy := reply{status: http.StatusTooManyRequests, headers: map[string]string{"Retry-After": "0"}, body: `{}`}
	rec := newRecorder(t, busy, busy, busy, busy)
	c, err := New(Config{BaseURL: rec.server.URL, MaxRetries: 2})
	require.NoError(t, err)
	res, err := c.Do(context.Background(), Request{Method: http.MethodPost, Path: "v2/tags", Body: []byte(`{"a":1}`)})
	var apiErr *APIError
	require.True(t, errors.As(err, &apiErr))
	require.Equal(t, http.StatusTooManyRequests, res.StatusCode)
	require.Len(t, rec.recorded(), 3)
	for _, r := range rec.recorded() {
		require.Equal(t, `{"a":1}`, r.Body)
	}
}

func TestShouldRetryHeaderOverridesStatus(t *testing.T) {
	rec := newRecorder(t, reply{status: http.StatusServiceUnavailable, headers: map[string]string{"x-should-retry": "false"}, body: `{}`})
	c, err := New(Config{BaseURL: rec.server.URL, MaxRetries: 2})
	require.NoError(t, err)
	_, err = c.Do(context.Background(), Request{Method: http.MethodGet, Path: "v2/tags/t"})
	require.Error(t, err)
	require.Len(t, rec.recorded(), 1)
}

func TestMiddlewareRunsInOrderPerAttempt(t *testing.T) {
	rec := newRecorder(t, reply{status: 500, headers: map[string]string{"Retry-After-Ms": "0"}, body: `{}`}, reply{status: 200, body: `{}`})
	c, err := New(Config{BaseURL: rec.server.URL, MaxRetries: 1})
	require.NoError(t, err)
	var calls []string
	mw := func(name string) Middleware {
		return func(r *http.Request, next MiddlewareNext) (*http.Response, error) {
			calls = append(calls, name)
			return next(r)
		}
	}
	res, err := c.Do(context.Background(), Request{Method: http.MethodGet, Path: "v2/tags/t", Middlewares: []Middleware{mw("outer"), mw("inner")}})
	require.NoError(t, err)
	res.Body.Close()
	require.Equal(t, []string{"outer", "inner", "outer", "inner"}, calls)
}

func TestListPagesStopsWhenCallbackDeclines(t *testing.T) {
	rec := newRecorder(t, reply{status: 200, body: `{"data":[{"id":"a"}],"nextPageToken":"p2"}`})
	c, err := New(Config{BaseURL: rec.server.URL})
	require.NoError(t, err)
	var pages int
	err = c.ListPages(context.Background(), Request{Path: "v2/tags"}, func(p Page) (bool, error) {
		pages++
		require.Equal(t, 1, p.ItemCount)
		return false, nil
	})
	require.NoError(t, err)
	require.Equal(t, 1, pages)
	require.Len(t, rec.recorded(), 1)
}
