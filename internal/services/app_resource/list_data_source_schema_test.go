// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_resource_test

import (
	"context"
	"testing"

	"github.com/ServalHQ/terraform-provider-serval/internal/services/app_resource"
	"github.com/ServalHQ/terraform-provider-serval/internal/test_helpers"
)

func TestAppResourcesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*app_resource.AppResourcesDataSourceModel)(nil)
	schema := app_resource.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
