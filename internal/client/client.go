// Package client is the provider's HTTP transport to the Serval public API.
package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	serval "github.com/ServalHQ/serval-go"
)

const (
	DefaultBaseURL    = "https://public.api.serval.com/"
	DefaultMaxRetries = 2
	tokenPath         = "v2/auth/token"
	maxRetryDelay     = 8 * time.Second
)

type MiddlewareNext = func(*http.Request) (*http.Response, error)

type Middleware = func(*http.Request, MiddlewareNext) (*http.Response, error)

type Config struct {
	// BaseURL defaults to DefaultBaseURL when empty.
	BaseURL      string
	ClientID     string
	ClientSecret string
	// BearerToken takes precedence over ClientID/ClientSecret.
	BearerToken string
	UserAgent   string
	// HTTPClient defaults to http.DefaultClient when nil.
	HTTPClient *http.Client
	// MaxRetries is the number of retries after the first attempt.
	MaxRetries int
	// Legacy serves the service calls that have not moved to Do yet.
	Legacy *serval.Client
}

type Client struct {
	*serval.Client
	baseURL     *url.URL
	bearerToken string
	userAgent   string
	httpClient  *http.Client
	maxRetries  int
	tokens      *tokenSource
}

func New(cfg Config) (*Client, error) {
	base := cfg.BaseURL
	if base == "" {
		base = DefaultBaseURL
	}
	u, err := url.Parse(base)
	if err != nil {
		return nil, fmt.Errorf("client: invalid base url %q: %w", base, err)
	}
	if !strings.HasSuffix(u.Path, "/") {
		u.Path += "/"
	}
	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	c := &Client{
		Client:      cfg.Legacy,
		baseURL:     u,
		bearerToken: cfg.BearerToken,
		userAgent:   cfg.UserAgent,
		httpClient:  httpClient,
		maxRetries:  cfg.MaxRetries,
	}
	if cfg.BearerToken == "" && (cfg.ClientID != "" || cfg.ClientSecret != "") {
		c.tokens = &tokenSource{clientID: cfg.ClientID, clientSecret: cfg.ClientSecret}
	}
	return c, nil
}

// Request describes one API call. Path is relative to the base URL.
type Request struct {
	Method      string
	Path        string
	Query       url.Values
	Body        []byte
	Middlewares []Middleware
}

// Do sends req with auth and retries. The response is returned even when err
// is an *APIError, with its body still readable. On success the caller owns
// and must close the response body.
func (c *Client) Do(ctx context.Context, req Request) (*http.Response, error) {
	u, err := c.resolve(req.Path, req.Query)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequestWithContext(ctx, req.Method, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("client: build request: %w", err)
	}
	if req.Body != nil {
		body := req.Body
		httpReq.ContentLength = int64(len(body))
		httpReq.GetBody = func() (io.ReadCloser, error) { return io.NopCloser(bytes.NewReader(body)), nil }
		httpReq.Body, _ = httpReq.GetBody()
		httpReq.Header.Set("Content-Type", "application/json")
	}
	httpReq.Header.Set("Accept", "application/json")
	if c.userAgent != "" {
		httpReq.Header.Set("User-Agent", c.userAgent)
	}
	if err := c.authorize(ctx, httpReq); err != nil {
		return nil, err
	}

	handler := c.httpClient.Do
	for i := len(req.Middlewares) - 1; i >= 0; i-- {
		mw, next := req.Middlewares[i], handler
		handler = func(r *http.Request) (*http.Response, error) { return mw(r, next) }
	}

	res, err := c.send(ctx, httpReq, handler)
	if err != nil {
		return res, err
	}
	if res.StatusCode >= 400 {
		return res, newAPIError(httpReq, res)
	}
	return res, nil
}

func (c *Client) resolve(path string, query url.Values) (*url.URL, error) {
	u, err := c.baseURL.Parse(strings.TrimLeft(path, "/"))
	if err != nil {
		return nil, fmt.Errorf("client: invalid path %q: %w", path, err)
	}
	if len(query) > 0 {
		u.RawQuery = query.Encode()
	}
	return u, nil
}

func (c *Client) authorize(ctx context.Context, req *http.Request) error {
	if c.bearerToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.bearerToken)
		return nil
	}
	if c.tokens == nil {
		return nil
	}
	token, err := c.tokens.token(ctx, c.httpClient, c.baseURL)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	return nil
}

func (c *Client) send(ctx context.Context, req *http.Request, handler MiddlewareNext) (*http.Response, error) {
	var res *http.Response
	var err error
	for attempt := 0; ; attempt++ {
		attemptReq := req.Clone(ctx)
		if req.GetBody != nil {
			attemptReq.Body, err = req.GetBody()
			if err != nil {
				return nil, fmt.Errorf("client: rewind request body: %w", err)
			}
		}
		res, err = handler(attemptReq)
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
		if attempt >= c.maxRetries || !shouldRetry(res) {
			break
		}
		if res != nil && res.Body != nil {
			_ = res.Body.Close()
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(retryDelay(res, attempt)):
		}
	}
	return res, err
}

func shouldRetry(res *http.Response) bool {
	if res == nil {
		return true
	}
	switch res.Header.Get("x-should-retry") {
	case "true":
		return true
	case "false":
		return false
	}
	return res.StatusCode == http.StatusRequestTimeout ||
		res.StatusCode == http.StatusConflict ||
		res.StatusCode == http.StatusTooManyRequests ||
		res.StatusCode >= http.StatusInternalServerError
}

func retryDelay(res *http.Response, attempt int) time.Duration {
	if d, ok := retryAfter(res); ok {
		return max(0, d)
	}
	delay := min(time.Duration(0.5*float64(time.Second)*math.Pow(2, float64(attempt))), maxRetryDelay)
	return delay - time.Duration(rand.Int63n(int64(delay/4)))
}

func retryAfter(res *http.Response) (time.Duration, bool) {
	if res == nil {
		return 0, false
	}
	if v := res.Header.Get("Retry-After-Ms"); v != "" {
		if ms, err := strconv.ParseFloat(v, 64); err == nil {
			return time.Duration(ms * float64(time.Millisecond)), true
		}
	}
	if v := res.Header.Get("Retry-After"); v != "" {
		if s, err := strconv.ParseFloat(v, 64); err == nil {
			return time.Duration(s * float64(time.Second)), true
		}
		if t, err := time.Parse(time.RFC1123, v); err == nil {
			return time.Until(t), true
		}
	}
	return 0, false
}
