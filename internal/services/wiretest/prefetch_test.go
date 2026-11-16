package wiretest

import (
	"context"
	"net/http"
	"testing"

	"github.com/ServalHQ/terraform-provider-serval/internal/cache"
	"github.com/ServalHQ/terraform-provider-serval/internal/client"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/access_policy"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/app_instance"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/custom_service"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/group"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/guidance"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/team"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/workflow"
)

func TestPrefetchWire(t *testing.T) {
	cases := []struct {
		name     string
		fixture  string
		prefetch func(context.Context, *client.Client, []string) (int, error)
		lookup   func(string) bool
		orgWide  bool
	}{
		{
			name:     "guidance",
			fixture:  `{"id":"id-1","teamId":"team-1","name":"g","content":"c"}`,
			prefetch: guidance.Prefetch,
			lookup:   func(id string) bool { _, ok, _ := guidance.TryRead(id); return ok },
		},
		{
			name:     "custom_service",
			fixture:  `{"id":"id-1","teamId":"team-1","name":"svc","domain":"svc.example.com"}`,
			prefetch: custom_service.Prefetch,
			lookup:   func(id string) bool { _, ok, _ := custom_service.TryRead(id); return ok },
		},
		{
			name:     "app_instance",
			fixture:  appInstanceFixture,
			prefetch: app_instance.Prefetch,
			lookup:   func(id string) bool { _, ok, _ := app_instance.TryRead(id); return ok },
		},
		{
			name:     "access_policy",
			fixture:  accessPolicyFixture,
			prefetch: access_policy.Prefetch,
			lookup:   func(id string) bool { _, ok, _ := access_policy.TryRead(id); return ok },
		},
		{
			name:     "workflow",
			fixture:  workflowFixture,
			prefetch: workflow.Prefetch,
			lookup:   func(id string) bool { _, ok, _ := workflow.TryRead(id); return ok },
		},
		{
			name:     "group",
			fixture:  groupFixture,
			prefetch: orgWide(group.Prefetch),
			orgWide:  true,
			lookup:   func(id string) bool { _, ok, _ := group.TryRead(id); return ok },
		},
		{
			name:     "team",
			fixture:  teamFixture,
			prefetch: orgWide(team.Prefetch),
			orgWide:  true,
			lookup:   func(id string) bool { _, ok, _ := team.TryRead(id); return ok },
		},
	}
	cache.PrefetchMode = true
	t.Cleanup(func() { cache.PrefetchMode = false })
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := newRecorder(t, tc.fixture)
			if tc.orgWide {
				r.step("prefetch two pages", r.page("page-2"), r.page(""))
			} else {
				r.step("prefetch two pages, then a server error", r.page("page-2"), r.page(""), reply{http.StatusServiceUnavailable, `{}`}, reply{http.StatusServiceUnavailable, `{}`}, reply{http.StatusServiceUnavailable, `{}`})
			}
			n, err := tc.prefetch(r.ctx, r.client, []string{"team-1", "team-2"})
			r.note("calls=%d err=%v cached=%v", n, err, tc.lookup("id-1"))
			r.step("prefetch rejected", reply{http.StatusForbidden, `{"error":"forbidden"}`})
			n, err = tc.prefetch(r.ctx, r.client, []string{"team-1"})
			r.note("calls=%d err=%v", n, err)
			r.compare(tc.name + "_prefetch")
		})
	}
}

func orgWide(prefetch func(context.Context, *client.Client) (int, error)) func(context.Context, *client.Client, []string) (int, error) {
	return func(ctx context.Context, c *client.Client, _ []string) (int, error) { return prefetch(ctx, c) }
}
