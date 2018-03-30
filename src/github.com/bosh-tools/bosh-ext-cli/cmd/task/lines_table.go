package task

import (
	boshui "github.com/cloudfoundry/bosh-cli/ui"
	boshtbl "github.com/cloudfoundry/bosh-cli/ui/table"

	linttask "github.com/bosh-tools/bosh-ext-cli/task"
)

type LinesTable struct {
	Lines        []*linttask.Line
	UnknownLines []linttask.UnknownLine
	LineFilter   func(*linttask.Line) bool
	UI           boshui.UI
}

func (t LinesTable) Print() {
	table := boshtbl.Table{
		Content: "lines",

		Header: []boshtbl.Header{
			boshtbl.NewHeader("Level"),
			boshtbl.NewHeader("Num"),
			boshtbl.NewHeader("Time"),
			boshtbl.NewHeader("Group"),
			boshtbl.NewHeader("Content"),
		},

		SortBy: []boshtbl.ColumnSort{
			{Column: 1, Asc: true},
		},
	}

	for _, line := range t.Lines {
		if t.LineFilter(line) {
			table.Rows = append(table.Rows, []boshtbl.Value{
				boshtbl.NewValueString(line.Level),
				boshtbl.NewValueInt(line.Number),
				boshtbl.NewValueTime(line.Time),
				boshtbl.NewValueString(line.Group),
				boshtbl.NewValueString(line.Action().ShortDescription()),
			})
		}
	}

	for _, line := range t.UnknownLines {
		table.Rows = append(table.Rows, []boshtbl.Value{
			boshtbl.NewValueString("<unknown>"),
			boshtbl.NewValueInt(line.Number),
			boshtbl.ValueNone{},
			boshtbl.NewValueString(""),
			boshtbl.NewValueString(line.Content),
		})
	}

	t.UI.PrintTable(table)
}
