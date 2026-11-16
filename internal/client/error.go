package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// APIError is returned for responses with status >= 400.
type APIError struct {
	StatusCode int
	Request    *http.Request
	Response   *http.Response
	Body       []byte
}

func newAPIError(req *http.Request, res *http.Response) error {
	body, err := io.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		return fmt.Errorf("client: read error response: %w", err)
	}
	res.Body = io.NopCloser(bytes.NewReader(body))
	return &APIError{StatusCode: res.StatusCode, Request: req, Response: res, Body: body}
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%s %q: %d %s %s", e.Request.Method, e.Request.URL, e.StatusCode, http.StatusText(e.StatusCode), e.Body)
}
