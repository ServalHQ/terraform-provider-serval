package internal

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/ServalHQ/serval-go/option"
	"github.com/ServalHQ/terraform-provider-serval/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestConfigureSharesSettingsBetweenLegacyAndProviderClient(t *testing.T) {
	type seen struct{ path, auth string }
	var (
		mu  sync.Mutex
		got []seen
	)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		got = append(got, seen{r.URL.Path, r.Header.Get("Authorization")})
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data":{"id":"t1"}}`))
	}))
	defer srv.Close()
	for _, env := range []string{"SERVAL_BASE_URL", "SERVAL_CLIENT_ID", "SERVAL_CLIENT_SECRET", "SERVAL_BEARER_TOKEN"} {
		t.Setenv(env, "")
	}
	t.Setenv("SERVAL_BEARER_TOKEN", "env-token")

	ctx := context.Background()
	p := &ServalProvider{}
	var schemaResp provider.SchemaResponse
	p.Schema(ctx, provider.SchemaRequest{}, &schemaResp)
	objType := schemaResp.Schema.Type().TerraformType(ctx).(tftypes.Object)
	vals := map[string]tftypes.Value{}
	for name, typ := range objType.AttributeTypes {
		vals[name] = tftypes.NewValue(typ, nil)
	}
	vals["base_url"] = tftypes.NewValue(tftypes.String, srv.URL+"/")

	var resp provider.ConfigureResponse
	p.Configure(ctx, provider.ConfigureRequest{Config: tfsdk.Config{
		Schema: schemaResp.Schema,
		Raw:    tftypes.NewValue(objType, vals),
	}}, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("configure: %v", resp.Diagnostics)
	}
	c, ok := resp.ResourceData.(*client.Client)
	if !ok || resp.DataSourceData != resp.ResourceData {
		t.Fatalf("provider data = %T / %T, want the same *client.Client", resp.ResourceData, resp.DataSourceData)
	}

	res := new(http.Response)
	if _, err := c.Tags.Get(ctx, "t1", option.WithResponseBodyInto(&res)); err != nil {
		t.Fatalf("legacy get: %v", err)
	}
	if _, err := c.Do(ctx, client.Request{Method: http.MethodGet, Path: "v2/tags/t1"}); err != nil {
		t.Fatalf("client get: %v", err)
	}

	mu.Lock()
	defer mu.Unlock()
	want := seen{"/v2/tags/t1", "Bearer env-token"}
	if len(got) != 2 || got[0] != want || got[1] != want {
		t.Fatalf("requests = %+v, want two of %+v", got, want)
	}
}

func TestConfigureRejectsEmptyBaseURL(t *testing.T) {
	for _, env := range []string{"SERVAL_BASE_URL", "SERVAL_CLIENT_ID", "SERVAL_CLIENT_SECRET", "SERVAL_BEARER_TOKEN"} {
		t.Setenv(env, "")
	}
	t.Setenv("SERVAL_BEARER_TOKEN", "env-token")

	ctx := context.Background()
	p := &ServalProvider{}
	var schemaResp provider.SchemaResponse
	p.Schema(ctx, provider.SchemaRequest{}, &schemaResp)
	objType := schemaResp.Schema.Type().TerraformType(ctx).(tftypes.Object)
	vals := map[string]tftypes.Value{}
	for name, typ := range objType.AttributeTypes {
		vals[name] = tftypes.NewValue(typ, nil)
	}

	var resp provider.ConfigureResponse
	p.Configure(ctx, provider.ConfigureRequest{Config: tfsdk.Config{
		Schema: schemaResp.Schema,
		Raw:    tftypes.NewValue(objType, vals),
	}}, &resp)
	if !resp.Diagnostics.HasError() || resp.ResourceData != nil {
		t.Fatalf("configure with SERVAL_BASE_URL=\"\" = %v, data %T; want an error and no client", resp.Diagnostics, resp.ResourceData)
	}
}
