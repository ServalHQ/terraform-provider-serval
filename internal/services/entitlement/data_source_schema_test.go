// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package entitlement_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/serval-terraform/internal/services/entitlement"
	"github.com/stainless-sdks/serval-terraform/internal/test_helpers"
)

func TestEntitlementDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*entitlement.EntitlementDataSourceModel)(nil)
	schema := entitlement.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
