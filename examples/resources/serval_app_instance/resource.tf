resource "serval_app_instance" "example_app_instance" {
  custom_service_id = "customServiceId"
  access_requests_enabled = true
  default_access_policy_id = "defaultAccessPolicyId"
  instance_id = "instanceId"
  name = "name"
  team_id = "teamId"
}
