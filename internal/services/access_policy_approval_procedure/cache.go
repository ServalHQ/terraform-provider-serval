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
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var Cache *cache.Store[AccessPolicyApprovalProcedureModel]

func Prefetch(ctx context.Context, client *serval.Client, accessPolicyIDs []string) (int, error) {
	Cache = cache.NewStore[AccessPolicyApprovalProcedureModel]()
	apiCalls := 0
	for _, apID := range accessPolicyIDs {
		res := new(http.Response)
		_, err := client.AccessPolicies.ApprovalProcedures.List(ctx, apID,
			option.WithResponseBodyInto(&res),
			option.WithMiddleware(logging.Middleware(ctx)),
		)
		apiCalls++
		if err != nil {
			// Skip access policies that don't support approval procedures
			continue
		}
		bytes, err := io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return apiCalls, err
		}
		var page struct {
			Data []AccessPolicyApprovalProcedureModel `json:"data"`
		}
		if err := apijson.Unmarshal(bytes, &page); err != nil {
			return apiCalls, err
		}
		for i := range page.Data {
			item := page.Data[i]
			item.AccessPolicyID = types.StringValue(apID)
			Cache.Put(item.ID.ValueString(), &item)
		}
	}
	return apiCalls, nil
}

func TryRead(id string) (*AccessPolicyApprovalProcedureModel, bool, error) {
	return cache.TryRead(Cache, id)
}
