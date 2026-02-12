// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package guidance_test

import (
	"context"
	"testing"

	"github.com/ServalHQ/terraform-provider-serval/internal/services/guidance"
	"github.com/ServalHQ/terraform-provider-serval/internal/test_helpers"
)

func TestGuidancesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*guidance.GuidancesDataSourceModel)(nil)
	schema := guidance.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
