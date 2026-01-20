// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_resource_role_test

import (
	"context"
	"testing"

	"github.com/ServalHQ/terraform-provider-serval/internal/services/app_resource_role"
	"github.com/ServalHQ/terraform-provider-serval/internal/test_helpers"
)

func TestAppResourceRoleDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*app_resource_role.AppResourceRoleDataSourceModel)(nil)
	schema := app_resource_role.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
