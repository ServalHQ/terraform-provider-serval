// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy_approval_procedure_test

import (
	"context"
	"testing"

	"github.com/ServalHQ/terraform-provider-serval/internal/services/access_policy_approval_procedure"
	"github.com/ServalHQ/terraform-provider-serval/internal/test_helpers"
)

func TestAccessPolicyApprovalProcedureModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*access_policy_approval_procedure.AccessPolicyApprovalProcedureModel)(nil)
	schema := access_policy_approval_procedure.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
