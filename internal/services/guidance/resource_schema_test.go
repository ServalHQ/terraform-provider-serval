// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package guidance_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/serval-terraform/internal/services/guidance"
	"github.com/stainless-sdks/serval-terraform/internal/test_helpers"
)

func TestGuidanceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*guidance.GuidanceModel)(nil)
	schema := guidance.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
