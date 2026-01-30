// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package internal

import (
	"context"
	"os"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/serval-go/option"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/access_policy"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/access_policy_approval_procedure"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/app_instance"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/app_resource"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/app_resource_role"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/custom_service"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/group"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/guidance"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/team"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/user"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/workflow"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/workflow_approval_procedure"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
	BaseURL      types.String `tfsdk:"base_url" json:"base_url,optional"`
	ClientID     types.String `tfsdk:"client_id" json:"client_id,optional"`
	ClientSecret types.String `tfsdk:"client_secret" json:"client_secret,optional"`
	BearerToken  types.String `tfsdk:"bearer_token" json:"bearer_token,optional"`
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

	// Initialize resource caches for improved performance with large datasets
	user.InitCache()
	group.InitCache()
	app_resource.InitCache()
	app_resource_role.InitCache()

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
		app_instance.NewAppInstanceDataSource,
		app_resource.NewAppResourceDataSource,
		app_resource_role.NewAppResourceRoleDataSource,
		user.NewUserDataSource,
		group.NewGroupDataSource,
		team.NewTeamDataSource,
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
