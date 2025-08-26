// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/serval-terraform/internal/services/access_policy"
	"github.com/stainless-sdks/serval-terraform/internal/test_helpers"
)

func TestAccessPolicyModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*access_policy.AccessPolicyModel)(nil)
	schema := access_policy.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
