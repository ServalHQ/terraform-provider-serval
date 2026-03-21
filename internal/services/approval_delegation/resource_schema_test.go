// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package approval_delegation_test

import (
	"context"
	"testing"

	"github.com/ServalHQ/terraform-provider-serval/internal/services/approval_delegation"
	"github.com/ServalHQ/terraform-provider-serval/internal/test_helpers"
)

func TestApprovalDelegationModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*approval_delegation.ApprovalDelegationModel)(nil)
	schema := approval_delegation.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
