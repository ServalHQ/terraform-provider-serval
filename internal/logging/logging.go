package logging

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ServalHQ/serval-go/option"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LoggingEnabled controls whether request/response bodies are logged.
// Disabled by default because reading the full response body causes hangs
// with Go's HTTP transport under high concurrency.
var LoggingEnabled = false

func Middleware(ctx context.Context) option.Middleware {
	return func(req *http.Request, next option.MiddlewareNext) (*http.Response, error) {
		// Skip body logging to prevent HTTP transport hangs
		if !LoggingEnabled {
			return next(req)
		}

		if req != nil {
			if err := LogRequest(ctx, req); err != nil {
				return nil, err
			}
		}

		resp, err := next(req)

		if resp != nil {
			if err := LogResponse(ctx, resp); err != nil {
				return nil, err
			}
		}

		return resp, err
	}
}

func LogRequest(ctx context.Context, req *http.Request) error {
	fmt.Println("teddysanity v1.0.4 - LogRequest")
	lines := []string{
		fmt.Sprintf("\n%s %s %s", req.Method, req.URL.Path, req.Proto),
	}

	// Log headers
	for name, values := range req.Header {
		for _, value := range values {
			lines = append(
				lines,
				fmt.Sprintf("> %s: %s", strings.ToLower(name), value),
			)
		}
	}

	if req.Body != nil {
		// Read the body without mutating the original response
		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			return err
		}

		// Restore the original body to the response so it can be read again
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// Log the body
		lines = append(lines, ">\n", string(bodyBytes), "\n")
	}

	tflog.Debug(ctx, strings.Join(lines, "\n"))

	return nil
}

func LogResponse(ctx context.Context, resp *http.Response) error {
	// Log the status code
	lines := []string{fmt.Sprintf("\n< %s %s", resp.Proto, resp.Status)}

	// Log headers
	for name, values := range resp.Header {
		for _, value := range values {
			lines = append(
				lines,
				fmt.Sprintf("< %s: %s", strings.ToLower(name), value),
			)
		}
	}

	// Read the body and close the original to properly release the connection
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		return err
	}
	// Close the original body to signal connection can be reused
	resp.Body.Close()

	// Restore a new body for downstream consumers
	resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	lines = append(lines, "<\n", string(bodyBytes), "\n")

	// Log the body
	tflog.Debug(ctx, strings.Join(lines, "\n"))

	return nil
}
