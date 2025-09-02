// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflow_test

import (
	"context"
	"testing"

	"github.com/ServalHQ/terraform-provider-serval/internal/services/workflow"
	"github.com/ServalHQ/terraform-provider-serval/internal/test_helpers"
)

func TestWorkflowDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*workflow.WorkflowDataSourceModel)(nil)
	schema := workflow.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
