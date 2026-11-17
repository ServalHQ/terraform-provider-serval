package access_policy_approval_procedure

import (
	"context"
	"errors"
	"net/http"

	"github.com/ServalHQ/terraform-provider-serval/internal/client"
	"github.com/ServalHQ/terraform-provider-serval/internal/logging"
)

func doRequest(ctx context.Context, c *client.Client, method, accessPolicyID, id string, body []byte) (*http.Response, error) {
	if accessPolicyID == "" {
		return nil, errors.New("missing required access_policy_id parameter")
	}
	collection := "v2/access-policies/" + accessPolicyID + "/approval-procedures"
	if method == http.MethodPost {
		return c.Do(ctx, client.Request{Method: method, Path: collection, Body: body, Middlewares: []client.Middleware{logging.Middleware(ctx)}})
	}
	return c.DoID(ctx, method, collection, id, body, logging.Middleware(ctx))
}
