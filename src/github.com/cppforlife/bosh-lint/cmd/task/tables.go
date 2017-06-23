package task

import (
	"fmt"

	boshui "github.com/cloudfoundry/bosh-cli/ui"
	boshtbl "github.com/cloudfoundry/bosh-cli/ui/table"

	linttask "github.com/cppforlife/bosh-lint/task"
)

type DetailsTable struct {
	Details linttask.Details
	UI      boshui.UI
}

func (t DetailsTable) Print() {
	table := boshtbl.Table{
		Content: "details",

		Header: []string{"Director Version", "Task ID", "Lines", "Unknown Lines"},
	}

	table.Rows = append(table.Rows, []boshtbl.Value{
		boshtbl.NewValueString(t.Details.DirectorVersion),
		boshtbl.NewValueString(t.Details.TaskID),
		boshtbl.NewValueInt(len(t.Details.Lines)),
		boshtbl.NewValueInt(len(t.Details.UnknownLines)),
	})

	t.UI.PrintTable(table)
}

type GroupsTable struct {
	Groups []linttask.Group
	UI     boshui.UI
}

func (t GroupsTable) Print() {
	table := boshtbl.Table{
		Content: "groups",

		Header: []string{"Name", "Started at", "Ended at", "Duration"},

		SortBy: []boshtbl.ColumnSort{
			{Column: 1, Asc: true},
			{Column: 0, Asc: true},
		},
	}

	for _, group := range t.Groups {
		table.Rows = append(table.Rows, []boshtbl.Value{
			boshtbl.NewValueString(group.Name),
			boshtbl.NewValueTime(group.StartedAt()),
			boshtbl.NewValueTime(group.EndedAt()),
			boshtbl.NewValueString(fmt.Sprintf("%s", group.Duration())),
		})
	}

	t.UI.PrintTable(table)
}
