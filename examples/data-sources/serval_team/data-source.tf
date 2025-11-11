data "serval_team" "example_team_by_id" {
  id = "id"
}

data "serval_team" "example_team_by_name" {
  name = "Engineering"
}

data "serval_team" "example_team_by_prefix" {
  prefix = "eng"
}
