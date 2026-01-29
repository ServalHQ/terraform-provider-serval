resource "serval_access_policy" "example_access_policy" {
  name = "name"
  team_id = "teamId"
  description = "description"
  max_access_minutes = 0
  recommended_access_minutes = 0
  require_business_justification = true
}
