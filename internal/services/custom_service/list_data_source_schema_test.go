// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_service_test

import (
	"context"
	"testing"

	"github.com/ServalHQ/terraform-provider-serval/internal/services/custom_service"
	"github.com/ServalHQ/terraform-provider-serval/internal/test_helpers"
)

func TestCustomServicesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*custom_service.CustomServicesDataSourceModel)(nil)
	schema := custom_service.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
