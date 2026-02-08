// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_request_test

import (
	"context"
	"testing"

	"github.com/ServalHQ/terraform-provider-serval/internal/services/access_request"
	"github.com/ServalHQ/terraform-provider-serval/internal/test_helpers"
)

func TestAccessRequestsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*access_request.AccessRequestsDataSourceModel)(nil)
	schema := access_request.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
