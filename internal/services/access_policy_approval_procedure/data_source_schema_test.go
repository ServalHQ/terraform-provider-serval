// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy_approval_procedure_test

import (
	"context"
	"testing"

	"github.com/ServalHQ/terraform-provider-serval/internal/services/access_policy_approval_procedure"
	"github.com/ServalHQ/terraform-provider-serval/internal/test_helpers"
)

func TestAccessPolicyApprovalProcedureDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*access_policy_approval_procedure.AccessPolicyApprovalProcedureDataSourceModel)(nil)
	schema := access_policy_approval_procedure.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
