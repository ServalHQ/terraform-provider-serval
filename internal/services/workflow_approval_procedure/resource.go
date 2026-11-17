// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflow_approval_procedure

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/ServalHQ/terraform-provider-serval/internal/apijson"
	"github.com/ServalHQ/terraform-provider-serval/internal/client"
	"github.com/ServalHQ/terraform-provider-serval/internal/importpath"
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
	client *client.Client
}

func (r *WorkflowApprovalProcedureResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_workflow_approval_procedure"
}

func (r *WorkflowApprovalProcedureResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	c, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = c
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
	env := WorkflowApprovalProcedureDataEnvelope{*data}
	res, err := doRequest(ctx, r.client, http.MethodPost, data.WorkflowID.ValueString(), "", dataBytes)
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
	env := WorkflowApprovalProcedureDataEnvelope{*data}
	res, err := doRequest(ctx, r.client, http.MethodPut, data.WorkflowID.ValueString(), data.ID.ValueString(), dataBytes)
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

	if item, ok, err := TryRead(data.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("prefetch cache miss", err.Error())
		return
	} else if ok {
		resp.Diagnostics.Append(resp.State.Set(ctx, item)...)
		return
	}

	env := WorkflowApprovalProcedureDataEnvelope{*data}
	res, err := doRequest(ctx, r.client, http.MethodGet, data.WorkflowID.ValueString(), data.ID.ValueString(), nil)
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

	res, err := doRequest(ctx, r.client, http.MethodDelete, data.WorkflowID.ValueString(), data.ID.ValueString(), nil)
	if res != nil {
		res.Body.Close()
	}
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

	if item, ok, err := TryRead(path_id); err != nil {
		resp.Diagnostics.AddError("prefetch cache miss", err.Error())
		return
	} else if ok {
		resp.Diagnostics.Append(resp.State.Set(ctx, item)...)
		return
	}

	env := WorkflowApprovalProcedureDataEnvelope{*data}
	res, err := doRequest(ctx, r.client, http.MethodGet, path_workflow_id, path_id, nil)
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
