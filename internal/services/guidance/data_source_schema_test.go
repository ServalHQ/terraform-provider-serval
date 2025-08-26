// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package guidance_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/serval-terraform/internal/services/guidance"
	"github.com/stainless-sdks/serval-terraform/internal/test_helpers"
)

func TestGuidanceDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*guidance.GuidanceDataSourceModel)(nil)
	schema := guidance.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
