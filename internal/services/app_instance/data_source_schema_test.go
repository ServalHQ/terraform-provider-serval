// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_instance_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/serval-terraform/internal/services/app_instance"
	"github.com/stainless-sdks/serval-terraform/internal/test_helpers"
)

func TestAppInstanceDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*app_instance.AppInstanceDataSourceModel)(nil)
	schema := app_instance.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
