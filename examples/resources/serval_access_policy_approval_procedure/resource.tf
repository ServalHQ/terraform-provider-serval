resource "serval_access_policy_approval_procedure" "example_access_policy_approval_procedure" {
  access_policy_id = "access_policy_id"
  steps = [{
    id = "id"
    allow_self_approval = true
    serval_group_ids = ["string"]
    specific_user_ids = ["string"]
    step_type = "APPROVAL_PROCEDURE_STEP_TYPE_UNSPECIFIED"
  }]
}
