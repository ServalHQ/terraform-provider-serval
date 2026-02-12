package hclcodec

import (
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestGenerateResourceBlock(t *testing.T) {
	type TestModel struct {
		ID    types.String `tfsdk:"id" json:"id,computed"`
		Name  types.String `tfsdk:"name" json:"name,required"`
		Email types.String `tfsdk:"email" json:"email,required"`
		Desc  types.String `tfsdk:"description" json:"description,optional"`
	}

	model := TestModel{
		ID:    types.StringValue("abc123"),
		Name:  types.StringValue("Alice"),
		Email: types.StringValue("alice@example.com"),
		Desc:  types.StringNull(),
	}

	hcl, err := GenerateResourceBlock("serval_user", "alice", model)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Should include resource wrapper
	if !strings.Contains(hcl, `resource "serval_user" "alice"`) {
		t.Errorf("missing resource block header in:\n%s", hcl)
	}

	// Should include required fields
	if !strings.Contains(hcl, `name = "Alice"`) {
		t.Errorf("missing required field name in:\n%s", hcl)
	}
	if !strings.Contains(hcl, `email = "alice@example.com"`) {
		t.Errorf("missing required field email in:\n%s", hcl)
	}

	// Should NOT include computed-only fields
	if strings.Contains(hcl, "id =") {
		t.Errorf("should not include computed field id in:\n%s", hcl)
	}

	// Should NOT include null optional fields
	if strings.Contains(hcl, "description") {
		t.Errorf("should not include null optional field description in:\n%s", hcl)
	}
}

func TestIncludesOptionalWhenSet(t *testing.T) {
	type TestModel struct {
		ID   types.String `tfsdk:"id" json:"id,computed"`
		Name types.String `tfsdk:"name" json:"name,required"`
		Desc types.String `tfsdk:"description" json:"description,optional"`
	}

	model := TestModel{
		ID:   types.StringValue("abc"),
		Name: types.StringValue("Test"),
		Desc: types.StringValue("A description"),
	}

	hcl, err := GenerateResourceBlock("serval_user", "test", model)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(hcl, `description = "A description"`) {
		t.Errorf("should include non-null optional field in:\n%s", hcl)
	}
}

func TestIncludesComputedOptionalWhenSet(t *testing.T) {
	type TestModel struct {
		ID       types.String `tfsdk:"id" json:"id,computed"`
		Name     types.String `tfsdk:"name" json:"name,required"`
		Required types.Bool   `tfsdk:"require_justification" json:"requireJustification,computed_optional"`
	}

	model := TestModel{
		ID:       types.StringValue("abc"),
		Name:     types.StringValue("Test"),
		Required: types.BoolValue(true),
	}

	hcl, err := GenerateResourceBlock("serval_policy", "test", model)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(hcl, `require_justification = true`) {
		t.Errorf("should include non-null computed_optional field in:\n%s", hcl)
	}
}

func TestExcludesNullComputedOptional(t *testing.T) {
	type TestModel struct {
		ID       types.String `tfsdk:"id" json:"id,computed"`
		Name     types.String `tfsdk:"name" json:"name,required"`
		Required types.Bool   `tfsdk:"require_justification" json:"requireJustification,computed_optional"`
	}

	model := TestModel{
		ID:       types.StringValue("abc"),
		Name:     types.StringValue("Test"),
		Required: types.BoolNull(),
	}

	hcl, err := GenerateResourceBlock("serval_policy", "test", model)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if strings.Contains(hcl, "require_justification") {
		t.Errorf("should not include null computed_optional field in:\n%s", hcl)
	}
}

func TestIntField(t *testing.T) {
	type TestModel struct {
		ID      types.String `tfsdk:"id" json:"id,computed"`
		MaxMins types.Int64  `tfsdk:"max_minutes" json:"maxMinutes,optional"`
	}

	model := TestModel{
		ID:      types.StringValue("abc"),
		MaxMins: types.Int64Value(60),
	}

	hcl, err := GenerateResourceBlock("serval_policy", "test", model)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(hcl, "max_minutes = 60") {
		t.Errorf("should include int field in:\n%s", hcl)
	}
}

func TestPointerSliceOfStrings(t *testing.T) {
	type TestModel struct {
		ID      types.String    `tfsdk:"id" json:"id,computed"`
		UserIDs *[]types.String `tfsdk:"user_ids" json:"userIds,optional"`
	}

	userIDs := []types.String{
		types.StringValue("u1"),
		types.StringValue("u2"),
	}

	model := TestModel{
		ID:      types.StringValue("abc"),
		UserIDs: &userIDs,
	}

	hcl, err := GenerateResourceBlock("serval_group", "test", model)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(hcl, `user_ids = ["u1", "u2"]`) {
		t.Errorf("should include string list in:\n%s", hcl)
	}
}

func TestNilPointerSlice(t *testing.T) {
	type TestModel struct {
		ID      types.String    `tfsdk:"id" json:"id,computed"`
		UserIDs *[]types.String `tfsdk:"user_ids" json:"userIds,optional"`
	}

	model := TestModel{
		ID:      types.StringValue("abc"),
		UserIDs: nil,
	}

	hcl, err := GenerateResourceBlock("serval_group", "test", model)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if strings.Contains(hcl, "user_ids") {
		t.Errorf("should not include nil pointer slice in:\n%s", hcl)
	}
}

func TestNestedStruct(t *testing.T) {
	type Inner struct {
		WorkflowID types.String `tfsdk:"workflow_id" json:"workflowId,optional"`
	}
	type TestModel struct {
		ID    types.String `tfsdk:"id" json:"id,computed"`
		Inner *Inner       `tfsdk:"custom_workflow" json:"customWorkflow,optional"`
	}

	model := TestModel{
		ID: types.StringValue("abc"),
		Inner: &Inner{
			WorkflowID: types.StringValue("wf-1"),
		},
	}

	hcl, err := GenerateResourceBlock("serval_role", "test", model)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(hcl, "custom_workflow {") {
		t.Errorf("should include nested block in:\n%s", hcl)
	}
	if !strings.Contains(hcl, `workflow_id = "wf-1"`) {
		t.Errorf("should include nested field in:\n%s", hcl)
	}
}

func TestNilNestedStruct(t *testing.T) {
	type Inner struct {
		WorkflowID types.String `tfsdk:"workflow_id" json:"workflowId,optional"`
	}
	type TestModel struct {
		ID    types.String `tfsdk:"id" json:"id,computed"`
		Inner *Inner       `tfsdk:"custom_workflow" json:"customWorkflow,optional"`
	}

	model := TestModel{
		ID:    types.StringValue("abc"),
		Inner: nil,
	}

	hcl, err := GenerateResourceBlock("serval_role", "test", model)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if strings.Contains(hcl, "custom_workflow") {
		t.Errorf("should not include nil nested block in:\n%s", hcl)
	}
}

func TestPathTagRequired(t *testing.T) {
	// Some fields use path tag instead of json tag for the option
	type TestModel struct {
		ID         types.String `tfsdk:"id" json:"id,computed"`
		WorkflowID types.String `tfsdk:"workflow_id" json:"workflowId" path:"workflow_id,required"`
	}

	model := TestModel{
		ID:         types.StringValue("abc"),
		WorkflowID: types.StringValue("wf-1"),
	}

	hcl, err := GenerateResourceBlock("serval_approval", "test", model)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(hcl, `workflow_id = "wf-1"`) {
		t.Errorf("should include path-tagged required field in:\n%s", hcl)
	}
}

func TestGenerateAttributes(t *testing.T) {
	type TestModel struct {
		ID   types.String `tfsdk:"id" json:"id,computed"`
		Name types.String `tfsdk:"name" json:"name,required"`
	}

	model := TestModel{
		ID:   types.StringValue("abc"),
		Name: types.StringValue("Test"),
	}

	attrs, err := GenerateAttributes(model)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Should NOT have resource wrapper
	if strings.Contains(attrs, "resource") {
		t.Errorf("should not include resource wrapper in attributes:\n%s", attrs)
	}

	// Should have the attribute
	if !strings.Contains(attrs, `name = "Test"`) {
		t.Errorf("should include attribute:\n%s", attrs)
	}
}
