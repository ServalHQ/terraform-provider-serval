package workflow

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
)

func TestUpdateBodyOmitsTeamID(t *testing.T) {
	state := WorkflowModel{
		ID:     types.StringValue("id-1"),
		TeamID: types.StringValue("team-1"),
		Name:   types.StringValue("old"),
	}
	plan := state
	plan.Name = types.StringValue("new")

	update, err := plan.MarshalJSONForUpdate(state)
	if err != nil {
		t.Fatal(err)
	}
	if gjson.GetBytes(update, "teamId").Exists() {
		t.Errorf("update body must not contain teamId: %s", update)
	}
	if got := gjson.GetBytes(update, "name").String(); got != "new" {
		t.Errorf("update body name = %q, want %q: %s", got, "new", update)
	}

	create, err := plan.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	if got := gjson.GetBytes(create, "teamId").String(); got != "team-1" {
		t.Errorf("create body teamId = %q, want %q: %s", got, "team-1", create)
	}
}
