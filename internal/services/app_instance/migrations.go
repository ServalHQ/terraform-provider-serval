// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_instance

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*AppInstanceResource)(nil)

func (r *AppInstanceResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
