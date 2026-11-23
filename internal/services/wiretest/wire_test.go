package wiretest

import (
	"testing"

	"github.com/ServalHQ/terraform-provider-serval/internal/services/approval_delegation"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/custom_service"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/guidance"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/tag"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type vals = map[string]tftypes.Value

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
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) { runDataSources(t, tc) })
	}
}
