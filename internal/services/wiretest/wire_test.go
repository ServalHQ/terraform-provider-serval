package wiretest

import (
	"testing"

	"github.com/ServalHQ/terraform-provider-serval/internal/services/access_policy"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/access_policy_approval_procedure"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/access_request"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/app_instance"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/app_resource"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/app_resource_role"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/approval_delegation"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/custom_service"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/group"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/guidance"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/tag"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/team"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/team_user"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/user"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/workflow"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/workflow_approval_procedure"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/workflow_run"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type vals = map[string]tftypes.Value

const (
	groupFixture                         = `{"id":"id-1","name":"eng","userIds":["u-1","u-2"],"organizationId":"org-1","createdAt":"2026-01-01T00:00:00Z"}`
	teamFixture                          = `{"id":"id-1","name":"it","description":"d","prefix":"IT","organizationId":"org-1","createdAt":"2026-01-01T00:00:00Z"}`
	appInstanceFixture                   = `{"id":"id-1","teamId":"team-1","service":"github","externalServiceInstanceId":"acme","name":"gh","accessRequestsEnabled":true}`
	accessPolicyFixture                  = `{"id":"id-1","teamId":"team-1","name":"p","description":"d","maxAccessMinutes":60,"requireBusinessJustification":true}`
	workflowFixture                      = `{"id":"id-1","teamId":"team-1","name":"w","content":"export default {}","executionScope":"TEAM_PRIVATE","isPublished":true,"hasUnpublishedChanges":false}`
	accessPolicyApprovalProcedureFixture = `{"id":"id-1","accessPolicyId":"ap-1","steps":[{"id":"s-1","allowSelfApproval":false,"approvers":[{"user":{"userId":"u-1"},"notify":true}]}]}`
	workflowApprovalProcedureFixture     = `{"id":"id-1","workflowId":"w-1","steps":[{"id":"s-1","allowSelfApproval":true,"approvers":[{"group":{"groupId":"g-1"},"notify":false}]}]}`
	appResourceFixture                   = `{"id":"id-1","appInstanceId":"ai-1","name":"repo","resourceType":"repository","description":"d","externalId":"ext-1"}`
	appResourceRoleFixture               = `{"id":"id-1","resourceId":"res-1","name":"admin","accessPolicyId":"ap-1","requestsEnabled":true,"provisioningMethod":{"manual":{"assignees":[{"assigneeId":"u-1","assigneeType":"MANUAL_PROVISIONING_ASSIGNEE_TYPE_USER"}]}}}`
	teamUserFixture                      = `{"id":"id-1","teamId":"team-1","userId":"u-1","role":"TEAM_USER_ROLE_AGENT","createdAt":"2026-01-01T00:00:00Z"}`
	userFixture                          = `{"id":"id-1","email":"a@example.com","firstName":"A","lastName":"B","role":"USER_ROLE_ORG_MEMBER","createdAt":"2026-01-01T00:00:00Z"}`
	accessRequestFixture                 = `{"id":"id-1","teamId":"team-1","status":"ACCESS_REQUEST_STATUS_PENDING","justification":"j"}`
	workflowRunFixture                   = `{"id":"id-1","teamId":"team-1","workflowId":"w-1","status":"WORKFLOW_RUN_STATUS_COMPLETED"}`
)

func TestResourceWire(t *testing.T) {
	cases := []resourceCase{
		{
			name:    "tag",
			new:     tag.NewResource,
			fixture: `{"id":"id-1","teamId":"team-1","name":"prod","color":"red","iconSlug":"star"}`,
			create:  vals{"team_id": str("team-1"), "name": str("prod"), "color": str("red")},
			update:  vals{"id": str("id-1"), "team_id": str("team-1"), "name": str("prod-2"), "color": str("blue")},
		},
		{
			name:    "guidance",
			new:     guidance.NewResource,
			fixture: `{"id":"id-1","teamId":"team-1","name":"g","content":"c","description":"d","isPublished":true,"shouldAlwaysUse":false,"hasUnpublishedChanges":false}`,
			create:  vals{"team_id": str("team-1"), "name": str("g"), "content": str("c"), "is_published": tftypes.NewValue(tftypes.Bool, true)},
			update:  vals{"id": str("id-1"), "team_id": str("team-1"), "name": str("g2"), "content": str("c2")},
		},
		{
			name:    "custom_service",
			new:     custom_service.NewResource,
			fixture: `{"id":"id-1","teamId":"team-1","name":"svc","domain":"svc.example.com"}`,
			create:  vals{"team_id": str("team-1"), "name": str("svc"), "domain": str("svc.example.com")},
			update:  vals{"id": str("id-1"), "team_id": str("team-1"), "name": str("svc-2"), "domain": str("svc.example.com")},
		},
		{
			name:    "approval_delegation",
			new:     approval_delegation.NewResource,
			fixture: `{"id":"id-1","delegatorUserId":"u-1","delegates":[{"id":"u-2","type":"APPROVAL_DELEGATE_TYPE_USER"}],"description":"d","createdAt":"2026-01-01T00:00:00Z"}`,
			create:  vals{"delegator_user_id": str("u-1"), "description": str("d")},
			update:  vals{"id": str("id-1"), "delegator_user_id": str("u-1"), "description": str("d2")},
		},
		{
			name:    "group",
			new:     group.NewResource,
			fixture: groupFixture,
			create:  vals{"name": str("eng")},
			update:  vals{"id": str("id-1"), "name": str("eng-2")},
		},
		{
			name:    "team",
			new:     team.NewResource,
			fixture: teamFixture,
			create:  vals{"name": str("it"), "description": str("d"), "prefix": str("IT")},
			update:  vals{"id": str("id-1"), "name": str("it-2"), "prefix": str("IT")},
		},
		{
			name:    "app_instance",
			new:     app_instance.NewResource,
			fixture: appInstanceFixture,
			create:  vals{"team_id": str("team-1"), "service": str("github"), "external_service_instance_id": str("acme"), "name": str("gh")},
			update:  vals{"id": str("id-1"), "team_id": str("team-1"), "service": str("github"), "external_service_instance_id": str("acme"), "name": str("gh-2"), "access_requests_enabled": tftypes.NewValue(tftypes.Bool, true)},
		},
		{
			name:    "access_policy",
			new:     access_policy.NewResource,
			fixture: accessPolicyFixture,
			create:  vals{"team_id": str("team-1"), "name": str("p"), "max_access_minutes": tftypes.NewValue(tftypes.Number, 60)},
			update:  vals{"id": str("id-1"), "team_id": str("team-1"), "name": str("p2"), "require_business_justification": tftypes.NewValue(tftypes.Bool, true)},
		},
		{
			name:    "workflow",
			new:     workflow.NewResource,
			fixture: workflowFixture,
			create:  vals{"team_id": str("team-1"), "name": str("w"), "content": str("export default {}"), "execution_scope": str("TEAM_PRIVATE")},
			update:  vals{"id": str("id-1"), "team_id": str("team-1"), "name": str("w2"), "content": str("export default {}"), "is_published": tftypes.NewValue(tftypes.Bool, true)},
		},
		{
			name:     "access_policy_approval_procedure",
			new:      access_policy_approval_procedure.NewResource,
			fixture:  accessPolicyApprovalProcedureFixture,
			create:   vals{"access_policy_id": str("ap-1")},
			update:   vals{"id": str("id-1"), "access_policy_id": str("ap-1")},
			importID: "ap-1/id-1",
		},
		{
			name:     "workflow_approval_procedure",
			new:      workflow_approval_procedure.NewResource,
			fixture:  workflowApprovalProcedureFixture,
			create:   vals{"workflow_id": str("w-1")},
			update:   vals{"id": str("id-1"), "workflow_id": str("w-1")},
			importID: "w-1/id-1",
		},
		{
			name:    "app_resource",
			new:     app_resource.NewResource,
			fixture: appResourceFixture,
			create:  vals{"app_instance_id": str("ai-1"), "name": str("repo"), "resource_type": str("repository")},
			update:  vals{"id": str("id-1"), "app_instance_id": str("ai-1"), "name": str("repo-2"), "resource_type": str("repository")},
		},
		{
			name:    "app_resource_role",
			new:     app_resource_role.NewResource,
			fixture: appResourceRoleFixture,
			create:  vals{"resource_id": str("res-1"), "name": str("admin")},
			update:  vals{"id": str("id-1"), "resource_id": str("res-1"), "name": str("admin-2"), "requests_enabled": tftypes.NewValue(tftypes.Bool, true)},
		},
		{
			name:    "team_user",
			new:     team_user.NewResource,
			fixture: teamUserFixture,
			create:  vals{"team_id": str("team-1"), "user_id": str("u-1"), "role": str("TEAM_USER_ROLE_AGENT")},
			update:  vals{"id": str("id-1"), "team_id": str("team-1"), "user_id": str("u-1"), "role": str("TEAM_USER_ROLE_MANAGER")},
		},
		{
			name:    "user",
			new:     user.NewResource,
			fixture: userFixture,
			create:  vals{"email": str("a@example.com")},
			update:  vals{"id": str("id-1"), "email": str("a@example.com"), "first_name": str("A")},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) { runResource(t, tc) })
	}
}

func TestDataSourceWire(t *testing.T) {
	cases := []dataSourceCase{
		{
			name:          "tag",
			single:        tag.NewTagDataSource,
			list:          tag.NewTagsDataSource,
			fixture:       `{"id":"id-1","teamId":"team-1","name":"prod","color":"red","iconSlug":"star"}`,
			byID:          vals{"id": str("id-1")},
			findBy:        vals{"team_id": str("team-1")},
			listArgs:      vals{"team_id": str("team-1")},
			emptyFindBy:   vals{"team_id": str("")},
			emptyListArgs: vals{"team_id": str("")},
		},
		{
			name:          "guidance",
			single:        guidance.NewGuidanceDataSource,
			list:          guidance.NewGuidancesDataSource,
			fixture:       `{"id":"id-1","teamId":"team-1","name":"g","content":"c","isPublished":true}`,
			byID:          vals{"id": str("id-1")},
			findBy:        vals{"team_id": str("team-1")},
			listArgs:      vals{"team_id": str("team-1")},
			emptyFindBy:   vals{"team_id": str("")},
			emptyListArgs: vals{"team_id": str("")},
		},
		{
			name:          "custom_service",
			single:        custom_service.NewCustomServiceDataSource,
			list:          custom_service.NewCustomServicesDataSource,
			fixture:       `{"id":"id-1","teamId":"team-1","name":"svc","domain":"svc.example.com"}`,
			byID:          vals{"id": str("id-1")},
			findBy:        vals{"team_id": str("team-1")},
			listArgs:      vals{"team_id": str("team-1")},
			emptyFindBy:   vals{"team_id": str("")},
			emptyListArgs: vals{"team_id": str("")},
		},
		{
			name:     "approval_delegation",
			single:   approval_delegation.NewApprovalDelegationDataSource,
			list:     approval_delegation.NewApprovalDelegationsDataSource,
			fixture:  `{"id":"id-1","delegatorUserId":"u-1","delegates":[{"id":"u-2","type":"APPROVAL_DELEGATE_TYPE_USER"}],"description":"d"}`,
			byID:     vals{"id": str("id-1")},
			findBy:   vals{"delegator_user_id": str("u-1")},
			listArgs: vals{"delegator_user_id": str("u-1")},
		},
		{
			name:    "group",
			single:  group.NewGroupDataSource,
			list:    group.NewGroupsDataSource,
			fixture: groupFixture,
			byID:    vals{"id": str("id-1")},
			lookups: []vals{{"name": str("eng")}, {"name": str("missing")}},
		},
		{
			name:    "team",
			single:  team.NewTeamDataSource,
			list:    team.NewTeamsDataSource,
			fixture: teamFixture,
			byID:    vals{"id": str("id-1")},
			lookups: []vals{{"name": str("it")}, {"prefix": str("IT")}, {"prefix": str("missing")}},
		},
		{
			name:          "app_instance",
			single:        app_instance.NewAppInstanceDataSource,
			list:          app_instance.NewAppInstancesDataSource,
			fixture:       appInstanceFixture,
			byID:          vals{"id": str("id-1")},
			findBy:        vals{"team_id": str("team-1")},
			listArgs:      vals{"team_id": str("team-1")},
			emptyFindBy:   vals{"team_id": str("")},
			emptyListArgs: vals{"team_id": str("")},
		},
		{
			name:          "access_policy",
			single:        access_policy.NewAccessPolicyDataSource,
			list:          access_policy.NewAccessPoliciesDataSource,
			fixture:       accessPolicyFixture,
			byID:          vals{"id": str("id-1")},
			findBy:        vals{"team_id": str("team-1")},
			listArgs:      vals{"team_id": str("team-1")},
			emptyFindBy:   vals{"team_id": str("")},
			emptyListArgs: vals{"team_id": str("")},
		},
		{
			name:          "workflow",
			single:        workflow.NewWorkflowDataSource,
			list:          workflow.NewWorkflowsDataSource,
			fixture:       workflowFixture,
			byID:          vals{"id": str("id-1")},
			findBy:        vals{"team_id": str("team-1")},
			listArgs:      vals{"team_id": str("team-1")},
			emptyFindBy:   vals{"team_id": str("")},
			emptyListArgs: vals{"team_id": str("")},
		},
		{
			name:    "access_policy_approval_procedure",
			single:  access_policy_approval_procedure.NewAccessPolicyApprovalProcedureDataSource,
			fixture: accessPolicyApprovalProcedureFixture,
			byID:    vals{"id": str("id-1"), "access_policy_id": str("ap-1")},
		},
		{
			name:    "workflow_approval_procedure",
			single:  workflow_approval_procedure.NewWorkflowApprovalProcedureDataSource,
			fixture: workflowApprovalProcedureFixture,
			byID:    vals{"id": str("id-1"), "workflow_id": str("w-1")},
		},
		{
			name:          "app_resource",
			single:        app_resource.NewAppResourceDataSource,
			list:          app_resource.NewAppResourcesDataSource,
			fixture:       appResourceFixture,
			byID:          vals{"id": str("id-1")},
			findBy:        vals{"team_id": str("team-1"), "app_instance_id": str("ai-1")},
			listArgs:      vals{"team_id": str("team-1"), "app_instance_id": str("ai-1")},
			emptyFindBy:   vals{"team_id": str("")},
			emptyListArgs: vals{"team_id": str("")},
		},
		{
			name:          "app_resource_role",
			single:        app_resource_role.NewAppResourceRoleDataSource,
			list:          app_resource_role.NewAppResourceRolesDataSource,
			fixture:       appResourceRoleFixture,
			byID:          vals{"id": str("id-1")},
			findBy:        vals{"resource_id": str("res-1"), "app_instance_id": str("ai-1")},
			listArgs:      vals{"team_id": str("team-1"), "resource_id": str("res-1")},
			emptyFindBy:   vals{"team_id": str("")},
			emptyListArgs: vals{"team_id": str("")},
		},
		{
			name:          "team_user",
			single:        team_user.NewTeamUserDataSource,
			list:          team_user.NewTeamUsersDataSource,
			fixture:       teamUserFixture,
			byID:          vals{"id": str("id-1"), "team_id": str("team-1"), "user_id": str("u-1")},
			lookups:       []vals{{"id": str("id-1"), "team_id": str("")}},
			listArgs:      vals{"team_id": str("team-1"), "user_id": str("u-1")},
			emptyListArgs: vals{"team_id": str("")},
		},
		{
			name:     "user",
			single:   user.NewUserDataSource,
			list:     user.NewUsersDataSource,
			fixture:  userFixture,
			byID:     vals{"id": str("id-1")},
			lookups:  []vals{{"email": str("a@example.com")}, {"email": str("missing@example.com")}},
			listArgs: vals{"include_deactivated": tftypes.NewValue(tftypes.Bool, true)},
		},
		{
			name:          "access_request",
			single:        access_request.NewAccessRequestDataSource,
			list:          access_request.NewAccessRequestsDataSource,
			fixture:       accessRequestFixture,
			byID:          vals{"id": str("id-1")},
			listArgs:      vals{"team_id": str("team-1")},
			emptyListArgs: vals{"team_id": str("")},
		},
		{
			name:          "workflow_run",
			single:        workflow_run.NewWorkflowRunDataSource,
			list:          workflow_run.NewWorkflowRunsDataSource,
			fixture:       workflowRunFixture,
			byID:          vals{"id": str("id-1")},
			listArgs:      vals{"team_id": str("team-1")},
			emptyListArgs: vals{"team_id": str("")},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) { runDataSources(t, tc) })
	}
}
