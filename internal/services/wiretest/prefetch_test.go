package wiretest

import (
	"context"
	"net/http"
	"testing"

	"github.com/ServalHQ/terraform-provider-serval/internal/cache"
	"github.com/ServalHQ/terraform-provider-serval/internal/client"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/custom_service"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/guidance"
)

func TestPrefetchWire(t *testing.T) {
	cases := []struct {
		name     string
		fixture  string
		prefetch func(context.Context, *client.Client, []string) (int, error)
		lookup   func(string) bool
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
	}
	cache.PrefetchMode = true
	t.Cleanup(func() { cache.PrefetchMode = false })
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := newRecorder(t, tc.fixture)
			r.step("prefetch two pages, then a server error", r.page("page-2"), r.page(""), reply{http.StatusServiceUnavailable, `{}`}, reply{http.StatusServiceUnavailable, `{}`}, reply{http.StatusServiceUnavailable, `{}`})
			n, err := tc.prefetch(r.ctx, r.client, []string{"team-1", "team-2"})
			r.note("calls=%d err=%v cached=%v", n, err, tc.lookup("id-1"))
			r.step("prefetch rejected", reply{http.StatusForbidden, `{"error":"forbidden"}`})
			n, err = tc.prefetch(r.ctx, r.client, []string{"team-1"})
			r.note("calls=%d err=%v", n, err)
			r.compare(tc.name + "_prefetch")
		})
	}
}
