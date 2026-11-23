// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy_approval_procedure

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
var _ resource.ResourceWithConfigure = (*AccessPolicyApprovalProcedureResource)(nil)
var _ resource.ResourceWithModifyPlan = (*AccessPolicyApprovalProcedureResource)(nil)
var _ resource.ResourceWithImportState = (*AccessPolicyApprovalProcedureResource)(nil)

func NewResource() resource.Resource {
	return &AccessPolicyApprovalProcedureResource{}
}

// AccessPolicyApprovalProcedureResource defines the resource implementation.
type AccessPolicyApprovalProcedureResource struct {
	client *client.Client
}

func (r *AccessPolicyApprovalProcedureResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_access_policy_approval_procedure"
}

func (r *AccessPolicyApprovalProcedureResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *AccessPolicyApprovalProcedureResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *AccessPolicyApprovalProcedureModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := data.MarshalJSON()
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	env := AccessPolicyApprovalProcedureDataEnvelope{*data}
	res, err := doRequest(ctx, r.client, http.MethodPost, data.AccessPolicyID.ValueString(), "", dataBytes)
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

func (r *AccessPolicyApprovalProcedureResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *AccessPolicyApprovalProcedureModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *AccessPolicyApprovalProcedureModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := data.MarshalJSONForUpdate(*state)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	env := AccessPolicyApprovalProcedureDataEnvelope{*data}
	res, err := doRequest(ctx, r.client, http.MethodPut, data.AccessPolicyID.ValueString(), data.ID.ValueString(), dataBytes)
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

func (r *AccessPolicyApprovalProcedureResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *AccessPolicyApprovalProcedureModel

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

	env := AccessPolicyApprovalProcedureDataEnvelope{*data}
	res, err := doRequest(ctx, r.client, http.MethodGet, data.AccessPolicyID.ValueString(), data.ID.ValueString(), nil)
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

func (r *AccessPolicyApprovalProcedureResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *AccessPolicyApprovalProcedureModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res, err := doRequest(ctx, r.client, http.MethodDelete, data.AccessPolicyID.ValueString(), data.ID.ValueString(), nil)
	if err == nil {
		err = client.CheckJSONBody(res)
	}
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AccessPolicyApprovalProcedureResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data = new(AccessPolicyApprovalProcedureModel)

	path_access_policy_id := ""
	path_id := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<access_policy_id>/<id>",
		&path_access_policy_id,
		&path_id,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.AccessPolicyID = types.StringValue(path_access_policy_id)
	data.ID = types.StringValue(path_id)

	if item, ok, err := TryRead(path_id); err != nil {
		resp.Diagnostics.AddError("prefetch cache miss", err.Error())
		return
	} else if ok {
		resp.Diagnostics.Append(resp.State.Set(ctx, item)...)
		return
	}

	env := AccessPolicyApprovalProcedureDataEnvelope{*data}
	res, err := doRequest(ctx, r.client, http.MethodGet, path_access_policy_id, path_id, nil)
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

func (r *AccessPolicyApprovalProcedureResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
