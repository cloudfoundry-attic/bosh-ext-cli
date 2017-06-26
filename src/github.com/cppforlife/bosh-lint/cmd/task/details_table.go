package task

import (
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
