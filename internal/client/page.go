package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Page is one page of a cursor-paginated list response.
type Page struct {
	Raw           []byte
	ItemCount     int
	NextPageToken string
}

// ListPages calls fn for each page of req until fn returns false, a page is
// empty, or there is no next page token.
func (c *Client) ListPages(ctx context.Context, req Request, fn func(Page) (bool, error)) error {
	query := url.Values{}
	for k, v := range req.Query {
		query[k] = v
	}
	req.Method = http.MethodGet
	for {
		req.Query = query
		page, err := c.getPage(ctx, req)
		if err != nil {
			return err
		}
		if page.ItemCount == 0 {
			return nil
		}
		more, err := fn(page)
		if err != nil || !more || page.NextPageToken == "" {
			return err
		}
		query.Set("pageToken", page.NextPageToken)
	}
}

func (c *Client) getPage(ctx context.Context, req Request) (Page, error) {
	res, err := c.Do(ctx, req)
	if err != nil {
		return Page{}, err
	}
	defer res.Body.Close()
	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return Page{}, fmt.Errorf("client: read page: %w", err)
	}
	var envelope struct {
		Data          []json.RawMessage `json:"data"`
		NextPageToken string            `json:"nextPageToken"`
	}
	if err := json.Unmarshal(raw, &envelope); err != nil {
		return Page{}, fmt.Errorf("client: parse page: %w", err)
	}
	return Page{Raw: raw, ItemCount: len(envelope.Data), NextPageToken: envelope.NextPageToken}, nil
}
