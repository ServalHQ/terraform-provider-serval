// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package group_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/serval-terraform/internal/services/group"
	"github.com/stainless-sdks/serval-terraform/internal/test_helpers"
)

func TestGroupDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*group.GroupDataSourceModel)(nil)
	schema := group.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
