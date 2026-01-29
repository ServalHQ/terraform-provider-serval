resource "serval_workflow" "example_workflow" {
  content = "content"
  description = "description"
  execution_scope = "WORKFLOW_EXECUTION_SCOPE_UNSPECIFIED"
  is_published = true
  is_temporary = true
  name = "name"
  parameters = "parameters"
  require_form_confirmation = true
  team_id = "teamId"
  type = "WORKFLOW_TYPE_UNSPECIFIED"
}
