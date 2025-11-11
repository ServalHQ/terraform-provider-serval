# Query by ID
data "serval_user" "example_user_by_id" {
  id = "id"
}

# Query by email
data "serval_user" "example_user_by_email" {
  email = "user@example.com"
}
