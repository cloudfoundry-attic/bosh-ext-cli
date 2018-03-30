package task

import (
	boshui "github.com/cloudfoundry/bosh-cli/ui"
	boshtbl "github.com/cloudfoundry/bosh-cli/ui/table"

	linttask "github.com/bosh-tools/bosh-ext-cli/task"
)

type ActionsTable struct {
	Group  linttask.Group
	SortBy string
	UI     boshui.UI
}

func (t ActionsTable) Print() {
	table := boshtbl.Table{
		Content: "actions",

		Header: []string{"Started at", "Ended at", "Duration", "Group", "Content"},
	}

	if t.SortBy == "duration" {
		table.SortBy = []boshtbl.ColumnSort{
			{Column: 2, Asc: false},
		}
	}

	for _, action := range t.Group.Actions() {
		table.Rows = append(table.Rows, []boshtbl.Value{
			boshtbl.NewValueTime(action.StartedAt()),
			boshtbl.NewValueTime(action.EndedAt()),
			NewValueDuration(action.Duration()),
			boshtbl.NewValueString(action.Group()),
			boshtbl.NewValueString(action.ShortDescription()),
		})
	}

	t.UI.PrintTable(table)
}
