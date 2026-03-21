// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package approval_delegation_test

import (
	"context"
	"testing"

	"github.com/ServalHQ/terraform-provider-serval/internal/services/approval_delegation"
	"github.com/ServalHQ/terraform-provider-serval/internal/test_helpers"
)

func TestApprovalDelegationDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*approval_delegation.ApprovalDelegationDataSourceModel)(nil)
	schema := approval_delegation.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
