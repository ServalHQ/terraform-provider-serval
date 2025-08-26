// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package resource_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/serval-terraform/internal/services/resource"
	"github.com/stainless-sdks/serval-terraform/internal/test_helpers"
)

func TestResourceDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*resource.ResourceDataSourceModel)(nil)
	schema := resource.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
