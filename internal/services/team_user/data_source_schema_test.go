// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package team_user_test

import (
	"context"
	"testing"

	"github.com/ServalHQ/terraform-provider-serval/internal/services/team_user"
	"github.com/ServalHQ/terraform-provider-serval/internal/test_helpers"
)

func TestTeamUserDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*team_user.TeamUserDataSourceModel)(nil)
	schema := team_user.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
