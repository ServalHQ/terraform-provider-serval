// Package directgen provides golden-file differential tests that verify the direct
// generator produces correct state and HCL output for all supported resource types.
//
// These tests serve as Layer 3 correctness guarantees: they catch any drift between
// the direct generation output and what the provider model structs expect.
//
// To update golden files: go test ./internal/directgen/ -update-golden
package directgen

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var updateGolden = flag.Bool("update-golden", false, "update golden files")

// resourceTestCase defines a test case for a resource type.
type resourceTestCase struct {
	ResourceType string
	ResourceName string
	FixtureFile  string

	// State attribute assertions: field name -> expected value
	// Uses interface{} to handle string, float64, bool, nil, []interface{}, map[string]interface{}
	ExpectedStateAttrs map[string]interface{}

	// HCL assertions: strings that must appear in the generated HCL
	ExpectedHCLContains []string

	// HCL exclusions: strings that must NOT appear in the generated HCL
	ExpectedHCLExcludes []string
}

// allResourceTests returns test cases for all 11 resource types.
func allResourceTests() []resourceTestCase {
	return []resourceTestCase{
		{
			ResourceType: ResourceTypeUser,
			ResourceName: "alice_smith",
			FixtureFile:  "testdata/user.json",
			ExpectedStateAttrs: map[string]interface{}{
				"id":         "usr-abc-123",
				"email":      "alice@example.com",
				"first_name": "Alice",
				"last_name":  "Smith",
				"role":       "admin",
				"avatar_url": "https://example.com/avatar.png",
				"name":       "Alice Smith",
				"timezone":   "America/New_York",
			},
			ExpectedHCLContains: []string{
				`resource "serval_user" "alice_smith"`,
				`email = "alice@example.com"`,
			},
			ExpectedHCLExcludes: []string{
				// Computed fields should not be in HCL
				`id = "usr-abc-123"`,
				"created_at",
				"deactivated_at",
				`name = "Alice Smith"`,
				`timezone =`,
			},
		},
		{
			ResourceType: ResourceTypeGroup,
			ResourceName: "engineering_admins",
			FixtureFile:  "testdata/group.json",
			ExpectedStateAttrs: map[string]interface{}{
				"id":              "grp-abc-123",
				"name":            "Engineering Admins",
				"organization_id": "org-abc-123",
			},
			ExpectedHCLContains: []string{
				`resource "serval_group" "engineering_admins"`,
				`name = "Engineering Admins"`,
				`user_ids = ["usr-001", "usr-002", "usr-003"]`,
			},
			ExpectedHCLExcludes: []string{
				`id = "grp-abc-123"`,
				"created_at",
				"organization_id",
			},
		},
		{
			ResourceType: ResourceTypeTeam,
			ResourceName: "platform_engineering",
			FixtureFile:  "testdata/team.json",
			ExpectedStateAttrs: map[string]interface{}{
				"id":              "team-abc-123",
				"name":            "Platform Engineering",
				"description":     "The platform engineering team",
				"prefix":          "plat-eng",
				"organization_id": "org-abc-123",
			},
			ExpectedHCLContains: []string{
				`resource "serval_team" "platform_engineering"`,
				`name = "Platform Engineering"`,
				`description = "The platform engineering team"`,
				`prefix = "plat-eng"`,
			},
			ExpectedHCLExcludes: []string{
				`id = "team-abc-123"`,
				"created_at",
				"organization_id",
			},
		},
		{
			ResourceType: ResourceTypeWorkflow,
			ResourceName: "deploy_to_prod",
			FixtureFile:  "testdata/workflow.json",
			ExpectedStateAttrs: map[string]interface{}{
				"id":      "wf-abc-123",
				"team_id": "team-abc-123",
				"name":    "deploy-to-prod",
				"type":    "EXECUTABLE",
			},
			ExpectedHCLContains: []string{
				`resource "serval_workflow" "deploy_to_prod"`,
				`team_id = "team-abc-123"`,
				`name = "deploy-to-prod"`,
				`type = "EXECUTABLE"`,
				"content =", // content should be included (required field)
			},
			ExpectedHCLExcludes: []string{
				`id = "wf-abc-123"`,
				"has_unpublished_changes",
			},
		},
		{
			ResourceType: ResourceTypeGuidance,
			ResourceName: "onboarding_guide",
			FixtureFile:  "testdata/guidance.json",
			ExpectedStateAttrs: map[string]interface{}{
				"id":      "guid-abc-123",
				"team_id": "team-abc-123",
				"name":    "Onboarding Guide",
				"content": "# Welcome\nThis is the onboarding guide.",
			},
			ExpectedHCLContains: []string{
				`resource "serval_guidance" "onboarding_guide"`,
				`team_id = "team-abc-123"`,
				`name = "Onboarding Guide"`,
				"content =",
			},
			ExpectedHCLExcludes: []string{
				`id = "guid-abc-123"`,
				"has_unpublished_changes",
			},
		},
		{
			ResourceType: ResourceTypeAccessPolicy,
			ResourceName: "standard_jit_access",
			FixtureFile:  "testdata/access_policy.json",
			ExpectedStateAttrs: map[string]interface{}{
				"id":                             "ap-abc-123",
				"team_id":                        "team-abc-123",
				"name":                           "Standard JIT Access",
				"max_access_minutes":             float64(60),
				"recommended_access_minutes":     float64(30),
				"require_business_justification": true,
			},
			ExpectedHCLContains: []string{
				`resource "serval_access_policy" "standard_jit_access"`,
				`team_id = "team-abc-123"`,
				`name = "Standard JIT Access"`,
				`max_access_minutes = 60`,
				`recommended_access_minutes = 30`,
				`require_business_justification = true`,
			},
			ExpectedHCLExcludes: []string{
				`id = "ap-abc-123"`,
			},
		},
		{
			ResourceType: ResourceTypeAppInstance,
			ResourceName: "okta_production",
			FixtureFile:  "testdata/app_instance.json",
			ExpectedStateAttrs: map[string]interface{}{
				"id":                           "ai-abc-123",
				"team_id":                      "team-abc-123",
				"service":                      "okta",
				"external_service_instance_id": "0oa1234567",
				"name":                         "Okta Production",
				"default_access_policy_id":     "ap-abc-123",
				"access_requests_enabled":      true,
			},
			ExpectedHCLContains: []string{
				`resource "serval_app_instance" "okta_production"`,
				`team_id = "team-abc-123"`,
				`service = "okta"`,
				`name = "Okta Production"`,
				`external_service_instance_id = "0oa1234567"`,
			},
			ExpectedHCLExcludes: []string{
				`id = "ai-abc-123"`,
			},
		},
		{
			ResourceType: ResourceTypeAppResource,
			ResourceName: "github_organization",
			FixtureFile:  "testdata/app_resource.json",
			ExpectedStateAttrs: map[string]interface{}{
				"id":              "ar-abc-123",
				"app_instance_id": "ai-abc-123",
				"name":            "GitHub Organization",
				"resource_type":   "github_org",
				"description":     "Main GitHub organization",
				"external_id":     "ext-gh-org-1",
			},
			ExpectedHCLContains: []string{
				`resource "serval_app_resource" "github_organization"`,
				`app_instance_id = "ai-abc-123"`,
				`name = "GitHub Organization"`,
				`resource_type = "github_org"`,
			},
			ExpectedHCLExcludes: []string{
				`id = "ar-abc-123"`,
			},
		},
		{
			ResourceType: ResourceTypeAppResourceRole,
			ResourceName: "admin",
			FixtureFile:  "testdata/app_resource_role.json",
			ExpectedStateAttrs: map[string]interface{}{
				"id":               "arr-abc-123",
				"resource_id":      "ar-abc-123",
				"name":             "admin",
				"access_policy_id": "ap-abc-123",
				"description":      "Admin role for GitHub org",
				"requests_enabled": true,
			},
			ExpectedHCLContains: []string{
				`resource "serval_app_resource_role" "admin"`,
				`resource_id = "ar-abc-123"`,
				`name = "admin"`,
				"provisioning_method = {",
				"custom_workflow = {",
			},
			ExpectedHCLExcludes: []string{
				`id = "arr-abc-123"`,
			},
		},
		{
			ResourceType: ResourceTypeCustomService,
			ResourceName: "internal_api_gateway",
			FixtureFile:  "testdata/custom_service.json",
			ExpectedStateAttrs: map[string]interface{}{
				"id":      "cs-abc-123",
				"team_id": "team-abc-123",
				"name":    "Internal API Gateway",
				"domain":  "api.internal.example.com",
			},
			ExpectedHCLContains: []string{
				`resource "serval_custom_service" "internal_api_gateway"`,
				`team_id = "team-abc-123"`,
				`name = "Internal API Gateway"`,
				`domain = "api.internal.example.com"`,
			},
			ExpectedHCLExcludes: []string{
				`id = "cs-abc-123"`,
			},
		},
	}
}

func TestGoldenAllResourceTypes(t *testing.T) {
	opts := DefaultOptions()

	for _, tc := range allResourceTests() {
		t.Run(tc.ResourceType, func(t *testing.T) {
			// Load fixture
			rawJSON, err := os.ReadFile(tc.FixtureFile)
			if err != nil {
				t.Fatalf("failed to read fixture %s: %v", tc.FixtureFile, err)
			}

			// Generate
			stateAttrs, hcl, err := GenerateForSingleResource(tc.ResourceType, tc.ResourceName, rawJSON, opts)
			if err != nil {
				t.Fatalf("generation failed: %v", err)
			}

			// Write golden files if requested
			if *updateGolden {
				goldenState := filepath.Join("testdata", tc.ResourceType+".state.golden.json")
				goldenHCL := filepath.Join("testdata", tc.ResourceType+".hcl.golden")

				prettyState, _ := json.MarshalIndent(json.RawMessage(stateAttrs), "", "  ")
				if err := os.WriteFile(goldenState, prettyState, 0644); err != nil {
					t.Fatalf("failed to write golden state: %v", err)
				}
				if err := os.WriteFile(goldenHCL, []byte(hcl), 0644); err != nil {
					t.Fatalf("failed to write golden HCL: %v", err)
				}
			}

			// Verify state attributes
			var attrs map[string]interface{}
			if err := json.Unmarshal(stateAttrs, &attrs); err != nil {
				t.Fatalf("failed to unmarshal state attrs: %v", err)
			}

			for key, expected := range tc.ExpectedStateAttrs {
				actual, exists := attrs[key]
				if !exists {
					t.Errorf("state: missing expected attribute %q", key)
					continue
				}
				if !deepEqual(expected, actual) {
					t.Errorf("state: attribute %q: expected %v (%T), got %v (%T)",
						key, expected, expected, actual, actual)
				}
			}

			// Verify HCL contains expected strings
			for _, expected := range tc.ExpectedHCLContains {
				if !strings.Contains(hcl, expected) {
					t.Errorf("HCL should contain %q but doesn't.\nHCL:\n%s", expected, hcl)
				}
			}

			// Verify HCL excludes certain strings
			for _, excluded := range tc.ExpectedHCLExcludes {
				if strings.Contains(hcl, excluded) {
					t.Errorf("HCL should NOT contain %q but does.\nHCL:\n%s", excluded, hcl)
				}
			}

			// Compare with golden files if they exist
			goldenState := filepath.Join("testdata", tc.ResourceType+".state.golden.json")
			goldenHCL := filepath.Join("testdata", tc.ResourceType+".hcl.golden")

			if goldenData, err := os.ReadFile(goldenState); err == nil {
				var expectedAttrs, actualAttrs interface{}
				if err := json.Unmarshal(goldenData, &expectedAttrs); err != nil {
					t.Fatalf("unmarshal golden state: %v", err)
				}
				if err := json.Unmarshal(stateAttrs, &actualAttrs); err != nil {
					t.Fatalf("unmarshal actual state: %v", err)
				}

				expectedJSON, _ := json.MarshalIndent(expectedAttrs, "", "  ")
				actualJSON, _ := json.MarshalIndent(actualAttrs, "", "  ")

				if string(expectedJSON) != string(actualJSON) {
					t.Errorf("state does not match golden file.\nExpected:\n%s\nActual:\n%s",
						string(expectedJSON), string(actualJSON))
				}
			}

			if goldenData, err := os.ReadFile(goldenHCL); err == nil {
				if string(goldenData) != hcl {
					t.Errorf("HCL does not match golden file.\nExpected:\n%s\nActual:\n%s",
						string(goldenData), hcl)
				}
			}
		})
	}
}

// TestGoldenFullStateFile tests that Generate produces a valid complete state file.
func TestGoldenFullStateFile(t *testing.T) {
	testCases := allResourceTests()
	resources := make([]Resource, 0, len(testCases))

	for _, tc := range testCases {
		rawJSON, err := os.ReadFile(tc.FixtureFile)
		if err != nil {
			t.Fatalf("failed to read fixture %s: %v", tc.FixtureFile, err)
		}
		resources = append(resources, Resource{
			Type:    tc.ResourceType,
			Name:    tc.ResourceName,
			RawJSON: rawJSON,
		})
	}

	opts := GenerateOptions{
		ProviderAddress:  "registry.opentofu.org/servalhq/serval",
		TerraformVersion: "1.9.0",
		Lineage:          "golden-test",
	}

	result, err := Generate(resources, opts)
	if err != nil {
		t.Fatalf("generation failed: %v", err)
	}

	// Verify state file structure
	var state map[string]interface{}
	if err := json.Unmarshal(result.StateJSON, &state); err != nil {
		t.Fatalf("invalid state JSON: %v", err)
	}

	if state["version"] != float64(4) {
		t.Errorf("expected version=4, got %v", state["version"])
	}
	if state["terraform_version"] != "1.9.0" {
		t.Errorf("expected terraform_version=1.9.0, got %v", state["terraform_version"])
	}

	stateResources, ok := state["resources"].([]interface{})
	if !ok {
		t.Fatal("resources is not an array")
	}

	expectedResourceCount := len(testCases)
	if len(stateResources) != expectedResourceCount {
		t.Errorf("expected %d resources, got %d", expectedResourceCount, len(stateResources))
	}

	// Verify each resource has required fields
	for _, r := range stateResources {
		res, ok := r.(map[string]interface{})
		if !ok {
			t.Error("resource is not an object")
			continue
		}
		if res["mode"] != "managed" {
			t.Errorf("expected mode=managed, got %v", res["mode"])
		}
		if res["type"] == nil || res["type"] == "" {
			t.Error("resource type is empty")
		}
		if res["name"] == nil || res["name"] == "" {
			t.Error("resource name is empty")
		}
		if !strings.Contains(fmt.Sprint(res["provider"]), "serval") {
			t.Errorf("unexpected provider: %v", res["provider"])
		}
	}

	// Verify HCL contains all resource blocks
	for _, tc := range testCases {
		expected := fmt.Sprintf(`resource %q %q`, tc.ResourceType, tc.ResourceName)
		if !strings.Contains(result.HCL, expected) {
			t.Errorf("HCL missing resource block for %s.%s", tc.ResourceType, tc.ResourceName)
		}
	}

	// Write golden file for full state if requested
	if *updateGolden {
		goldenFile := filepath.Join("testdata", "full_state.golden.json")
		prettyState, _ := json.MarshalIndent(json.RawMessage(result.StateJSON), "", "  ")
		if err := os.WriteFile(goldenFile, prettyState, 0644); err != nil {
			t.Fatalf("failed to write golden state: %v", err)
		}

		goldenHCL := filepath.Join("testdata", "full_state.hcl.golden")
		if err := os.WriteFile(goldenHCL, []byte(result.HCL), 0644); err != nil {
			t.Fatalf("failed to write golden HCL: %v", err)
		}
	}
}

// TestGoldenNullOptionalFields verifies that null optional fields are handled correctly.
func TestGoldenNullOptionalFields(t *testing.T) {
	// User with minimal fields (only required + computed)
	rawJSON := json.RawMessage(`{
		"data": {
			"id": "usr-minimal",
			"email": "minimal@example.com",
			"createdAt": "2025-01-01T00:00:00Z",
			"name": "minimal"
		}
	}`)

	opts := DefaultOptions()
	stateAttrs, hcl, err := GenerateForSingleResource(ResourceTypeUser, "minimal", rawJSON, opts)
	if err != nil {
		t.Fatalf("generation failed: %v", err)
	}

	// State should have null optional fields
	var attrs map[string]interface{}
	if err := json.Unmarshal(stateAttrs, &attrs); err != nil {
		t.Fatalf("unmarshal state attrs: %v", err)
	}

	if attrs["first_name"] != nil {
		t.Errorf("expected first_name=null, got %v", attrs["first_name"])
	}
	if attrs["last_name"] != nil {
		t.Errorf("expected last_name=null, got %v", attrs["last_name"])
	}

	// HCL should only include required field (email)
	if !strings.Contains(hcl, `email = "minimal@example.com"`) {
		t.Errorf("missing required field email in HCL:\n%s", hcl)
	}
	if strings.Contains(hcl, "first_name") {
		t.Errorf("should not include null optional field first_name in HCL:\n%s", hcl)
	}
}

// deepEqual compares two values for equality, handling JSON number comparison.
func deepEqual(expected, actual interface{}) bool {
	// Handle nil
	if expected == nil && actual == nil {
		return true
	}
	if expected == nil || actual == nil {
		return false
	}

	// Handle slices
	expectedSlice, expectedIsSlice := expected.([]interface{})
	actualSlice, actualIsSlice := actual.([]interface{})
	if expectedIsSlice && actualIsSlice {
		if len(expectedSlice) != len(actualSlice) {
			return false
		}
		for i := range expectedSlice {
			if !deepEqual(expectedSlice[i], actualSlice[i]) {
				return false
			}
		}
		return true
	}

	// Handle maps
	expectedMap, expectedIsMap := expected.(map[string]interface{})
	actualMap, actualIsMap := actual.(map[string]interface{})
	if expectedIsMap && actualIsMap {
		if len(expectedMap) != len(actualMap) {
			return false
		}
		for k, v := range expectedMap {
			if !deepEqual(v, actualMap[k]) {
				return false
			}
		}
		return true
	}

	return fmt.Sprint(expected) == fmt.Sprint(actual)
}
