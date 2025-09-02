// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflow_approval_procedure_test

import (
	"context"
	"testing"

	"github.com/ServalHQ/terraform-provider-serval/internal/services/workflow_approval_procedure"
	"github.com/ServalHQ/terraform-provider-serval/internal/test_helpers"
)

func TestWorkflowApprovalProcedureDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*workflow_approval_procedure.WorkflowApprovalProcedureDataSourceModel)(nil)
	schema := workflow_approval_procedure.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
