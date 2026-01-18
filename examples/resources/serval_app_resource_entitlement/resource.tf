resource "serval_app_resource_entitlement" "example_app_resource_entitlement" {
  access_policy_id = "accessPolicyId"
  description = "description"
  external_data = "externalData"
  external_id = "externalId"
  name = "name"
  provisioning_method = {
    builtin_workflow = {

    }
    custom_workflow = {
      deprovision_workflow_id = "deprovisionWorkflowId"
      provision_workflow_id = "provisionWorkflowId"
    }
    linked_entitlements = {
      linked_entitlement_ids = ["string"]
    }
    manual = {
      assignees = [{
        assignee_id = "assigneeId"
        assignee_type = "MANUAL_PROVISIONING_ASSIGNEE_TYPE_UNSPECIFIED"
      }]
    }
  }
  requests_enabled = true
  resource_id = "resourceId"
}
