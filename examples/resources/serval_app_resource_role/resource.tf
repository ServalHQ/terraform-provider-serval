resource "serval_app_resource_role" "example_app_resource_role" {
  name = "name"
  provisioning_method = {
    builtin_workflow = {

    }
  }
  resource_id = "resourceId"
  access_policy_id = "accessPolicyId"
  description = "description"
  external_data = "externalData"
  external_id = "externalId"
  requests_enabled = true
}
