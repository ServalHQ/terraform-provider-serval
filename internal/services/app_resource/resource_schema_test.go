// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_resource_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/serval-terraform/internal/services/app_resource"
	"github.com/stainless-sdks/serval-terraform/internal/test_helpers"
)

func TestAppResourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*app_resource.AppResourceModel)(nil)
	schema := app_resource.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
