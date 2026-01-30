resource "serval_workflow" "example_workflow" {
  content = "content"
  name = "name"
  team_id = "teamId"
  type = "WORKFLOW_TYPE_UNSPECIFIED"
  description = "description"
  execution_scope = "WORKFLOW_EXECUTION_SCOPE_UNSPECIFIED"
  is_published = true
  is_temporary = true
  parameters = "parameters"
  require_form_confirmation = true
}
