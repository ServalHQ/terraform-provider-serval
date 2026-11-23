package workflow_approval_procedure

import (
	"context"
	"errors"
	"net/http"

	"github.com/ServalHQ/terraform-provider-serval/internal/client"
	"github.com/ServalHQ/terraform-provider-serval/internal/logging"
)

func doRequest(ctx context.Context, c *client.Client, method, workflowID, id string, body []byte) (*http.Response, error) {
	if workflowID == "" {
		return nil, errors.New("missing required workflow_id parameter")
	}
	collection := "v2/workflows/" + workflowID + "/approval-procedures"
	if method == http.MethodPost {
		return c.Do(ctx, client.Request{Method: method, Path: collection, Body: body, Middlewares: []client.Middleware{logging.Middleware(ctx)}})
	}
	return c.DoID(ctx, method, collection, id, body, logging.Middleware(ctx))
}
