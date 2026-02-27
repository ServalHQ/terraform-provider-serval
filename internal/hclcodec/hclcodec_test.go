package hclcodec

import (
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
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

	hcl, err := GenerateResourceBlock("serval_user", "alice", model, nil)
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

	hcl, err := GenerateResourceBlock("serval_user", "test", model, nil)
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

	hcl, err := GenerateResourceBlock("serval_policy", "test", model, nil)
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

	hcl, err := GenerateResourceBlock("serval_policy", "test", model, nil)
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

	hcl, err := GenerateResourceBlock("serval_policy", "test", model, nil)
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

	hcl, err := GenerateResourceBlock("serval_group", "test", model, nil)
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

	hcl, err := GenerateResourceBlock("serval_group", "test", model, nil)
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

	hcl, err := GenerateResourceBlock("serval_role", "test", model, nil)
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

	hcl, err := GenerateResourceBlock("serval_role", "test", model, nil)
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

	hcl, err := GenerateResourceBlock("serval_approval", "test", model, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(hcl, `workflow_id = "wf-1"`) {
		t.Errorf("should include path-tagged required field in:\n%s", hcl)
	}
}

func TestNestedStructAttributeMode(t *testing.T) {
	type Inner struct {
		WorkflowID types.String `tfsdk:"workflow_id" json:"workflowId,optional"`
	}
	type Outer struct {
		InnerA *Inner `tfsdk:"inner_a" json:"innerA,optional"`
		InnerB *Inner `tfsdk:"inner_b" json:"innerB,optional"`
	}
	type TestModel struct {
		ID    types.String `tfsdk:"id" json:"id,computed"`
		Outer *Outer       `tfsdk:"method" json:"method,required"`
	}

	model := TestModel{
		ID: types.StringValue("abc"),
		Outer: &Outer{
			InnerA: &Inner{WorkflowID: types.StringValue("wf-1")},
			InnerB: nil,
		},
	}

	// With schema info marking method and its children as attribute mode
	schema := SchemaInfo{
		"method": FieldSchema{
			NestedMode: NestedModeAttr,
			Children: SchemaInfo{
				"inner_a": {NestedMode: NestedModeAttr},
				"inner_b": {NestedMode: NestedModeAttr},
			},
		},
	}

	hcl, err := GenerateResourceBlock("serval_role", "test", model, schema)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Top-level nested attribute should use "= {"
	if !strings.Contains(hcl, "method = {") {
		t.Errorf("should use attribute syntax for method in:\n%s", hcl)
	}

	// Child nested attribute should also use "= {"
	if !strings.Contains(hcl, "inner_a = {") {
		t.Errorf("should use attribute syntax for inner_a in:\n%s", hcl)
	}

	// inner_b is nil, should not appear
	if strings.Contains(hcl, "inner_b") {
		t.Errorf("should not include nil nested struct inner_b in:\n%s", hcl)
	}

	// The inner field should still be present
	if !strings.Contains(hcl, `workflow_id = "wf-1"`) {
		t.Errorf("should include nested field in:\n%s", hcl)
	}
}

func TestNestedStructBlockModeDefault(t *testing.T) {
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

	// Without schema info → default block mode
	hcl, err := GenerateResourceBlock("serval_role", "test", model, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Should use block syntax (no "=")
	if !strings.Contains(hcl, "custom_workflow {") {
		t.Errorf("should use block syntax for custom_workflow in:\n%s", hcl)
	}
	if strings.Contains(hcl, "custom_workflow = {") {
		t.Errorf("should NOT use attribute syntax without schema info in:\n%s", hcl)
	}
}

func TestMapEnumValue(t *testing.T) {
	allowed := []string{
		"USER_ROLE_UNSPECIFIED",
		"USER_ROLE_ORG_MEMBER",
		"USER_ROLE_ORG_ADMIN",
		"USER_ROLE_ORG_GUEST",
	}

	tests := []struct {
		input    string
		expected string
	}{
		// Exact match (case-insensitive)
		{"USER_ROLE_ORG_ADMIN", "USER_ROLE_ORG_ADMIN"},
		{"user_role_org_admin", "USER_ROLE_ORG_ADMIN"},

		// Suffix match
		{"admin", "USER_ROLE_ORG_ADMIN"},
		{"ADMIN", "USER_ROLE_ORG_ADMIN"},
		{"member", "USER_ROLE_ORG_MEMBER"},
		{"guest", "USER_ROLE_ORG_GUEST"},

		// No match — pass through unchanged
		{"superadmin", "superadmin"},
		{"unknown", "unknown"},
	}

	for _, tt := range tests {
		got := mapEnumValue(tt.input, allowed)
		if got != tt.expected {
			t.Errorf("mapEnumValue(%q) = %q, want %q", tt.input, got, tt.expected)
		}
	}
}

func TestMapEnumValueAmbiguous(t *testing.T) {
	// "PRIVATE" is a suffix of both values — ambiguous, should pass through
	allowed := []string{"TEAM_PRIVATE", "SCOPE_PRIVATE"}
	got := mapEnumValue("private", allowed)
	if got != "private" {
		t.Errorf("ambiguous suffix should pass through, got %q", got)
	}
}

func TestMapEnumValueNilAllowed(t *testing.T) {
	got := mapEnumValue("anything", nil)
	if got != "anything" {
		t.Errorf("nil allowed should pass through, got %q", got)
	}
}

func TestHclQuote(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Plain string — no escaping needed
		{"hello", `"hello"`},
		// Dollar-brace interpolation
		{"Hello ${name}", `"Hello $${name}"`},
		// Percent-brace template directive
		{"Hello %{if true}yes%{endif}", `"Hello %%{if true}yes%%{endif}"`},
		// Both sequences
		{"${a} and %{b}", `"$${a} and %%{b}"`},
		// Already-escaped Go characters should be preserved
		{"line1\nline2", `"line1\nline2"`},
		// No braces after $ — should not be escaped
		{"price is $5", `"price is $5"`},
	}

	for _, tt := range tests {
		got := hclQuote(tt.input)
		if got != tt.expected {
			t.Errorf("hclQuote(%q) = %s, want %s", tt.input, got, tt.expected)
		}
	}
}

func TestStringWithInterpolation(t *testing.T) {
	type TestModel struct {
		ID   types.String `tfsdk:"id" json:"id,computed"`
		Body types.String `tfsdk:"body" json:"body,required"`
	}

	model := TestModel{
		ID:   types.StringValue("abc"),
		Body: types.StringValue("Hello ${userEmail}, welcome to ${teamName}"),
	}

	hcl, err := GenerateResourceBlock("serval_workflow", "test", model, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// The output should contain escaped interpolation sequences
	if !strings.Contains(hcl, `$${userEmail}`) {
		t.Errorf("should escape ${ as $${ in:\n%s", hcl)
	}
	if !strings.Contains(hcl, `$${teamName}`) {
		t.Errorf("should escape ${ as $${ in:\n%s", hcl)
	}
	// Should NOT contain unescaped ${
	if strings.Contains(hcl, `${userEmail}`) && !strings.Contains(hcl, `$${userEmail}`) {
		t.Errorf("should not contain unescaped interpolation in:\n%s", hcl)
	}
}

func TestListOfObjectsAttrSkipsNullElements(t *testing.T) {
	// Regression test: when null ObjectValue elements precede valid ones in a
	// list, the comma separator must only be written between actually-emitted
	// objects. Previously the code used the loop index (i > 0) which produced
	// a leading comma like [, {...}] when the first element was null.
	nullObj := types.ObjectNull(map[string]attr.Type{
		"name": types.StringType,
	})
	validObj := types.ObjectValueMust(map[string]attr.Type{
		"name": types.StringType,
	}, map[string]attr.Value{
		"name": types.StringValue("hello"),
	})

	elements := []attr.Value{nullObj, validObj}
	hcl, err := serializeListOfObjectsAsAttrToHCL(
		"items", elements, "  ", 2, nil,
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Should NOT start with a comma after the opening bracket
	if strings.Contains(hcl, "[,") || strings.Contains(hcl, "[ ,") {
		t.Errorf("should not have leading comma in list:\n%s", hcl)
	}

	// Should contain the valid object
	if !strings.Contains(hcl, `name = "hello"`) {
		t.Errorf("should include valid object attributes in:\n%s", hcl)
	}

	// Verify the output parses as valid HCL by checking basic structure
	if !strings.Contains(hcl, "items = [{") {
		t.Errorf("should have proper list-of-objects syntax in:\n%s", hcl)
	}
}

func TestListOfObjectsAttrMultipleWithSkips(t *testing.T) {
	// Verify commas are correct when null elements appear between valid ones.
	objType := map[string]attr.Type{"id": types.StringType}
	null := types.ObjectNull(objType)
	obj1 := types.ObjectValueMust(objType, map[string]attr.Value{
		"id": types.StringValue("a"),
	})
	obj2 := types.ObjectValueMust(objType, map[string]attr.Value{
		"id": types.StringValue("b"),
	})

	elements := []attr.Value{null, obj1, null, obj2}
	hcl, err := serializeListOfObjectsAsAttrToHCL(
		"items", elements, "  ", 2, nil,
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Should have exactly one comma (between the two valid objects)
	commaCount := strings.Count(hcl, ", {")
	if commaCount != 1 {
		t.Errorf("expected 1 comma between objects, got %d in:\n%s", commaCount, hcl)
	}

	if !strings.Contains(hcl, `id = "a"`) || !strings.Contains(hcl, `id = "b"`) {
		t.Errorf("should include both valid objects in:\n%s", hcl)
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

	attrs, err := GenerateAttributes(model, nil)
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
