resource "serval_workflow" "example_workflow" {
  content = "content"
  name = "name"
  team_id = "teamId"
  description = "description"
  execution_scope = "WORKFLOW_EXECUTION_SCOPE_UNSPECIFIED"
  is_published = true
  require_form_confirmation = true
}
