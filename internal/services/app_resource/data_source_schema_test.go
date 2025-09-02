// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_resource_test

import (
	"context"
	"testing"

	"github.com/ServalHQ/terraform-provider-serval/internal/services/app_resource"
	"github.com/ServalHQ/terraform-provider-serval/internal/test_helpers"
)

func TestAppResourceDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*app_resource.AppResourceDataSourceModel)(nil)
	schema := app_resource.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
