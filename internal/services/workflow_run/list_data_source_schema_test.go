// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflow_run_test

import (
	"context"
	"testing"

	"github.com/ServalHQ/terraform-provider-serval/internal/services/workflow_run"
	"github.com/ServalHQ/terraform-provider-serval/internal/test_helpers"
)

func TestWorkflowRunsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*workflow_run.WorkflowRunsDataSourceModel)(nil)
	schema := workflow_run.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
