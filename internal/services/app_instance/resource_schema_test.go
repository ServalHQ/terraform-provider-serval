// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_instance_test

import (
	"context"
	"testing"

	"github.com/ServalHQ/terraform-provider-serval/internal/services/app_instance"
	"github.com/ServalHQ/terraform-provider-serval/internal/test_helpers"
)

func TestAppInstanceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*app_instance.AppInstanceModel)(nil)
	schema := app_instance.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
