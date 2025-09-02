// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy_test

import (
	"context"
	"testing"

	"github.com/ServalHQ/terraform-provider-serval/internal/services/access_policy"
	"github.com/ServalHQ/terraform-provider-serval/internal/test_helpers"
)

func TestAccessPolicyDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*access_policy.AccessPolicyDataSourceModel)(nil)
	schema := access_policy.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
