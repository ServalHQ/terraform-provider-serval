package client

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	authapi "github.com/ServalHQ/terraform-provider-serval/internal/gen/proto/svauth/public"
	flowapi "github.com/ServalHQ/terraform-provider-serval/internal/gen/proto/svflow/publicapi"
)

func TestQueryFromMessageListParams(t *testing.T) {
	size, tok, team := int32(1000), "page-2", "team-1"
	cases := []struct {
		name string
		want string
		ours proto.Message
	}{
		{"access policies", "pageSize=1000&pageToken=page-2&teamId=team-1",
			&flowapi.ListAccessPoliciesRequest{PageSize: &size, PageToken: &tok, TeamId: team}},
		{"access policy approval procedures by team", "pageSize=1000&pageToken=page-2&teamId=team-1",
			&flowapi.ListTeamAccessPolicyApprovalProceduresRequest{PageSize: &size, PageToken: &tok, TeamId: team}},
		{"access requests", "pageSize=1000&pageToken=page-2&teamId=team-1",
			&flowapi.ListAccessRequestsRequest{PageSize: &size, PageToken: &tok, TeamId: team}},
		{"app instances", "pageSize=1000&pageToken=page-2&teamId=team-1",
			&flowapi.ListAppInstancesRequest{PageSize: &size, PageToken: &tok, TeamId: team}},
		{"app resources", "appInstanceId=ai-1&pageSize=1000&pageToken=page-2&teamId=team-1",
			&flowapi.ListAppResourcesRequest{AppInstanceId: proto.String("ai-1"), PageSize: &size, PageToken: &tok, TeamId: &team}},
		{"app resource roles", "appInstanceId=ai-1&pageSize=1000&pageToken=page-2&resourceId=r-1&teamId=team-1",
			&flowapi.ListAppResourceRolesRequest{AppInstanceId: proto.String("ai-1"), ResourceId: proto.String("r-1"), PageSize: &size, PageToken: &tok, TeamId: &team}},
		{"approval delegations", "delegatorUserId=u-1&pageSize=1000&pageToken=page-2",
			&flowapi.ListApprovalDelegationsRequest{DelegatorUserId: proto.String("u-1"), PageSize: &size, PageToken: &tok}},
		{"custom services", "pageSize=1000&pageToken=page-2&teamId=team-1",
			&flowapi.ListCustomServicesRequest{PageSize: &size, PageToken: &tok, TeamId: team}},
		{"groups", "pageSize=1000&pageToken=page-2",
			&authapi.ListGroupsRequest{PageSize: &size, PageToken: &tok}},
		{"guidances", "pageSize=1000&pageToken=page-2&teamId=team-1",
			&flowapi.ListGuidancesRequest{PageSize: &size, PageToken: &tok, TeamId: team}},
		{"tags", "pageSize=1000&pageToken=page-2&teamId=team-1",
			&flowapi.ListTagsRequest{PageSize: &size, PageToken: &tok, TeamId: team}},
		{"teams", "pageSize=1000&pageToken=page-2",
			&authapi.ListTeamsRequest{PageSize: &size, PageToken: &tok}},
		{"team users", "pageSize=1000&pageToken=page-2&teamId=team-1&userId=u-1",
			&authapi.ListTeamUsersRequest{UserId: proto.String("u-1"), PageSize: &size, PageToken: &tok, TeamId: team}},
		{"users", "includeDeactivated=true&pageSize=1000&pageToken=page-2",
			&authapi.ListUsersRequest{IncludeDeactivated: proto.Bool(true), PageSize: &size, PageToken: &tok}},
		{"users without deactivated", "includeDeactivated=false",
			&authapi.ListUsersRequest{IncludeDeactivated: proto.Bool(false)}},
		{"workflow approval procedures by team", "pageSize=1000&pageToken=page-2&teamId=team-1",
			&flowapi.ListTeamWorkflowApprovalProceduresRequest{PageSize: &size, PageToken: &tok, TeamId: team}},
		{"workflows", "pageSize=1000&pageToken=page-2&teamId=team-1",
			&flowapi.ListWorkflowsRequest{PageSize: &size, PageToken: &tok, TeamId: team}},
		{"workflows first page", "pageSize=1000&teamId=team-1",
			&flowapi.ListWorkflowsRequest{PageSize: &size, TeamId: team}},
		{"workflow runs", "pageSize=1000&pageToken=page-2&teamId=team-1",
			&flowapi.ListWorkflowRunsRequest{PageSize: &size, PageToken: &tok, TeamId: team}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := QueryFromMessage(tc.ours)
			require.NoError(t, err)
			require.Equal(t, tc.want, got.Encode())
		})
	}
}

func TestWithExplicitSendsEmptyTeamID(t *testing.T) {
	got, err := QueryFromMessage(WithExplicit(&flowapi.ListTagsRequest{}, "teamId"))
	require.NoError(t, err)
	require.Equal(t, "teamId=", got.Encode())
}

func TestWithExplicitKeepsSetValue(t *testing.T) {
	got, err := QueryFromMessage(WithExplicit(&flowapi.ListTagsRequest{TeamId: "team-1"}, "teamId"))
	require.NoError(t, err)
	require.Equal(t, "teamId=team-1", got.Encode())
}

func TestWithExplicitRejectsUnknownField(t *testing.T) {
	_, err := QueryFromMessage(WithExplicit(&flowapi.ListTagsRequest{}, "nope"))
	require.ErrorContains(t, err, `no scalar field "nope"`)
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
