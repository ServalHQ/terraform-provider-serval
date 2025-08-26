resource "serval_access_policy" "example_access_policy" {
  description = "description"
  max_access_minutes = 0
  name = "name"
  require_business_justification = true
  team_id = "teamId"
}
