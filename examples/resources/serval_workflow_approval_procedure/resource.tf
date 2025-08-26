resource "serval_workflow_approval_procedure" "example_workflow_approval_procedure" {
  workflow_id = "workflow_id"
  steps = [{
    id = "id"
    allow_self_approval = true
    serval_group_ids = ["string"]
    specific_user_ids = ["string"]
    step_type = "APPROVAL_PROCEDURE_STEP_TYPE_UNSPECIFIED"
  }]
}
