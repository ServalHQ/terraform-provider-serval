// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_resource_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/serval-terraform/internal/services/access_resource"
	"github.com/stainless-sdks/serval-terraform/internal/test_helpers"
)

func TestAccessResourceDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*access_resource.AccessResourceDataSourceModel)(nil)
	schema := access_resource.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
