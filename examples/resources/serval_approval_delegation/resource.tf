resource "serval_approval_delegation" "example_approval_delegation" {
  delegates = [{
    id = "id"
    type = "APPROVAL_DELEGATE_TYPE_UNSPECIFIED"
  }]
  delegator_user_id = "delegatorUserId"
  description = "description"
  priority = 0
}
