package client

import (
	"context"
	"errors"
	"net/http"
	"net/url"

	"google.golang.org/protobuf/proto"
)

var errMissingID = errors.New("missing required id parameter")

// DoID sends a request to collection/id, rejecting an empty id.
func (c *Client) DoID(ctx context.Context, method, collection, id string, body []byte, mws ...Middleware) (*http.Response, error) {
	if id == "" {
		return nil, errMissingID
	}
	return c.DoIDQuery(ctx, method, collection, id, nil, body, mws...)
}

func (c *Client) DoIDQuery(ctx context.Context, method, collection, id string, query url.Values, body []byte, mws ...Middleware) (*http.Response, error) {
	if id == "" {
		return nil, errMissingID
	}
	return c.Do(ctx, Request{Method: method, Path: collection + "/" + id, Query: query, Body: body, Middlewares: mws})
}

// FirstPage fetches one page of a list endpoint with params encoded as its query.
func (c *Client) FirstPage(ctx context.Context, path string, params proto.Message) (Page, error) {
	query, err := QueryFromMessage(params)
	if err != nil {
		return Page{}, err
	}
	return c.getPage(ctx, Request{Method: http.MethodGet, Path: path, Query: query})
}

// PageError is a failure fetching page Index (0 for the first) of a list.
type PageError struct {
	Index int
	Err   error
}

func (e *PageError) Error() string { return e.Err.Error() }
func (e *PageError) Unwrap() error { return e.Err }

// EachPage calls fn with the raw JSON of each page until fn returns false or the
// list ends. Fetch failures come back as *PageError; fn's errors come back as is.
func (c *Client) EachPage(ctx context.Context, path string, params proto.Message, fn func(raw []byte) (bool, error)) error {
	query, err := QueryFromMessage(params)
	if err != nil {
		return &PageError{Err: err}
	}
	fetched := 0
	var fnErr error
	err = c.ListPages(ctx, Request{Path: path, Query: query}, func(p Page) (bool, error) {
		fetched++
		more, err := fn(p.Raw)
		fnErr = err
		return more, err
	})
	if err != nil && fnErr == nil {
		return &PageError{Index: fetched, Err: err}
	}
	return err
}

// ListErrorSummary is the diagnostic summary list data sources report for an EachPage error.
func ListErrorSummary(err error) string {
	var pageErr *PageError
	switch {
	case !errors.As(err, &pageErr):
		return "failed to unmarshal http request"
	case pageErr.Index == 0:
		return "failed to make http request"
	default:
		return "failed to fetch next page"
	}
}
