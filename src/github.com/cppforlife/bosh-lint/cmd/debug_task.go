package cmd

import (
	boshui "github.com/cloudfoundry/bosh-cli/ui"

	lintcmdtask "github.com/cppforlife/bosh-lint/cmd/task"
	linttask "github.com/cppforlife/bosh-lint/task"
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

	lintcmdtask.DetailsTable{details, c.ui}.Print()
	lintcmdtask.GroupsTable{details.Groups(), c.ui}.Print()

	if opts.Lines {
		return c.showLines(details, opts)
	}

	return c.showActions(details, opts)
}

func (c DebugTaskCmd) showActions(details linttask.Details, opts DebugTaskOpts) error {
	group, err := details.FindGroup(opts.Group)
	if err != nil {
		return err
	}

	lintcmdtask.ActionsTable{group, c.ui}.Print()

	return nil
}

func (c DebugTaskCmd) showLines(details linttask.Details, opts DebugTaskOpts) error {
	lineFilter := func(line *linttask.Line) bool {
		if opts.Errors && !line.IsError() {
			return false
		}
		if len(opts.Group) > 0 && line.Group != opts.Group {
			return false
		}
		if !opts.DB {
			if _, ok := line.Action().(linttask.DBStatement); ok {
				return false
			}
		}
		return true
	}

	lintcmdtask.LinesTable{details.Lines, details.UnknownLines, lineFilter, c.ui}.Print()

	return nil
}
