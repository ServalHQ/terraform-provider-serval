// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tag_test

import (
	"context"
	"testing"

	"github.com/ServalHQ/terraform-provider-serval/internal/services/tag"
	"github.com/ServalHQ/terraform-provider-serval/internal/test_helpers"
)

func TestTagsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*tag.TagsDataSourceModel)(nil)
	schema := tag.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
