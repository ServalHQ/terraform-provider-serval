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
)

// ByWorkflowCache stores WorkflowApprovalProcedures keyed by workflow_id.
var ByWorkflowCache *cache.KeyedCache[WorkflowApprovalProcedureModel]

// InitCache initializes the workflow approval procedure cache.
func InitCache() {
	ByWorkflowCache = cache.NewKeyedCache[WorkflowApprovalProcedureModel]()
}

// GetCached retrieves an approval procedure from cache, loading via List API if needed.
func GetCached(
	ctx context.Context,
	client *serval.Client,
	id string,
	workflowID string,
) (*WorkflowApprovalProcedureModel, bool, error) {
	if ByWorkflowCache == nil {
		return nil, false, nil
	}

	// Check if already in any loaded cache
	if item, _ := ByWorkflowCache.FindInLoadedCaches(id); item != nil {
		return item, true, nil
	}

	// Get or create cache for this workflow
	workflowCache := ByWorkflowCache.GetOrCreateCache(workflowID)

	return workflowCache.GetOrLoad(
		id,
		func() (map[string]*WorkflowApprovalProcedureModel, error) {
			return fetchAllForWorkflow(ctx, client, workflowID)
		},
	)
}

func fetchAllForWorkflow(
	ctx context.Context,
	client *serval.Client,
	workflowID string,
) (map[string]*WorkflowApprovalProcedureModel, error) {
	result := make(map[string]*WorkflowApprovalProcedureModel)

	res := new(http.Response)
	_, err := client.Workflows.ApprovalProcedures.List(ctx, workflowID,
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	var page struct {
		Data []WorkflowApprovalProcedureModel `json:"data"`
	}
	if err := apijson.Unmarshal(bytes, &page); err != nil {
		return nil, err
	}

	for i := range page.Data {
		item := page.Data[i]
		result[item.ID.ValueString()] = &item
	}

	return result, nil
}
