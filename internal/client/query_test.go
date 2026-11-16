package client

import (
	"net/url"
	"testing"

	"github.com/ServalHQ/serval-go"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	authapi "github.com/ServalHQ/terraform-provider-serval/internal/gen/proto/svauth/public"
	flowapi "github.com/ServalHQ/terraform-provider-serval/internal/gen/proto/svflow/publicapi"
)

type sdkQuery interface {
	URLQuery() (url.Values, error)
}

func TestQueryFromMessageMatchesServalGoListParams(t *testing.T) {
	size, tok, team := int32(1000), "page-2", "team-1"
	cases := []struct {
		name string
		sdk  sdkQuery
		ours proto.Message
	}{
		{"access policies", serval.AccessPolicyListParams{PageSize: serval.Int(1000), PageToken: serval.String(tok), TeamID: serval.String(team)},
			&flowapi.ListAccessPoliciesRequest{PageSize: &size, PageToken: &tok, TeamId: team}},
		{"access policy approval procedures by team", serval.AccessPolicyApprovalProcedureListByTeamParams{PageSize: serval.Int(1000), PageToken: serval.String(tok), TeamID: serval.String(team)},
			&flowapi.ListTeamAccessPolicyApprovalProceduresRequest{PageSize: &size, PageToken: &tok, TeamId: team}},
		{"access requests", serval.AccessRequestListParams{PageSize: serval.Int(1000), PageToken: serval.String(tok), TeamID: serval.String(team)},
			&flowapi.ListAccessRequestsRequest{PageSize: &size, PageToken: &tok, TeamId: team}},
		{"app instances", serval.AppInstanceListParams{PageSize: serval.Int(1000), PageToken: serval.String(tok), TeamID: serval.String(team)},
			&flowapi.ListAppInstancesRequest{PageSize: &size, PageToken: &tok, TeamId: team}},
		{"app resources", serval.AppResourceListParams{AppInstanceID: serval.String("ai-1"), PageSize: serval.Int(1000), PageToken: serval.String(tok), TeamID: serval.String(team)},
			&flowapi.ListAppResourcesRequest{AppInstanceId: proto.String("ai-1"), PageSize: &size, PageToken: &tok, TeamId: &team}},
		{"app resource roles", serval.AppResourceRoleListParams{AppInstanceID: serval.String("ai-1"), ResourceID: serval.String("r-1"), PageSize: serval.Int(1000), PageToken: serval.String(tok), TeamID: serval.String(team)},
			&flowapi.ListAppResourceRolesRequest{AppInstanceId: proto.String("ai-1"), ResourceId: proto.String("r-1"), PageSize: &size, PageToken: &tok, TeamId: &team}},
		{"approval delegations", serval.ApprovalDelegationListParams{DelegatorUserID: serval.String("u-1"), PageSize: serval.Int(1000), PageToken: serval.String(tok)},
			&flowapi.ListApprovalDelegationsRequest{DelegatorUserId: proto.String("u-1"), PageSize: &size, PageToken: &tok}},
		{"custom services", serval.CustomServiceListParams{PageSize: serval.Int(1000), PageToken: serval.String(tok), TeamID: serval.String(team)},
			&flowapi.ListCustomServicesRequest{PageSize: &size, PageToken: &tok, TeamId: team}},
		{"groups", serval.GroupListParams{PageSize: serval.Int(1000), PageToken: serval.String(tok)},
			&authapi.ListGroupsRequest{PageSize: &size, PageToken: &tok}},
		{"guidances", serval.GuidanceListParams{PageSize: serval.Int(1000), PageToken: serval.String(tok), TeamID: serval.String(team)},
			&flowapi.ListGuidancesRequest{PageSize: &size, PageToken: &tok, TeamId: team}},
		{"tags", serval.TagListParams{PageSize: serval.Int(1000), PageToken: serval.String(tok), TeamID: serval.String(team)},
			&flowapi.ListTagsRequest{PageSize: &size, PageToken: &tok, TeamId: team}},
		{"teams", serval.TeamListParams{PageSize: serval.Int(1000), PageToken: serval.String(tok)},
			&authapi.ListTeamsRequest{PageSize: &size, PageToken: &tok}},
		{"team users", serval.TeamUserListParams{UserID: serval.String("u-1"), PageSize: serval.Int(1000), PageToken: serval.String(tok), TeamID: serval.String(team)},
			&authapi.ListTeamUsersRequest{UserId: proto.String("u-1"), PageSize: &size, PageToken: &tok, TeamId: team}},
		{"users", serval.UserListParams{IncludeDeactivated: serval.Bool(true), PageSize: serval.Int(1000), PageToken: serval.String(tok)},
			&authapi.ListUsersRequest{IncludeDeactivated: proto.Bool(true), PageSize: &size, PageToken: &tok}},
		{"users without deactivated", serval.UserListParams{IncludeDeactivated: serval.Bool(false)},
			&authapi.ListUsersRequest{IncludeDeactivated: proto.Bool(false)}},
		{"workflow approval procedures by team", serval.WorkflowApprovalProcedureListByTeamParams{PageSize: serval.Int(1000), PageToken: serval.String(tok), TeamID: serval.String(team)},
			&flowapi.ListTeamWorkflowApprovalProceduresRequest{PageSize: &size, PageToken: &tok, TeamId: team}},
		{"workflows", serval.WorkflowListParams{PageSize: serval.Int(1000), PageToken: serval.String(tok), TeamID: serval.String(team)},
			&flowapi.ListWorkflowsRequest{PageSize: &size, PageToken: &tok, TeamId: team}},
		{"workflows first page", serval.WorkflowListParams{PageSize: serval.Int(1000), TeamID: serval.String(team)},
			&flowapi.ListWorkflowsRequest{PageSize: &size, TeamId: team}},
		{"workflow runs", serval.WorkflowRunListParams{PageSize: serval.Int(1000), PageToken: serval.String(tok), TeamID: serval.String(team)},
			&flowapi.ListWorkflowRunsRequest{PageSize: &size, PageToken: &tok, TeamId: team}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			want, err := tc.sdk.URLQuery()
			require.NoError(t, err)
			got, err := QueryFromMessage(tc.ours)
			require.NoError(t, err)
			require.Equal(t, want.Encode(), got.Encode())
		})
	}
}

func TestQueryFromMessageJoinsRepeatedFields(t *testing.T) {
	got, err := QueryFromMessage(&authapi.ListUsersRequest{Emails: []string{"a@x.com", "b@x.com"}})
	require.NoError(t, err)
	require.Equal(t, url.Values{"emails": {"a@x.com,b@x.com"}}, got)
}

func TestQueryFromMessageRejectsEnumFields(t *testing.T) {
	_, err := QueryFromMessage(&flowapi.ListSubjectTagsRequest{SubjectType: 1, SubjectId: "s-1"})
	require.ErrorContains(t, err, "cannot be sent as query parameters")
}
