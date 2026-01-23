resource "serval_access_policy_approval_procedure" "example_access_policy_approval_procedure" {
  access_policy_id = "access_policy_id"
  steps = [{
    allow_self_approval = true
    custom_workflow_id = "customWorkflowId"
    serval_group_ids = ["string"]
    specific_user_ids = ["string"]
  }]
}
