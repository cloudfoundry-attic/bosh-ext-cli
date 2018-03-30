package cmd

import (
	boshui "github.com/cloudfoundry/bosh-cli/ui"

	lintcmdtask "github.com/bosh-tools/bosh-ext-cli/cmd/task"
	linttask "github.com/bosh-tools/bosh-ext-cli/task"
)

type DebugTaskCmd struct {
	ui boshui.UI
}

func NewDebugTaskCmd(ui boshui.UI) DebugTaskCmd {
	return DebugTaskCmd{ui}
}

func (c DebugTaskCmd) Run(opts DebugTaskOpts) error {
	task := linttask.NewTask(opts.Args.File.Bytes)

	details, err := task.Details()
	if err != nil {
		return err
	}

	switch {
	case opts.Lines:
		return c.showLines(details, opts)

	case opts.Actions:
		lintcmdtask.ActionsTable{details.Group(), opts.SortBy, c.ui}.Print()
		return nil

	default:
		lintcmdtask.DetailsTable{details, c.ui}.Print()
		lintcmdtask.GroupsTable{details.Groups(), opts.SortBy, c.ui}.Print()
		return nil
	}
}

func (c DebugTaskCmd) showLines(details linttask.Details, opts DebugTaskOpts) error {
	lineFilter := func(line *linttask.Line) bool {
		if opts.Errors && !line.IsError() {
			return false
		}
		if !opts.DB {
			if _, ok := line.Action().(*linttask.DBStatement); ok {
				return false
			}
		}
		return true
	}

	lintcmdtask.LinesTable{details.Lines, details.UnknownLines, lineFilter, c.ui}.Print()

	return nil
}
