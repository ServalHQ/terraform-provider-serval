// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tag_test

import (
	"context"
	"testing"

	"github.com/ServalHQ/terraform-provider-serval/internal/services/tag"
	"github.com/ServalHQ/terraform-provider-serval/internal/test_helpers"
)

func TestTagModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*tag.TagModel)(nil)
	schema := tag.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
