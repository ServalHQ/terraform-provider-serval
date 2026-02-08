package access_policy_approval_procedure

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/serval-go/option"
	"github.com/ServalHQ/terraform-provider-serval/internal/apijson"
	"github.com/ServalHQ/terraform-provider-serval/internal/cache"
	"github.com/ServalHQ/terraform-provider-serval/internal/logging"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var Cache *cache.Store[AccessPolicyApprovalProcedureModel]

func Prefetch(ctx context.Context, client *serval.Client, teamIDs []string) (int, error) {
	Cache = cache.NewStore[AccessPolicyApprovalProcedureModel]()
	apiCalls := 0
	for _, teamID := range teamIDs {
		var pageToken *string
		for {
			params := serval.AccessPolicyApprovalProcedureListByTeamParams{
				PageSize: serval.Int(1000),
				TeamID:   serval.String(teamID),
			}
			if pageToken != nil {
				params.PageToken = serval.String(*pageToken)
			}
			res := new(http.Response)
			_, err := client.AccessPolicies.ApprovalProcedures.ListByTeam(ctx, params,
				option.WithResponseBodyInto(&res),
				option.WithMiddleware(logging.Middleware(ctx)),
			)
			apiCalls++
			if err != nil {
				if cache.IsServerError(err) {
					tflog.Warn(ctx, fmt.Sprintf("prefetch: skipping access_policy_approval_procedures for team %s due to server error: %s", teamID, err))
					break
				}
				return apiCalls, err
			}
			bytes, err := io.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				return apiCalls, err
			}
			var page struct {
				Data          []AccessPolicyApprovalProcedureModel `json:"data"`
				NextPageToken *string                              `json:"nextPageToken,omitempty"`
			}
			if err := apijson.Unmarshal(bytes, &page); err != nil {
				return apiCalls, err
			}
			for i := range page.Data {
				item := page.Data[i]
				Cache.Put(item.ID.ValueString(), &item)
			}
			if page.NextPageToken == nil || *page.NextPageToken == "" {
				break
			}
			pageToken = page.NextPageToken
		}
	}
	return apiCalls, nil
}

func TryRead(id string) (*AccessPolicyApprovalProcedureModel, bool, error) {
	return cache.TryRead(Cache, id)
}
