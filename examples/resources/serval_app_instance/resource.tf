resource "serval_app_instance" "example_app_instance" {
  access_requests_enabled = true
  custom_service_id = "customServiceId"
  default_access_policy_id = "defaultAccessPolicyId"
  instance_id = "instanceId"
  name = "name"
  service = "service"
  team_id = "teamId"
}
