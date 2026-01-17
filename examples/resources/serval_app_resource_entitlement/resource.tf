resource "serval_app_resource_entitlement" "example_app_resource_entitlement" {
  access_policy_id = "accessPolicyId"
  description = "description"
  external_data = "externalData"
  external_id = "externalId"
  name = "name"
  provisioning_method = {
    builtin_workflow = {

    }
  }
  requests_enabled = true
  resource_id = "resourceId"
}
