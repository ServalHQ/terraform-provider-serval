package access_policy_approval_procedure

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

// ByAccessPolicyCache stores AccessPolicyApprovalProcedures keyed by access_policy_id.
var ByAccessPolicyCache *cache.KeyedCache[AccessPolicyApprovalProcedureModel]

// InitCache initializes the access policy approval procedure cache.
func InitCache() {
	ByAccessPolicyCache = cache.NewKeyedCache[AccessPolicyApprovalProcedureModel]()
}

// GetCached retrieves an approval procedure from cache, loading via List API if needed.
func GetCached(
	ctx context.Context,
	client *serval.Client,
	id string,
	accessPolicyID string,
) (*AccessPolicyApprovalProcedureModel, bool, error) {
	if ByAccessPolicyCache == nil {
		return nil, false, nil
	}

	// Check if already in any loaded cache
	if item, _ := ByAccessPolicyCache.FindInLoadedCaches(id); item != nil {
		return item, true, nil
	}

	// Get or create cache for this access policy
	policyCache := ByAccessPolicyCache.GetOrCreateCache(accessPolicyID)

	return policyCache.GetOrLoad(
		id,
		func() (map[string]*AccessPolicyApprovalProcedureModel, error) {
			return fetchAllForAccessPolicy(ctx, client, accessPolicyID)
		},
	)
}

func fetchAllForAccessPolicy(
	ctx context.Context,
	client *serval.Client,
	accessPolicyID string,
) (map[string]*AccessPolicyApprovalProcedureModel, error) {
	result := make(map[string]*AccessPolicyApprovalProcedureModel)

	res := new(http.Response)
	_, err := client.AccessPolicies.ApprovalProcedures.List(ctx, accessPolicyID,
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
		Data []AccessPolicyApprovalProcedureModel `json:"data"`
	}
	if err := apijson.UnmarshalComputed(bytes, &page); err != nil {
		return nil, err
	}

	for i := range page.Data {
		item := page.Data[i]
		result[item.ID.ValueString()] = &item
	}

	return result, nil
}
