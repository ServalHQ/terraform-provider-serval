// Command testgen exercises the pkg/export API by reading the directgen test fixtures,
// generating Terraform state + HCL via the direct generation pipeline, and writing
// the output to a test workspace directory.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ServalHQ/terraform-provider-serval/pkg/export"
)

// testResource maps a resource type and name to its test fixture file.
type testResource struct {
	Type    string
	Name    string
	Fixture string
}

func main() {
	outDir := "test-workspace"
	if len(os.Args) > 1 {
		outDir = os.Args[1]
	}

	fixtureDir := filepath.Join("internal", "directgen", "testdata")

	// One of each resource type, matching the golden test fixtures.
	resources := []testResource{
		{export.ResourceTypeUser, "alice_smith", "user.json"},
		{export.ResourceTypeGroup, "engineering_admins", "group.json"},
		{export.ResourceTypeTeam, "platform_engineering", "team.json"},
		{export.ResourceTypeWorkflow, "deploy_to_prod", "workflow.json"},
		{export.ResourceTypeGuidance, "onboarding_guide", "guidance.json"},
		{export.ResourceTypeAccessPolicy, "standard_jit_access", "access_policy.json"},
		{export.ResourceTypeAppInstance, "okta_production", "app_instance.json"},
		{export.ResourceTypeAppResource, "github_organization", "app_resource.json"},
		{export.ResourceTypeAppResourceRole, "admin", "app_resource_role.json"},
		{export.ResourceTypeCustomService, "internal_api_gateway", "custom_service.json"},
	}

	// Build the input for Generate.
	var input []export.Resource
	for _, r := range resources {
		raw, err := os.ReadFile(filepath.Join(fixtureDir, r.Fixture))
		if err != nil {
			log.Fatalf("read fixture %s: %v", r.Fixture, err)
		}
		input = append(input, export.Resource{
			Type:    r.Type,
			Name:    r.Name,
			RawJSON: json.RawMessage(raw),
		})
	}

	// Use OpenTofu-compatible provider address.
	opts := export.GenerateOptions{
		ProviderAddress:  "registry.opentofu.org/servalhq/serval",
		TerraformVersion: "1.9.0",
		Lineage:          "testgen-e2e",
	}

	result, err := export.Generate(input, opts)
	if err != nil {
		log.Fatalf("Generate failed: %v", err)
	}

	// Create output directory.
	if err := os.MkdirAll(outDir, 0755); err != nil {
		log.Fatalf("mkdir %s: %v", outDir, err)
	}

	// Write terraform.tfstate.
	statePath := filepath.Join(outDir, "terraform.tfstate")
	prettyState, _ := json.MarshalIndent(json.RawMessage(result.StateJSON), "", "  ")
	if err := os.WriteFile(statePath, append(prettyState, '\n'), 0644); err != nil {
		log.Fatalf("write state: %v", err)
	}
	fmt.Printf("Wrote %s (%d bytes, %d resources)\n", statePath, len(prettyState), len(resources))

	// Write main.tf with provider config + generated resource blocks.
	providerBlock := `terraform {
  required_providers {
    serval = {
      source = "servalhq/serval"
    }
  }
}

provider "serval" {}

`
	hclPath := filepath.Join(outDir, "main.tf")
	hcl := providerBlock + result.HCL
	if err := os.WriteFile(hclPath, []byte(hcl), 0644); err != nil {
		log.Fatalf("write HCL: %v", err)
	}
	fmt.Printf("Wrote %s (%d bytes)\n", hclPath, len(hcl))

	fmt.Println("\nDirect generation complete. Supported types exercised:")
	for _, r := range resources {
		fmt.Printf("  %s.%s\n", r.Type, r.Name)
	}
}
