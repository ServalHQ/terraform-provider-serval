// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package team_user

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*TeamUserResource)(nil)

func (r *TeamUserResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
