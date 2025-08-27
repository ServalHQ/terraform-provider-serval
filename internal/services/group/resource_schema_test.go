// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package group_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/serval-terraform/internal/services/group"
	"github.com/stainless-sdks/serval-terraform/internal/test_helpers"
)

func TestGroupModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*group.GroupModel)(nil)
	schema := group.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
