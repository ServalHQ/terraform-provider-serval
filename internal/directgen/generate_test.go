package directgen

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestGenerateForSingleUser(t *testing.T) {
	rawJSON := json.RawMessage(`{
		"data": {
			"id": "usr-001",
			"email": "alice@example.com",
			"firstName": "Alice",
			"lastName": "Smith",
			"role": "admin",
			"avatarUrl": "https://example.com/avatar.png",
			"createdAt": "2025-01-15T10:00:00Z",
			"name": "Alice Smith",
			"timezone": "America/New_York"
		}
	}`)

	opts := DefaultOptions()
	attrs, hcl, err := GenerateForSingleResource(ResourceTypeUser, "alice_smith", rawJSON, opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Check state attributes
	var stateAttrs map[string]interface{}
	if err := json.Unmarshal(attrs, &stateAttrs); err != nil {
		t.Fatalf("failed to unmarshal state attrs: %v", err)
	}

	if stateAttrs["id"] != "usr-001" {
		t.Errorf("expected id=usr-001, got %v", stateAttrs["id"])
	}
	if stateAttrs["email"] != "alice@example.com" {
		t.Errorf("expected email=alice@example.com, got %v", stateAttrs["email"])
	}
	if stateAttrs["first_name"] != "Alice" {
		t.Errorf("expected first_name=Alice, got %v", stateAttrs["first_name"])
	}

	// Check HCL
	if !strings.Contains(hcl, `resource "serval_user" "alice_smith"`) {
		t.Errorf("missing resource block in HCL:\n%s", hcl)
	}
	if !strings.Contains(hcl, `email = "alice@example.com"`) {
		t.Errorf("missing required field email in HCL:\n%s", hcl)
	}
	// Computed fields should NOT be in HCL
	if strings.Contains(hcl, `id = "usr-001"`) {
		t.Errorf("computed field id should not be in HCL:\n%s", hcl)
	}
	if strings.Contains(hcl, "created_at") {
		t.Errorf("computed field created_at should not be in HCL:\n%s", hcl)
	}
}

func TestGenerateForSingleTeam(t *testing.T) {
	rawJSON := json.RawMessage(`{
		"data": {
			"id": "team-001",
			"name": "Engineering",
			"description": "The engineering team",
			"prefix": "eng",
			"createdAt": "2025-01-10T09:00:00Z",
			"organizationId": "org-001"
		}
	}`)

	opts := DefaultOptions()
	attrs, hcl, err := GenerateForSingleResource(ResourceTypeTeam, "engineering", rawJSON, opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var stateAttrs map[string]interface{}
	if err := json.Unmarshal(attrs, &stateAttrs); err != nil {
		t.Fatalf("failed to unmarshal state attrs: %v", err)
	}

	if stateAttrs["id"] != "team-001" {
		t.Errorf("expected id=team-001, got %v", stateAttrs["id"])
	}
	if stateAttrs["name"] != "Engineering" {
		t.Errorf("expected name=Engineering, got %v", stateAttrs["name"])
	}
	if stateAttrs["prefix"] != "eng" {
		t.Errorf("expected prefix=eng, got %v", stateAttrs["prefix"])
	}

	// HCL should include required + optional fields, not computed
	if !strings.Contains(hcl, `name = "Engineering"`) {
		t.Errorf("missing name in HCL:\n%s", hcl)
	}
	if !strings.Contains(hcl, `description = "The engineering team"`) {
		t.Errorf("missing description in HCL:\n%s", hcl)
	}
	if strings.Contains(hcl, "created_at") {
		t.Errorf("computed field should not be in HCL:\n%s", hcl)
	}
}

func TestGenerateForAccessPolicy(t *testing.T) {
	rawJSON := json.RawMessage(`{
		"data": {
			"id": "ap-001",
			"teamId": "team-001",
			"name": "Standard Access",
			"description": "Standard access policy",
			"maxAccessMinutes": 60,
			"recommendedAccessMinutes": 30,
			"requireBusinessJustification": true
		}
	}`)

	opts := DefaultOptions()
	attrs, hcl, err := GenerateForSingleResource(ResourceTypeAccessPolicy, "standard_access", rawJSON, opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var stateAttrs map[string]interface{}
	if err := json.Unmarshal(attrs, &stateAttrs); err != nil {
		t.Fatalf("failed to unmarshal state attrs: %v", err)
	}

	if stateAttrs["team_id"] != "team-001" {
		t.Errorf("expected team_id=team-001, got %v", stateAttrs["team_id"])
	}
	if stateAttrs["max_access_minutes"] != float64(60) {
		t.Errorf("expected max_access_minutes=60, got %v", stateAttrs["max_access_minutes"])
	}
	if stateAttrs["require_business_justification"] != true {
		t.Errorf("expected require_business_justification=true, got %v", stateAttrs["require_business_justification"])
	}

	// HCL checks
	if !strings.Contains(hcl, `team_id = "team-001"`) {
		t.Errorf("missing required field team_id in HCL:\n%s", hcl)
	}
	if !strings.Contains(hcl, "max_access_minutes = 60") {
		t.Errorf("missing optional field max_access_minutes in HCL:\n%s", hcl)
	}
	if !strings.Contains(hcl, "require_business_justification = true") {
		t.Errorf("missing computed_optional field in HCL:\n%s", hcl)
	}
}

func TestGenerateForGroup(t *testing.T) {
	rawJSON := json.RawMessage(`{
		"data": {
			"id": "grp-001",
			"name": "Admins",
			"userIds": ["usr-001", "usr-002"],
			"createdAt": "2025-01-15T10:00:00Z",
			"organizationId": "org-001"
		}
	}`)

	opts := DefaultOptions()
	attrs, hcl, err := GenerateForSingleResource(ResourceTypeGroup, "admins", rawJSON, opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var stateAttrs map[string]interface{}
	if err := json.Unmarshal(attrs, &stateAttrs); err != nil {
		t.Fatalf("failed to unmarshal state attrs: %v", err)
	}

	userIDs, ok := stateAttrs["user_ids"].([]interface{})
	if !ok {
		t.Fatalf("expected user_ids to be array, got %T: %v", stateAttrs["user_ids"], stateAttrs["user_ids"])
	}
	if len(userIDs) != 2 {
		t.Errorf("expected 2 user_ids, got %d", len(userIDs))
	}

	// HCL should include user_ids
	if !strings.Contains(hcl, `user_ids = ["usr-001", "usr-002"]`) {
		t.Errorf("missing user_ids in HCL:\n%s", hcl)
	}
}

func TestGenerateMultipleResources(t *testing.T) {
	resources := []Resource{
		{
			Type: ResourceTypeUser,
			Name: "alice",
			RawJSON: json.RawMessage(`{
				"data": {
					"id": "usr-001",
					"email": "alice@example.com",
					"createdAt": "2025-01-15T10:00:00Z",
					"name": "alice"
				}
			}`),
		},
		{
			Type: ResourceTypeTeam,
			Name: "eng",
			RawJSON: json.RawMessage(`{
				"data": {
					"id": "team-001",
					"name": "Engineering",
					"createdAt": "2025-01-10T09:00:00Z"
				}
			}`),
		},
	}

	opts := DefaultOptions()
	result, err := Generate(resources, opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Verify state file structure
	var state map[string]interface{}
	if err := json.Unmarshal(result.StateJSON, &state); err != nil {
		t.Fatalf("failed to unmarshal state: %v", err)
	}

	if state["version"] != float64(4) {
		t.Errorf("expected version=4, got %v", state["version"])
	}

	stateResources, ok := state["resources"].([]interface{})
	if !ok {
		t.Fatalf("expected resources to be array")
	}
	if len(stateResources) != 2 {
		t.Errorf("expected 2 resources in state, got %d", len(stateResources))
	}

	// Verify HCL contains both resources
	if !strings.Contains(result.HCL, `resource "serval_user" "alice"`) {
		t.Errorf("missing user resource in HCL")
	}
	if !strings.Contains(result.HCL, `resource "serval_team" "eng"`) {
		t.Errorf("missing team resource in HCL")
	}
}

func TestGenerateUnsupportedType(t *testing.T) {
	_, _, err := GenerateForSingleResource("serval_unknown", "test", json.RawMessage(`{}`), DefaultOptions())
	if err == nil {
		t.Error("expected error for unsupported type")
	}
	if !strings.Contains(err.Error(), "unsupported resource type") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestGenerateForAppInstance(t *testing.T) {
	rawJSON := json.RawMessage(`{
		"data": {
			"id": "ai-001",
			"teamId": "team-001",
			"service": "okta",
			"externalServiceInstanceId": "ext-123",
			"name": "Okta Production",
			"accessRequestsEnabled": true
		}
	}`)

	opts := DefaultOptions()
	attrs, hcl, err := GenerateForSingleResource(ResourceTypeAppInstance, "okta_prod", rawJSON, opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var stateAttrs map[string]interface{}
	if err := json.Unmarshal(attrs, &stateAttrs); err != nil {
		t.Fatalf("failed to unmarshal state attrs: %v", err)
	}

	if stateAttrs["service"] != "okta" {
		t.Errorf("expected service=okta, got %v", stateAttrs["service"])
	}
	if stateAttrs["access_requests_enabled"] != true {
		t.Errorf("expected access_requests_enabled=true, got %v", stateAttrs["access_requests_enabled"])
	}

	if !strings.Contains(hcl, `name = "Okta Production"`) {
		t.Errorf("missing name in HCL:\n%s", hcl)
	}
}

func TestGenerateForCustomService(t *testing.T) {
	rawJSON := json.RawMessage(`{
		"data": {
			"id": "cs-001",
			"teamId": "team-001",
			"name": "Internal API",
			"domain": "api.internal.com"
		}
	}`)

	opts := DefaultOptions()
	attrs, hcl, err := GenerateForSingleResource(ResourceTypeCustomService, "internal_api", rawJSON, opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var stateAttrs map[string]interface{}
	if err := json.Unmarshal(attrs, &stateAttrs); err != nil {
		t.Fatalf("failed to unmarshal state attrs: %v", err)
	}

	if stateAttrs["name"] != "Internal API" {
		t.Errorf("expected name=Internal API, got %v", stateAttrs["name"])
	}
	if stateAttrs["domain"] != "api.internal.com" {
		t.Errorf("expected domain=api.internal.com, got %v", stateAttrs["domain"])
	}

	if !strings.Contains(hcl, `domain = "api.internal.com"`) {
		t.Errorf("missing domain in HCL:\n%s", hcl)
	}
}

func TestGenerateStateFileStructure(t *testing.T) {
	resources := []Resource{
		{
			Type: ResourceTypeUser,
			Name: "alice",
			RawJSON: json.RawMessage(`{
				"data": {
					"id": "usr-001",
					"email": "alice@example.com",
					"createdAt": "2025-01-15T10:00:00Z",
					"name": "alice"
				}
			}`),
		},
	}

	opts := GenerateOptions{
		ProviderAddress:  "registry.opentofu.org/servalhq/serval",
		TerraformVersion: "1.9.0",
		Lineage:          "test-lineage-123",
	}

	result, err := Generate(resources, opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var state stateFile
	if err := json.Unmarshal(result.StateJSON, &state); err != nil {
		t.Fatalf("failed to unmarshal state: %v", err)
	}

	if state.Version != 4 {
		t.Errorf("expected version=4, got %d", state.Version)
	}
	if state.TerraformVersion != "1.9.0" {
		t.Errorf("expected terraform_version=1.9.0, got %s", state.TerraformVersion)
	}
	if state.Lineage != "test-lineage-123" {
		t.Errorf("expected lineage=test-lineage-123, got %s", state.Lineage)
	}
	if state.Serial != 1 {
		t.Errorf("expected serial=1, got %d", state.Serial)
	}

	if len(state.Resources) != 1 {
		t.Fatalf("expected 1 resource, got %d", len(state.Resources))
	}

	res := state.Resources[0]
	if res.Mode != "managed" {
		t.Errorf("expected mode=managed, got %s", res.Mode)
	}
	if res.Type != "serval_user" {
		t.Errorf("expected type=serval_user, got %s", res.Type)
	}
	if res.Name != "alice" {
		t.Errorf("expected name=alice, got %s", res.Name)
	}
	if !strings.Contains(res.Provider, "registry.opentofu.org/servalhq/serval") {
		t.Errorf("unexpected provider: %s", res.Provider)
	}
}
