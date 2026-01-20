// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflow_approval_procedure

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/serval-go/option"
	"github.com/ServalHQ/terraform-provider-serval/internal/apijson"
	"github.com/ServalHQ/terraform-provider-serval/internal/importpath"
	"github.com/ServalHQ/terraform-provider-serval/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*WorkflowApprovalProcedureResource)(nil)
var _ resource.ResourceWithModifyPlan = (*WorkflowApprovalProcedureResource)(nil)
var _ resource.ResourceWithImportState = (*WorkflowApprovalProcedureResource)(nil)

func NewResource() resource.Resource {
	return &WorkflowApprovalProcedureResource{}
}

// WorkflowApprovalProcedureResource defines the resource implementation.
type WorkflowApprovalProcedureResource struct {
	client *serval.Client
}

func (r *WorkflowApprovalProcedureResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_workflow_approval_procedure"
}

func (r *WorkflowApprovalProcedureResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*serval.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *serval.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *WorkflowApprovalProcedureResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *WorkflowApprovalProcedureModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := data.MarshalJSON()
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := WorkflowApprovalProcedureDataEnvelope{*data}
	_, err = r.client.Workflows.ApprovalProcedures.New(
		ctx,
		data.WorkflowID.ValueString(),
		serval.WorkflowApprovalProcedureNewParams{},
		option.WithRequestBody("application/json", dataBytes),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.UnmarshalComputed(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Data

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkflowApprovalProcedureResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *WorkflowApprovalProcedureModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *WorkflowApprovalProcedureModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := data.MarshalJSONForUpdate(*state)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := WorkflowApprovalProcedureDataEnvelope{*data}
	_, err = r.client.Workflows.ApprovalProcedures.Update(
		ctx,
		data.ID.ValueString(),
		serval.WorkflowApprovalProcedureUpdateParams{
			WorkflowID: data.WorkflowID.ValueString(),
		},
		option.WithRequestBody("application/json", dataBytes),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.UnmarshalComputed(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Data

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkflowApprovalProcedureResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *WorkflowApprovalProcedureModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := WorkflowApprovalProcedureDataEnvelope{*data}
	_, err := r.client.Workflows.ApprovalProcedures.Get(
		ctx,
		data.ID.ValueString(),
		serval.WorkflowApprovalProcedureGetParams{
			WorkflowID: data.WorkflowID.ValueString(),
		},
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if res != nil && res.StatusCode == 404 {
		resp.Diagnostics.AddWarning("Resource not found", "The resource was not found on the server and will be removed from state.")
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Data

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkflowApprovalProcedureResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *WorkflowApprovalProcedureModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.Workflows.ApprovalProcedures.Delete(
		ctx,
		data.ID.ValueString(),
		serval.WorkflowApprovalProcedureDeleteParams{
			WorkflowID: data.WorkflowID.ValueString(),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkflowApprovalProcedureResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data = new(WorkflowApprovalProcedureModel)

	path_workflow_id := ""
	path_id := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<workflow_id>/<id>",
		&path_workflow_id,
		&path_id,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.WorkflowID = types.StringValue(path_workflow_id)
	data.ID = types.StringValue(path_id)

	res := new(http.Response)
	env := WorkflowApprovalProcedureDataEnvelope{*data}
	_, err := r.client.Workflows.ApprovalProcedures.Get(
		ctx,
		path_id,
		serval.WorkflowApprovalProcedureGetParams{
			WorkflowID: path_workflow_id,
		},
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Data

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkflowApprovalProcedureResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
