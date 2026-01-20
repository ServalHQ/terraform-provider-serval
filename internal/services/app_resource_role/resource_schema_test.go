// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_resource_role_test

import (
	"context"
	"testing"

	"github.com/ServalHQ/terraform-provider-serval/internal/services/app_resource_role"
	"github.com/ServalHQ/terraform-provider-serval/internal/test_helpers"
)

func TestAppResourceRoleModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*app_resource_role.AppResourceRoleModel)(nil)
	schema := app_resource_role.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
