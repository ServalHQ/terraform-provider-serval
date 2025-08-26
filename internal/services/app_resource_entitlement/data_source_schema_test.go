// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_resource_entitlement_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/serval-terraform/internal/services/app_resource_entitlement"
	"github.com/stainless-sdks/serval-terraform/internal/test_helpers"
)

func TestAppResourceEntitlementDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*app_resource_entitlement.AppResourceEntitlementDataSourceModel)(nil)
	schema := app_resource_entitlement.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
