resource "serval_app_instance" "example_app_instance" {
  instance_id = "instanceId"
  name = "name"
  team_id = "teamId"
  access_requests_enabled = true
  custom_service_id = "customServiceId"
  default_access_policy_id = "defaultAccessPolicyId"
  service = "service"
}
