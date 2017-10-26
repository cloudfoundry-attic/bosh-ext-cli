package task

import (
	boshui "github.com/cloudfoundry/bosh-cli/ui"
	boshtbl "github.com/cloudfoundry/bosh-cli/ui/table"

	linttask "github.com/cppforlife/bosh-lint/task"
)

type GroupsTable struct {
	Groups []linttask.Group
	SortBy string
	UI     boshui.UI
}

func (t GroupsTable) Print() {
	table := boshtbl.Table{
		Content: "groups",

		Header: []boshtbl.Header{
			{Title: "Started at"},
			{Title: "Ended at"},
			{Title: "Duration"},
			{Title: "Name"},
		},

		SortBy: []boshtbl.ColumnSort{
			{Column: 0, Asc: true},
			{Column: 3, Asc: true},
		},
	}

	if t.SortBy == "duration" {
		table.SortBy = []boshtbl.ColumnSort{
			{Column: 2, Asc: false},
		}
	}

	for _, group := range t.Groups {
		table.Rows = append(table.Rows, []boshtbl.Value{
			boshtbl.NewValueTime(group.StartedAt()),
			boshtbl.NewValueTime(group.EndedAt()),
			NewValueDuration(group.Duration()),
			boshtbl.NewValueString(group.Name),
		})
	}

	t.UI.PrintTable(table)
}
