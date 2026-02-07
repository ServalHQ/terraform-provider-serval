package workflow_approval_procedure

import (
	"context"
	"io"
	"net/http"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/serval-go/option"
	"github.com/ServalHQ/terraform-provider-serval/internal/apijson"
	"github.com/ServalHQ/terraform-provider-serval/internal/cache"
	"github.com/ServalHQ/terraform-provider-serval/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var Cache *cache.Store[WorkflowApprovalProcedureModel]

func Prefetch(ctx context.Context, client *serval.Client, workflowIDs []string) (int, error) {
	Cache = cache.NewStore[WorkflowApprovalProcedureModel]()
	apiCalls := 0
	for _, wfID := range workflowIDs {
		res := new(http.Response)
		_, err := client.Workflows.ApprovalProcedures.List(ctx, wfID,
			option.WithResponseBodyInto(&res),
			option.WithMiddleware(logging.Middleware(ctx)),
		)
		apiCalls++
		if err != nil {
			// Skip workflows that don't support approval procedures
			continue
		}
		bytes, err := io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return apiCalls, err
		}
		var page struct {
			Data []WorkflowApprovalProcedureModel `json:"data"`
		}
		if err := apijson.Unmarshal(bytes, &page); err != nil {
			return apiCalls, err
		}
		for i := range page.Data {
			item := page.Data[i]
			item.WorkflowID = types.StringValue(wfID)
			Cache.Put(item.ID.ValueString(), &item)
		}
	}
	return apiCalls, nil
}

func TryRead(id string) (*WorkflowApprovalProcedureModel, bool, error) {
	return cache.TryRead(Cache, id)
}
