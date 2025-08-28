resource "serval_app_instance" "example_app_instance" {
  access_requests_enabled = true
  default_access_policy_id = "defaultAccessPolicyId"
  instance_id = "instanceId"
  name = "name"
  service = "service"
  team_id = "teamId"
}
