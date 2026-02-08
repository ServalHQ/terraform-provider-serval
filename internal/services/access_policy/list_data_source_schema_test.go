// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy_test

import (
	"context"
	"testing"

	"github.com/ServalHQ/terraform-provider-serval/internal/services/access_policy"
	"github.com/ServalHQ/terraform-provider-serval/internal/test_helpers"
)

func TestAccessPoliciesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*access_policy.AccessPoliciesDataSourceModel)(nil)
	schema := access_policy.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
