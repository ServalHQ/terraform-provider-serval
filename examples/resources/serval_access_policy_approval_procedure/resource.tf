resource "serval_access_policy_approval_procedure" "example_access_policy_approval_procedure" {
  access_policy_id = "access_policy_id"
  steps = [
    {
      allow_self_approval = false
      approvers = [
        {
          group = {
            group_id = "d4f5a926-1a4b-4c3d-9e8f-7b6a5c4d3e2f"
          }
          notify = true
        },
        {
          user = {
            user_id = "a1b2c3d4-e5f6-7890-abcd-ef1234567890"
          }
          notify = true
        },
        {
          manager = "{}"
          notify = true
        },
        {
          app_owner = "{}"
          notify = true
        },
      ]
    },
    {
      custom_workflow = {
        workflow_id = "b2c3d4e5-f6a7-8901-bcde-f12345678901"
      }
    },
  ]
}
