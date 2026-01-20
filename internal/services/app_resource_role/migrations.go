// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_resource_role

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*AppResourceRoleResource)(nil)

func (r *AppResourceRoleResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
