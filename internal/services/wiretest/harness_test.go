package wiretest

import (
	"cmp"
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/ServalHQ/terraform-provider-serval/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var update = flag.Bool("update", false, "rewrite testdata/*.golden")

type reply struct {
	status int
	body   string
}

type recorder struct {
	t       *testing.T
	srv     *httptest.Server
	mu      sync.Mutex
	queue   []reply
	log     strings.Builder
	client  *client.Client
	ctx     context.Context
	fixture string
}

func newRecorder(t *testing.T, fixture string) *recorder {
	r := &recorder{t: t, ctx: context.Background(), fixture: fixture}
	r.srv = httptest.NewServer(http.HandlerFunc(r.serve))
	t.Cleanup(r.srv.Close)
	for _, env := range []string{"SERVAL_BASE_URL", "SERVAL_CLIENT_ID", "SERVAL_CLIENT_SECRET", "SERVAL_BEARER_TOKEN"} {
		t.Setenv(env, "")
	}
	c, err := client.New(client.Config{BaseURL: r.srv.URL + "/", BearerToken: "tok", MaxRetries: client.DefaultMaxRetries})
	if err != nil {
		t.Fatal(err)
	}
	r.client = c
	return r
}

func (r *recorder) serve(w http.ResponseWriter, req *http.Request) {
	body, _ := io.ReadAll(req.Body)
	r.mu.Lock()
	defer r.mu.Unlock()
	fmt.Fprintf(&r.log, "> %s %s", req.Method, req.URL.Path)
	if req.URL.RawQuery != "" {
		fmt.Fprintf(&r.log, "?%s", req.URL.RawQuery)
	}
	fmt.Fprintf(&r.log, "\n  auth=%q content-type=%q\n", req.Header.Get("Authorization"), req.Header.Get("Content-Type"))
	if len(body) > 0 {
		fmt.Fprintf(&r.log, "  body=%s\n", body)
	}
	if len(r.queue) == 0 {
		r.t.Errorf("unexpected request %s %s", req.Method, req.URL)
		w.WriteHeader(http.StatusTeapot)
		return
	}
	rep := r.queue[0]
	r.queue = r.queue[1:]
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(rep.status)
	_, _ = w.Write([]byte(rep.body))
}

func (r *recorder) step(name string, replies ...reply) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if len(r.queue) != 0 {
		r.t.Errorf("unused replies before %q: %v", name, r.queue)
	}
	r.queue = replies
	fmt.Fprintf(&r.log, "\n## %s\n", name)
}

func (r *recorder) item() reply { return reply{http.StatusOK, `{"data":` + r.fixture + `}`} }

func (r *recorder) page(token string) reply {
	return reply{http.StatusOK, `{"data":[` + r.fixture + `],"nextPageToken":"` + token + `"}`}
}

func (r *recorder) outcome(diags diag.Diagnostics, state tftypes.Value) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, d := range diags {
		fmt.Fprintf(&r.log, "  %s: %s: %s\n", d.Severity(), d.Summary(), strings.ReplaceAll(d.Detail(), r.srv.URL, "SERVER"))
	}
	fmt.Fprintf(&r.log, "  state=%s\n", state)
}

func (r *recorder) compare(name string) {
	got := r.log.String()
	path := filepath.Join("testdata", name+".golden")
	if *update {
		if err := os.WriteFile(path, []byte(got), 0o644); err != nil {
			r.t.Fatal(err)
		}
		return
	}
	want, err := os.ReadFile(path)
	if err != nil {
		r.t.Fatal(err)
	}
	if got != string(want) {
		r.t.Errorf("wire log differs from %s:\n%s", path, got)
	}
}

func objectWith(ctx context.Context, typ tftypes.Type, set map[string]tftypes.Value) tftypes.Value {
	obj := typ.(tftypes.Object)
	vals := map[string]tftypes.Value{}
	for name, at := range obj.AttributeTypes {
		if v, ok := set[name]; ok {
			vals[name] = v
		} else {
			vals[name] = tftypes.NewValue(at, nil)
		}
	}
	return tftypes.NewValue(obj, vals)
}

func str(s string) tftypes.Value { return tftypes.NewValue(tftypes.String, s) }

type resourceCase struct {
	name     string
	new      func() resource.Resource
	fixture  string
	create   map[string]tftypes.Value
	update   map[string]tftypes.Value
	importID string
}

func runResource(t *testing.T, tc resourceCase) {
	r := newRecorder(t, tc.fixture)
	ctx := r.ctx
	res := tc.new()
	res.(resource.ResourceWithConfigure).Configure(ctx, resource.ConfigureRequest{ProviderData: r.client}, &resource.ConfigureResponse{})
	var sr resource.SchemaResponse
	res.Schema(ctx, resource.SchemaRequest{}, &sr)
	s := sr.Schema
	typ := s.Type().TerraformType(ctx)
	empty := tfsdk.State{Schema: s, Raw: tftypes.NewValue(typ, nil)}

	r.step("create", r.item())
	cr := resource.CreateResponse{State: empty}
	res.Create(ctx, resource.CreateRequest{Plan: tfsdk.Plan{Schema: s, Raw: objectWith(ctx, typ, tc.create)}}, &cr)
	r.outcome(cr.Diagnostics, cr.State.Raw)
	state := cr.State

	r.step("read", r.item())
	rr := resource.ReadResponse{State: state}
	res.Read(ctx, resource.ReadRequest{State: state}, &rr)
	r.outcome(rr.Diagnostics, rr.State.Raw)

	r.step("update", r.item())
	ur := resource.UpdateResponse{State: state}
	res.Update(ctx, resource.UpdateRequest{Plan: tfsdk.Plan{Schema: s, Raw: objectWith(ctx, typ, tc.update)}, State: state}, &ur)
	r.outcome(ur.Diagnostics, ur.State.Raw)

	r.step("read not found", reply{http.StatusNotFound, `{"error":"not found"}`})
	nr := resource.ReadResponse{State: state}
	res.Read(ctx, resource.ReadRequest{State: state}, &nr)
	r.outcome(nr.Diagnostics, nr.State.Raw)

	r.step("delete", reply{http.StatusOK, `{}`})
	dr := resource.DeleteResponse{State: state}
	res.Delete(ctx, resource.DeleteRequest{State: state}, &dr)
	r.outcome(dr.Diagnostics, dr.State.Raw)

	r.step("import", r.item())
	ir := resource.ImportStateResponse{State: empty}
	res.(resource.ResourceWithImportState).ImportState(ctx, resource.ImportStateRequest{ID: cmp.Or(tc.importID, "id-1")}, &ir)
	r.outcome(ir.Diagnostics, ir.State.Raw)

	r.step("create rejected", reply{http.StatusBadRequest, `{"error":"invalid name"}`})
	er := resource.CreateResponse{State: empty}
	res.Create(ctx, resource.CreateRequest{Plan: tfsdk.Plan{Schema: s, Raw: objectWith(ctx, typ, tc.create)}}, &er)
	r.outcome(er.Diagnostics, er.State.Raw)

	r.compare(tc.name + "_resource")
}

type dataSourceCase struct {
	name     string
	single   func() datasource.DataSource
	list     func() datasource.DataSource
	fixture  string
	byID     map[string]tftypes.Value
	findBy   map[string]tftypes.Value
	lookups  []map[string]tftypes.Value
	listArgs map[string]tftypes.Value

	emptyFindBy   map[string]tftypes.Value
	emptyListArgs map[string]tftypes.Value
}

func readDataSource(r *recorder, ds datasource.DataSource, set map[string]tftypes.Value) {
	ctx := r.ctx
	ds.(datasource.DataSourceWithConfigure).Configure(ctx, datasource.ConfigureRequest{ProviderData: r.client}, &datasource.ConfigureResponse{})
	var sr datasource.SchemaResponse
	ds.Schema(ctx, datasource.SchemaRequest{}, &sr)
	typ := sr.Schema.Type().TerraformType(ctx)
	resp := datasource.ReadResponse{State: tfsdk.State{Schema: sr.Schema, Raw: tftypes.NewValue(typ, nil)}}
	ds.Read(ctx, datasource.ReadRequest{Config: tfsdk.Config{Schema: sr.Schema, Raw: objectWith(ctx, typ, set)}}, &resp)
	r.outcome(resp.Diagnostics, resp.State.Raw)
}

func findOneBy(ctx context.Context, ds datasource.DataSource, set map[string]tftypes.Value) map[string]tftypes.Value {
	var sr datasource.SchemaResponse
	ds.Schema(ctx, datasource.SchemaRequest{}, &sr)
	typ := sr.Schema.Type().TerraformType(ctx).(tftypes.Object).AttributeTypes["find_one_by"]
	return map[string]tftypes.Value{"find_one_by": objectWith(ctx, typ, set)}
}

func runDataSources(t *testing.T, tc dataSourceCase) {
	r := newRecorder(t, tc.fixture)

	r.step("data source by id", r.item())
	readDataSource(r, tc.single(), tc.byID)

	if tc.findBy != nil {
		r.step("data source find_one_by", r.page(""), r.item())
		readDataSource(r, tc.single(), findOneBy(r.ctx, tc.single(), tc.findBy))
	}

	for _, set := range tc.lookups {
		r.step(fmt.Sprintf("data source lookup %v", set), r.page(""))
		readDataSource(r, tc.single(), set)
	}

	if tc.list != nil {
		r.step("list data source, two pages", r.page("page-2"), r.page(""))
		readDataSource(r, tc.list(), tc.listArgs)

		r.step("list data source rejected", reply{http.StatusForbidden, `{"error":"forbidden"}`})
		readDataSource(r, tc.list(), tc.listArgs)
	}

	if tc.emptyFindBy != nil {
		r.step("data source find_one_by, empty filter", r.page(""), r.item())
		readDataSource(r, tc.single(), findOneBy(r.ctx, tc.single(), tc.emptyFindBy))
	}

	if tc.emptyListArgs != nil {
		r.step("list data source, empty filter", r.page(""))
		readDataSource(r, tc.list(), tc.emptyListArgs)
	}

	r.compare(tc.name + "_data_sources")
}

func (r *recorder) note(format string, args ...any) {
	r.mu.Lock()
	defer r.mu.Unlock()
	fmt.Fprintf(&r.log, "  %s\n", strings.ReplaceAll(fmt.Sprintf(format, args...), r.srv.URL, "SERVER"))
}
