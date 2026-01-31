// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflow_approval_procedure_test

import (
	"context"
	"testing"

	"github.com/ServalHQ/terraform-provider-serval/internal/services/workflow_approval_procedure"
	"github.com/ServalHQ/terraform-provider-serval/internal/test_helpers"
)

func TestWorkflowApprovalProcedureModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*workflow_approval_procedure.WorkflowApprovalProcedureModel)(nil)
	schema := workflow_approval_procedure.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	// WORKAROUND: steps[].id is intentionally Optional+Computed (not just Computed) to fix
	// OpenTofu import config generation. The model tag says "computed" but schema is "computed_optional".
	errs.Ignore(t, ".@WorkflowApprovalProcedureModel.steps.[].@WorkflowApprovalProcedureStepsModel.id")
	errs.Report(t)
}
