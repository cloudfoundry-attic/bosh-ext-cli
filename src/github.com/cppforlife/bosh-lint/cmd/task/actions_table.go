package task

import (
	"fmt"

	boshui "github.com/cloudfoundry/bosh-cli/ui"
	boshtbl "github.com/cloudfoundry/bosh-cli/ui/table"

	linttask "github.com/cppforlife/bosh-lint/task"
)

type ActionsTable struct {
	Group linttask.Group
	UI    boshui.UI
}

func (t ActionsTable) Print() {
	table := boshtbl.Table{
		Content: "actions",

		Header: []string{"Started at", "Ended at", "Duration", "Content"},
	}

	for _, action := range t.Group.Actions() {
		table.Rows = append(table.Rows, []boshtbl.Value{
			boshtbl.NewValueTime(action.StartedAt()),
			boshtbl.NewValueTime(action.EndedAt()),
			boshtbl.NewValueString(fmt.Sprintf("%s", action.Duration())),
			boshtbl.NewValueString(action.ShortDescription()),
		})
	}

	t.UI.PrintTable(table)
}
