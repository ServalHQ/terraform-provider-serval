// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package group_test

import (
	"context"
	"testing"

	"github.com/ServalHQ/terraform-provider-serval/internal/services/group"
	"github.com/ServalHQ/terraform-provider-serval/internal/test_helpers"
)

func TestGroupsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*group.GroupsDataSourceModel)(nil)
	schema := group.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
