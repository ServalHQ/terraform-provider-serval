resource "serval_workflow_approval_procedure" "example_workflow_approval_procedure" {
  workflow_id = "workflow_id"
  steps = [{
    allow_self_approval = true
    approvers = [{
      app_owner = {

      }
      notify = true
    }]
    custom_workflow = {
      workflow_id = "workflowId"
    }
  }]
}
