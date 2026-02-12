// Command e2etest fetches one of each resource type from a real Serval instance,
// runs direct generation to produce state + HCL, and writes the output to a
// workspace that can be verified with `tofu plan -refresh-only`.
//
// Authentication (pick one):
//
//	SERVAL_BEARER_TOKEN=<token> go run ./cmd/e2etest/
//
//	SERVAL_CLIENT_ID=<id> SERVAL_CLIENT_SECRET=<secret> go run ./cmd/e2etest/
//
// The program writes to e2e-workspace/ in the repo root.
package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/ServalHQ/terraform-provider-serval/pkg/export"
)

const baseURL = "https://public.api.serval.com"

var token string

// apiResource tracks a resource we fetched from the API.
type apiResource struct {
	Type     string // e.g., "serval_user"
	Name     string // sanitized TF name
	ID       string // resource ID
	Envelope json.RawMessage
}

func main() {
	token = os.Getenv("SERVAL_BEARER_TOKEN")
	if token == "" {
		// Try client credentials flow
		clientID := os.Getenv("SERVAL_CLIENT_ID")
		clientSecret := os.Getenv("SERVAL_CLIENT_SECRET")
		if clientID == "" || clientSecret == "" {
			log.Fatal("Set SERVAL_BEARER_TOKEN or both SERVAL_CLIENT_ID + SERVAL_CLIENT_SECRET")
		}
		fmt.Println("Exchanging client credentials for bearer token...")
		var err error
		token, err = exchangeClientCredentials(clientID, clientSecret)
		if err != nil {
			log.Fatalf("Token exchange failed: %v", err)
		}
		fmt.Println("Got bearer token")
	}

	outDir := "e2e-workspace"
	if len(os.Args) > 1 {
		outDir = os.Args[1]
	}

	var resources []apiResource

	// ---- Non-team-scoped resources ----

	fmt.Println("Fetching users...")
	if r, err := fetchFirst("/v2/users?pageSize=1", export.ResourceTypeUser, "name"); err == nil {
		resources = append(resources, r)
	} else {
		fmt.Printf("  skip: %v\n", err)
	}

	fmt.Println("Fetching teams...")
	if r, err := fetchFirst("/v2/teams?pageSize=1", export.ResourceTypeTeam, "name"); err == nil {
		resources = append(resources, r)
	} else {
		fmt.Printf("  skip: %v\n", err)
	}

	fmt.Println("Fetching groups...")
	if r, err := fetchFirst("/v2/groups?pageSize=1", export.ResourceTypeGroup, "name"); err == nil {
		resources = append(resources, r)
	} else {
		fmt.Printf("  skip: %v\n", err)
	}

	// ---- Get a team ID for team-scoped resources ----

	teamID := ""
	for _, r := range resources {
		if r.Type == export.ResourceTypeTeam {
			teamID = r.ID
			break
		}
	}

	if teamID != "" {
		fmt.Printf("Using team %s for team-scoped resources\n", teamID)

		teamScoped := []struct {
			path     string
			resType  string
			nameKey  string
		}{
			{"/v2/workflows?pageSize=1&teamId=" + teamID, export.ResourceTypeWorkflow, "name"},
			{"/v2/guidances?pageSize=1&teamId=" + teamID, export.ResourceTypeGuidance, "name"},
			{"/v2/access-policies?pageSize=1&teamId=" + teamID, export.ResourceTypeAccessPolicy, "name"},
			{"/v2/app-instances?pageSize=1&teamId=" + teamID, export.ResourceTypeAppInstance, "name"},
			{"/v2/app-resources?pageSize=1&teamId=" + teamID, export.ResourceTypeAppResource, "name"},
			{"/v2/app-resource-roles?pageSize=1&teamId=" + teamID, export.ResourceTypeAppResourceRole, "name"},
			{"/v2/custom-services?pageSize=1&teamId=" + teamID, export.ResourceTypeCustomService, "name"},
		}

		for _, ts := range teamScoped {
			typeName := strings.TrimPrefix(ts.resType, "serval_")
			fmt.Printf("Fetching %s...\n", typeName)
			if r, err := fetchFirst(ts.path, ts.resType, ts.nameKey); err == nil {
				resources = append(resources, r)
			} else {
				fmt.Printf("  skip: %v\n", err)
			}
		}
	} else {
		fmt.Println("No teams found — skipping team-scoped resources")
	}

	if len(resources) == 0 {
		log.Fatal("No resources found in this Serval instance")
	}

	fmt.Printf("\nFetched %d resources:\n", len(resources))
	for _, r := range resources {
		fmt.Printf("  %s.%s  (id=%s)\n", r.Type, r.Name, r.ID)
	}

	// ---- Run direct generation ----

	var input []export.Resource
	for _, r := range resources {
		input = append(input, export.Resource{
			Type:    r.Type,
			Name:    r.Name,
			RawJSON: r.Envelope,
		})
	}

	opts := export.GenerateOptions{
		ProviderAddress:  "registry.opentofu.org/servalhq/serval",
		TerraformVersion: "1.9.0",
		Lineage:          "e2e-directgen",
	}

	result, err := export.Generate(input, opts)
	if err != nil {
		log.Fatalf("Generate failed: %v", err)
	}

	// ---- Write workspace ----

	if err := os.MkdirAll(outDir, 0755); err != nil {
		log.Fatalf("mkdir: %v", err)
	}

	// State file
	statePath := filepath.Join(outDir, "terraform.tfstate")
	prettyState, _ := json.MarshalIndent(json.RawMessage(result.StateJSON), "", "  ")
	if err := os.WriteFile(statePath, append(prettyState, '\n'), 0644); err != nil {
		log.Fatalf("write state: %v", err)
	}
	fmt.Printf("\nWrote %s (%d bytes)\n", statePath, len(prettyState))

	// HCL config
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
	if err := os.WriteFile(hclPath, []byte(providerBlock+result.HCL), 0644); err != nil {
		log.Fatalf("write HCL: %v", err)
	}
	fmt.Printf("Wrote %s\n", hclPath)

	// Also write the raw API JSON for debugging
	rawDir := filepath.Join(outDir, "raw-api")
	os.MkdirAll(rawDir, 0755)
	for _, r := range resources {
		name := r.Type + "." + r.Name + ".json"
		pretty, _ := json.MarshalIndent(json.RawMessage(r.Envelope), "", "  ")
		os.WriteFile(filepath.Join(rawDir, name), pretty, 0644)
	}
	fmt.Printf("Wrote raw API responses to %s/\n", rawDir)

	fmt.Println("\n--- Next steps ---")
	fmt.Println("Verify directgen matches the provider's Read output:")
	fmt.Printf("  cd %s && tofu plan -refresh-only\n", outDir)
	fmt.Println("\nIf the plan shows no changes, directgen is producing correct state.")
	fmt.Println("If there are changes, the diff shows what directgen gets wrong.")
}

// exchangeClientCredentials performs the OAuth2 client credentials flow to
// obtain a bearer token from a client ID and secret.
func exchangeClientCredentials(clientID, clientSecret string) (string, error) {
	data := url.Values{"grant_type": {"client_credentials"}}
	req, err := http.NewRequest("POST", baseURL+"/v2/auth/token", strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString(
		[]byte(clientID+":"+clientSecret),
	))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body[:min(len(body), 300)]))
	}

	var tokenResp struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return "", fmt.Errorf("parse token response: %w", err)
	}
	if tokenResp.AccessToken == "" {
		return "", fmt.Errorf("empty access_token in response: %s", string(body[:min(len(body), 200)]))
	}
	return tokenResp.AccessToken, nil
}

// fetchFirst calls a list endpoint and returns the first resource wrapped in
// a single-resource envelope {"data": {...}}.
func fetchFirst(path, resourceType, nameKey string) (apiResource, error) {
	body, err := httpGet(path)
	if err != nil {
		return apiResource{}, err
	}

	// Parse list response: {"data": [...]}
	var listResp struct {
		Data []json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(body, &listResp); err != nil {
		return apiResource{}, fmt.Errorf("parse list response: %w", err)
	}
	if len(listResp.Data) == 0 {
		return apiResource{}, fmt.Errorf("no %s found", resourceType)
	}

	item := listResp.Data[0]

	// Extract id and name for the resource
	var fields map[string]json.RawMessage
	json.Unmarshal(item, &fields)

	id := unquote(fields["id"])
	name := unquote(fields[nameKey])
	if name == "" {
		name = id
	}

	// Wrap as single-resource envelope: {"data": {...}}
	envelope, _ := json.Marshal(map[string]json.RawMessage{"data": item})

	return apiResource{
		Type:     resourceType,
		Name:     sanitizeName(name),
		ID:       id,
		Envelope: envelope,
	}, nil
}

// httpGet makes an authenticated GET request to the Serval API.
func httpGet(path string) ([]byte, error) {
	req, err := http.NewRequest("GET", baseURL+path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body[:min(len(body), 200)]))
	}

	return body, nil
}

// unquote strips JSON quotes from a raw value.
func unquote(raw json.RawMessage) string {
	var s string
	json.Unmarshal(raw, &s)
	return s
}

var nonIdentRe = regexp.MustCompile(`[^a-zA-Z0-9_]`)

// sanitizeName converts a display name to a valid Terraform resource name.
func sanitizeName(name string) string {
	s := strings.ToLower(name)
	s = strings.ReplaceAll(s, " ", "_")
	s = strings.ReplaceAll(s, "-", "_")
	s = nonIdentRe.ReplaceAllString(s, "")
	if s == "" {
		s = "resource"
	}
	// Must start with a letter or underscore
	if s[0] >= '0' && s[0] <= '9' {
		s = "_" + s
	}
	return s
}
