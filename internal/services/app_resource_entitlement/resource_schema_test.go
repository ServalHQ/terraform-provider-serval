// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_resource_entitlement_test

import (
	"context"
	"testing"

	"github.com/ServalHQ/terraform-provider-serval/internal/services/app_resource_entitlement"
	"github.com/ServalHQ/terraform-provider-serval/internal/test_helpers"
)

func TestAppResourceEntitlementModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*app_resource_entitlement.AppResourceEntitlementModel)(nil)
	schema := app_resource_entitlement.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
