resource "serval_access_policy_approval_procedure" "example_access_policy_approval_procedure" {
  access_policy_id = "access_policy_id"
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
