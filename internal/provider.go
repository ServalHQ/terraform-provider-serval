// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package internal

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/serval-go/option"
	"github.com/ServalHQ/terraform-provider-serval/internal/cache"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/access_policy"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/access_policy_approval_procedure"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/access_request"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/app_instance"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/app_resource"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/app_resource_role"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/custom_service"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/group"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/guidance"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/team"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/team_user"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/user"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/workflow"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/workflow_approval_procedure"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/workflow_run"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ provider.ProviderWithConfigValidators = (*ServalProvider)(nil)

// ServalProvider defines the provider implementation.
type ServalProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// ServalProviderModel describes the provider data model.
type ServalProviderModel struct {
	BaseURL           types.String `tfsdk:"base_url" json:"base_url,optional"`
	ClientID          types.String `tfsdk:"client_id" json:"client_id,optional"`
	ClientSecret      types.String `tfsdk:"client_secret" json:"client_secret,optional"`
	BearerToken       types.String `tfsdk:"bearer_token" json:"bearer_token,optional"`
	PrefetchForTeams  types.List   `tfsdk:"prefetch_for_teams" json:"prefetch_for_teams,optional"`
	PrefetchCachePath types.String `tfsdk:"prefetch_cache_path" json:"prefetch_cache_path,optional"`
}

func (p *ServalProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "serval"
	resp.Version = p.version
}

func ProviderSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				Description: "Set the base url that the provider connects to.",
				Optional:    true,
			},
			"client_id": schema.StringAttribute{
				Optional: true,
			},
			"client_secret": schema.StringAttribute{
				Optional: true,
			},
			"bearer_token": schema.StringAttribute{
				Optional: true,
			},
			"prefetch_for_teams": schema.ListAttribute{
				Description: "List of team IDs to pre-fetch resources for. " +
					"When set, all resources for these teams are loaded into memory during initialization. " +
					"Optimized for bulk import operations.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"prefetch_cache_path": schema.StringAttribute{
				Description: "Path to a directory containing pre-built cache files. " +
					"When set, loads cache from files instead of making API calls. " +
					"For internal use by Serval tooling.",
				Optional: true,
			},
		},
	}
}

func (p *ServalProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = ProviderSchema(ctx)
}

func (p *ServalProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	var data ServalProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	opts := []option.RequestOption{}

	if !data.BaseURL.IsNull() && !data.BaseURL.IsUnknown() {
		opts = append(opts, option.WithBaseURL(data.BaseURL.ValueString()))
	} else if o, ok := os.LookupEnv("SERVAL_BASE_URL"); ok {
		opts = append(opts, option.WithBaseURL(o))
	}

	if !data.ClientID.IsNull() && !data.ClientID.IsUnknown() {
		opts = append(opts, option.WithClientID(data.ClientID.ValueString()))
	} else if o, ok := os.LookupEnv("SERVAL_CLIENT_ID"); ok {
		opts = append(opts, option.WithClientID(o))
	}

	if !data.ClientSecret.IsNull() && !data.ClientSecret.IsUnknown() {
		opts = append(opts, option.WithClientSecret(data.ClientSecret.ValueString()))
	} else if o, ok := os.LookupEnv("SERVAL_CLIENT_SECRET"); ok {
		opts = append(opts, option.WithClientSecret(o))
	}

	if !data.BearerToken.IsNull() && !data.BearerToken.IsUnknown() {
		opts = append(opts, option.WithBearerToken(data.BearerToken.ValueString()))
	} else if o, ok := os.LookupEnv("SERVAL_BEARER_TOKEN"); ok {
		opts = append(opts, option.WithBearerToken(o))
	}

	client := serval.NewClient(
		opts...,
	)

	// Extract team IDs from prefetch_for_teams list attribute
	var teamIDs []string
	if !data.PrefetchForTeams.IsNull() && !data.PrefetchForTeams.IsUnknown() {
		var ids []types.String
		resp.Diagnostics.Append(data.PrefetchForTeams.ElementsAs(ctx, &ids, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
		for _, id := range ids {
			teamIDs = append(teamIDs, id.ValueString())
		}
	}

	prefetchFromAPI := len(teamIDs) > 0
	prefetchFromFile := !data.PrefetchCachePath.IsNull() && data.PrefetchCachePath.ValueString() != ""

	if prefetchFromAPI || prefetchFromFile {
		cache.PrefetchMode = true
	}

	if prefetchFromFile {
		// FUTURE: Load caches from files at data.PrefetchCachePath.ValueString()
		// Each service will have a LoadFromFile function that reads its cache.
		// When implemented, the svmeta Fetcher will write these files using
		// serval-go SDK types, and the provider will load them here.
		prefetchFromAPI = true // Fallback until file loading is implemented
	}

	if prefetchFromAPI {
		totalStart := time.Now()

		type resourceStats struct {
			name       string
			durationMs int64
			apiCalls   int
			items      int
		}

		var mu sync.Mutex
		var allErrors []string
		var stats []resourceStats

		recordStats := func(name string, durationMs int64, apiCalls, items int) {
			mu.Lock()
			stats = append(stats, resourceStats{name: name, durationMs: durationMs, apiCalls: apiCalls, items: items})
			mu.Unlock()
		}
		recordError := func(name string, err error) {
			mu.Lock()
			allErrors = append(allErrors, fmt.Sprintf("prefetch failed: %s: %s", name, err.Error()))
			mu.Unlock()
		}

		type prefetchJob struct {
			name string
			fn   func() (int, error)
			len  func() int
		}

		// All resources prefetched in a single parallel phase (team-scoped endpoints)
		jobs := []prefetchJob{
			{"users", func() (int, error) { return user.Prefetch(ctx, &client) }, func() int { return user.Cache.Len() }},
			{"groups", func() (int, error) { return group.Prefetch(ctx, &client) }, func() int { return group.Cache.Len() }},
			{"teams", func() (int, error) { return team.Prefetch(ctx, &client) }, func() int { return team.Cache.Len() }},
			{"workflows", func() (int, error) { return workflow.Prefetch(ctx, &client, teamIDs) }, func() int { return workflow.Cache.Len() }},
			{"guidances", func() (int, error) { return guidance.Prefetch(ctx, &client, teamIDs) }, func() int { return guidance.Cache.Len() }},
			{"access_policies", func() (int, error) { return access_policy.Prefetch(ctx, &client, teamIDs) }, func() int { return access_policy.Cache.Len() }},
			{"app_instances", func() (int, error) { return app_instance.Prefetch(ctx, &client, teamIDs) }, func() int { return app_instance.Cache.Len() }},
			{"custom_services", func() (int, error) { return custom_service.Prefetch(ctx, &client, teamIDs) }, func() int { return custom_service.Cache.Len() }},
			{"app_resources", func() (int, error) { return app_resource.Prefetch(ctx, &client, teamIDs) }, func() int { return app_resource.Cache.Len() }},
			{"app_resource_roles", func() (int, error) { return app_resource_role.Prefetch(ctx, &client, teamIDs) }, func() int { return app_resource_role.Cache.Len() }},
			{"workflow_approval_procedures", func() (int, error) {
				return workflow_approval_procedure.Prefetch(ctx, &client, teamIDs)
			}, func() int { return workflow_approval_procedure.Cache.Len() }},
			{"access_policy_approval_procedures", func() (int, error) {
				return access_policy_approval_procedure.Prefetch(ctx, &client, teamIDs)
			}, func() int { return access_policy_approval_procedure.Cache.Len() }},
		}

		var wg sync.WaitGroup
		for _, pf := range jobs {
			wg.Add(1)
			go func(j prefetchJob) {
				defer wg.Done()
				start := time.Now()
				apiCalls, err := j.fn()
				dur := time.Since(start).Milliseconds()
				if err != nil {
					recordError(j.name, err)
					return
				}
				recordStats(j.name, dur, apiCalls, j.len())
			}(pf)
		}
		wg.Wait()

		if len(allErrors) > 0 {
			for _, e := range allErrors {
				resp.Diagnostics.AddError("prefetch failed", e)
			}
			return
		}

		// Build and log the comprehensive prefetch report
		totalDuration := time.Since(totalStart).Milliseconds()
		totalAPICalls := 0
		totalItems := 0
		report := map[string]interface{}{
			"total_duration_ms": totalDuration,
			"team_ids":          teamIDs,
		}
		for _, s := range stats {
			totalAPICalls += s.apiCalls
			totalItems += s.items
			report[s.name+"_duration_ms"] = s.durationMs
			report[s.name+"_api_calls"] = s.apiCalls
			report[s.name+"_items"] = s.items
		}
		report["total_api_calls"] = totalAPICalls
		report["total_items"] = totalItems

		tflog.Info(ctx, "prefetch: complete", report)
	}

	resp.DataSourceData = &client
	resp.ResourceData = &client
}

func (p *ServalProvider) ConfigValidators(_ context.Context) []provider.ConfigValidator {
	return []provider.ConfigValidator{}
}

func (p *ServalProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		access_policy.NewResource,
		access_policy_approval_procedure.NewResource,
		guidance.NewResource,
		workflow.NewResource,
		workflow_approval_procedure.NewResource,
		app_instance.NewResource,
		app_resource.NewResource,
		app_resource_role.NewResource,
		user.NewResource,
		group.NewResource,
		team.NewResource,
		team_user.NewResource,
		custom_service.NewResource,
	}
}

func (p *ServalProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		access_policy.NewAccessPolicyDataSource,
		access_policy_approval_procedure.NewAccessPolicyApprovalProcedureDataSource,
		guidance.NewGuidanceDataSource,
		workflow.NewWorkflowDataSource,
		workflow_approval_procedure.NewWorkflowApprovalProcedureDataSource,
		workflow_run.NewWorkflowRunDataSource,
		access_request.NewAccessRequestDataSource,
		app_instance.NewAppInstanceDataSource,
		app_resource.NewAppResourceDataSource,
		app_resource_role.NewAppResourceRoleDataSource,
		user.NewUserDataSource,
		group.NewGroupDataSource,
		team.NewTeamDataSource,
		team_user.NewTeamUserDataSource,
		custom_service.NewCustomServiceDataSource,
	}
}

func NewProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &ServalProvider{
			version: version,
		}
	}
}
