package cmd

import (
	boshui "github.com/cloudfoundry/bosh-cli/ui"
	boshtbl "github.com/cloudfoundry/bosh-cli/ui/table"

	check "github.com/bosh-tools/bosh-ext-cli/lint/check"
)

type ChecksTable struct {
	Descriptions []check.Description
	Suggestions  []check.Suggestion

	UI boshui.UI
}

func (t ChecksTable) Print(verbose int) {
	switch verbose {
	case 1:
		t.printChecks()
	case 2:
		t.printChecksWithContext()
	default:
		// do nothing
	}

	t.printSuggestions()
}

func (t ChecksTable) printChecks() {
	uniqPurposes := map[string]struct{}{}

	for _, desc := range t.Descriptions {
		uniqPurposes[desc.Purpose()] = struct{}{}
	}

	table := boshtbl.Table{
		Content: "checks",

		SortBy: []boshtbl.ColumnSort{
			{Column: 0, Asc: true},
		},

		Header: []boshtbl.Header{
			boshtbl.NewHeader("Purpose"),
		},
	}

	for purpose, _ := range uniqPurposes {
		table.Rows = append(table.Rows, []boshtbl.Value{
			boshtbl.NewValueString(purpose),
		})
	}

	t.UI.PrintTable(table)
}

func (t ChecksTable) printChecksWithContext() {
	table := boshtbl.Table{
		Content: "checks",

		SortBy: []boshtbl.ColumnSort{
			{Column: 0, Asc: true},
			{Column: 1, Asc: true},
		},

		Header: []boshtbl.Header{
			boshtbl.NewHeader("Purpose"),
			boshtbl.NewHeader("Context"),
		},
	}

	for _, desc := range t.Descriptions {
		table.Rows = append(table.Rows, []boshtbl.Value{
			boshtbl.NewValueString(desc.Purpose()),
			boshtbl.NewValueString(desc.Context().String()),
		})
	}

	t.UI.PrintTable(table)
}

func (t ChecksTable) printSuggestions() {
	table := boshtbl.Table{
		Content: "suggestions",

		SortBy: []boshtbl.ColumnSort{
			{Column: 0, Asc: true},
			{Column: 1, Asc: true},
		},

		Header: []boshtbl.Header{
			boshtbl.NewHeader("Purpose"),
			boshtbl.NewHeader("Context"),
			boshtbl.NewHeader("Resolution"),
		},
	}

	for _, sug := range t.Suggestions {
		table.Rows = append(table.Rows, []boshtbl.Value{
			boshtbl.NewValueString(sug.Problem()),
			boshtbl.NewValueString(sug.Context().String()),
			boshtbl.NewValueString(sug.Resolution()),
		})
	}

	t.UI.PrintTable(table)
}
