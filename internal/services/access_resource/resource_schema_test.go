// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_resource_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/serval-terraform/internal/services/access_resource"
	"github.com/stainless-sdks/serval-terraform/internal/test_helpers"
)

func TestAccessResourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*access_resource.AccessResourceModel)(nil)
	schema := access_resource.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
