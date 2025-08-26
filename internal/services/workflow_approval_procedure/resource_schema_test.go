// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflow_approval_procedure_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/serval-terraform/internal/services/workflow_approval_procedure"
	"github.com/stainless-sdks/serval-terraform/internal/test_helpers"
)

func TestWorkflowApprovalProcedureModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*workflow_approval_procedure.WorkflowApprovalProcedureModel)(nil)
	schema := workflow_approval_procedure.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
