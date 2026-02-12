package statecodec

import (
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestSerializeSimpleStruct(t *testing.T) {
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

	raw, err := SerializeToStateAttributes(model)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var attrs map[string]interface{}
	if err := json.Unmarshal(raw, &attrs); err != nil {
		t.Fatalf("failed to unmarshal result: %v", err)
	}

	if attrs["id"] != "abc123" {
		t.Errorf("expected id=abc123, got %v", attrs["id"])
	}
	if attrs["name"] != "Alice" {
		t.Errorf("expected name=Alice, got %v", attrs["name"])
	}
	if attrs["email"] != "alice@example.com" {
		t.Errorf("expected email=alice@example.com, got %v", attrs["email"])
	}
	if attrs["description"] != nil {
		t.Errorf("expected description=null, got %v", attrs["description"])
	}
}

func TestSerializeWithIntAndBool(t *testing.T) {
	type TestModel struct {
		ID       types.String `tfsdk:"id"`
		Name     types.String `tfsdk:"name"`
		MaxMins  types.Int64  `tfsdk:"max_minutes"`
		Required types.Bool   `tfsdk:"required"`
	}

	model := TestModel{
		ID:       types.StringValue("xyz"),
		Name:     types.StringValue("Test"),
		MaxMins:  types.Int64Value(60),
		Required: types.BoolValue(true),
	}

	raw, err := SerializeToStateAttributes(model)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var attrs map[string]interface{}
	if err := json.Unmarshal(raw, &attrs); err != nil {
		t.Fatalf("failed to unmarshal result: %v", err)
	}

	if attrs["max_minutes"] != float64(60) {
		t.Errorf("expected max_minutes=60, got %v", attrs["max_minutes"])
	}
	if attrs["required"] != true {
		t.Errorf("expected required=true, got %v", attrs["required"])
	}
}

func TestSerializeNullInt(t *testing.T) {
	type TestModel struct {
		ID      types.String `tfsdk:"id"`
		MaxMins types.Int64  `tfsdk:"max_minutes"`
	}

	model := TestModel{
		ID:      types.StringValue("xyz"),
		MaxMins: types.Int64Null(),
	}

	raw, err := SerializeToStateAttributes(model)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var attrs map[string]interface{}
	if err := json.Unmarshal(raw, &attrs); err != nil {
		t.Fatalf("failed to unmarshal result: %v", err)
	}

	if attrs["max_minutes"] != nil {
		t.Errorf("expected max_minutes=null, got %v", attrs["max_minutes"])
	}
}

func TestSerializePointerSliceOfStrings(t *testing.T) {
	type TestModel struct {
		ID      types.String    `tfsdk:"id"`
		UserIDs *[]types.String `tfsdk:"user_ids"`
	}

	userIDs := []types.String{
		types.StringValue("u1"),
		types.StringValue("u2"),
	}

	model := TestModel{
		ID:      types.StringValue("abc"),
		UserIDs: &userIDs,
	}

	raw, err := SerializeToStateAttributes(model)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var attrs map[string]interface{}
	if err := json.Unmarshal(raw, &attrs); err != nil {
		t.Fatalf("failed to unmarshal result: %v", err)
	}

	ids, ok := attrs["user_ids"].([]interface{})
	if !ok {
		t.Fatalf("expected user_ids to be array, got %T", attrs["user_ids"])
	}
	if len(ids) != 2 {
		t.Errorf("expected 2 user_ids, got %d", len(ids))
	}
	if ids[0] != "u1" || ids[1] != "u2" {
		t.Errorf("unexpected user_ids values: %v", ids)
	}
}

func TestSerializeNilPointerSlice(t *testing.T) {
	type TestModel struct {
		ID      types.String    `tfsdk:"id"`
		UserIDs *[]types.String `tfsdk:"user_ids"`
	}

	model := TestModel{
		ID:      types.StringValue("abc"),
		UserIDs: nil,
	}

	raw, err := SerializeToStateAttributes(model)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var attrs map[string]interface{}
	if err := json.Unmarshal(raw, &attrs); err != nil {
		t.Fatalf("failed to unmarshal result: %v", err)
	}

	if attrs["user_ids"] != nil {
		t.Errorf("expected user_ids=null, got %v", attrs["user_ids"])
	}
}

func TestSerializeNestedStruct(t *testing.T) {
	type Inner struct {
		Name types.String `tfsdk:"name"`
		ID   types.String `tfsdk:"id"`
	}
	type TestModel struct {
		ID    types.String `tfsdk:"id"`
		Inner *Inner       `tfsdk:"inner"`
	}

	model := TestModel{
		ID: types.StringValue("abc"),
		Inner: &Inner{
			Name: types.StringValue("test"),
			ID:   types.StringValue("inner1"),
		},
	}

	raw, err := SerializeToStateAttributes(model)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var attrs map[string]interface{}
	if err := json.Unmarshal(raw, &attrs); err != nil {
		t.Fatalf("failed to unmarshal result: %v", err)
	}

	inner, ok := attrs["inner"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected inner to be object, got %T", attrs["inner"])
	}
	if inner["name"] != "test" {
		t.Errorf("expected inner.name=test, got %v", inner["name"])
	}
}

func TestSerializeNilNestedStruct(t *testing.T) {
	type Inner struct {
		Name types.String `tfsdk:"name"`
	}
	type TestModel struct {
		ID    types.String `tfsdk:"id"`
		Inner *Inner       `tfsdk:"inner"`
	}

	model := TestModel{
		ID:    types.StringValue("abc"),
		Inner: nil,
	}

	raw, err := SerializeToStateAttributes(model)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var attrs map[string]interface{}
	if err := json.Unmarshal(raw, &attrs); err != nil {
		t.Fatalf("failed to unmarshal result: %v", err)
	}

	if attrs["inner"] != nil {
		t.Errorf("expected inner=null, got %v", attrs["inner"])
	}
}

func TestSerializeSkipsFieldsWithoutTfsdkTag(t *testing.T) {
	type TestModel struct {
		ID       types.String `tfsdk:"id"`
		Internal string       // no tfsdk tag
	}

	model := TestModel{
		ID:       types.StringValue("abc"),
		Internal: "should be skipped",
	}

	raw, err := SerializeToStateAttributes(model)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var attrs map[string]interface{}
	if err := json.Unmarshal(raw, &attrs); err != nil {
		t.Fatalf("failed to unmarshal result: %v", err)
	}

	if len(attrs) != 1 {
		t.Errorf("expected 1 attribute, got %d: %v", len(attrs), attrs)
	}
}

func TestSerializePointerToSliceOfPointers(t *testing.T) {
	type Assignee struct {
		AssigneeID   types.String `tfsdk:"assignee_id"`
		AssigneeType types.String `tfsdk:"assignee_type"`
	}
	type TestModel struct {
		ID        types.String `tfsdk:"id"`
		Assignees *[]*Assignee `tfsdk:"assignees"`
	}

	assignees := []*Assignee{
		{AssigneeID: types.StringValue("u1"), AssigneeType: types.StringValue("user")},
		{AssigneeID: types.StringValue("g1"), AssigneeType: types.StringValue("group")},
	}

	model := TestModel{
		ID:        types.StringValue("abc"),
		Assignees: &assignees,
	}

	raw, err := SerializeToStateAttributes(model)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var attrs map[string]interface{}
	if err := json.Unmarshal(raw, &attrs); err != nil {
		t.Fatalf("failed to unmarshal result: %v", err)
	}

	assigneesArr, ok := attrs["assignees"].([]interface{})
	if !ok {
		t.Fatalf("expected assignees to be array, got %T", attrs["assignees"])
	}
	if len(assigneesArr) != 2 {
		t.Errorf("expected 2 assignees, got %d", len(assigneesArr))
	}

	first, ok := assigneesArr[0].(map[string]interface{})
	if !ok {
		t.Fatalf("expected first assignee to be object, got %T", assigneesArr[0])
	}
	if first["assignee_id"] != "u1" {
		t.Errorf("expected first assignee_id=u1, got %v", first["assignee_id"])
	}
}

func TestSerializePointer(t *testing.T) {
	type TestModel struct {
		ID   types.String `tfsdk:"id"`
		Name types.String `tfsdk:"name"`
	}

	model := &TestModel{
		ID:   types.StringValue("abc"),
		Name: types.StringValue("test"),
	}

	raw, err := SerializeToStateAttributes(model)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var attrs map[string]interface{}
	if err := json.Unmarshal(raw, &attrs); err != nil {
		t.Fatalf("failed to unmarshal result: %v", err)
	}

	if attrs["id"] != "abc" {
		t.Errorf("expected id=abc, got %v", attrs["id"])
	}
}
