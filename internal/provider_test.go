package internal

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ServalHQ/terraform-provider-serval/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestConfigurePassesSettingsToClient(t *testing.T) {
	type seen struct{ path, auth string }
	var got []seen
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = append(got, seen{r.URL.Path, r.Header.Get("Authorization")})
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

	if _, err := c.Do(ctx, client.Request{Method: http.MethodGet, Path: "v2/tags/t1"}); err != nil {
		t.Fatalf("client get: %v", err)
	}

	want := seen{"/v2/tags/t1", "Bearer env-token"}
	if len(got) != 1 || got[0] != want {
		t.Fatalf("requests = %+v, want one %+v", got, want)
	}
}
