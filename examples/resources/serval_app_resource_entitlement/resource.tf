resource "serval_app_resource_entitlement" "example_app_resource_entitlement" {
  access_policy_id = "accessPolicyId"
  description = "description"
  linked_entitlement_ids = ["string"]
  name = "name"
  provisioning_method = "provisioningMethod"
  requests_enabled = true
  resource_id = "resourceId"
}
