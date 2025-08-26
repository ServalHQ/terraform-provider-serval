// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package entitlement_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/serval-terraform/internal/services/entitlement"
	"github.com/stainless-sdks/serval-terraform/internal/test_helpers"
)

func TestEntitlementModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*entitlement.EntitlementModel)(nil)
	schema := entitlement.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
