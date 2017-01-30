package cmd

import (
	boshui "github.com/cloudfoundry/bosh-cli/ui"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"

	lintman "github.com/cppforlife/bosh-lint/manifest"
)

type LintManifestCmd struct {
	ui boshui.UI
}

func NewLintManifestCmd(ui boshui.UI) LintManifestCmd {
	return LintManifestCmd{ui}
}

func (c LintManifestCmd) Run(opts LintManifestOpts) error {
	lintableManifest, err := lintman.NewLintableManifest(opts.Args.File.Bytes)
	if err != nil {
		return err
	}

	descriptions, suggestions, err := lintableManifest.Lint()
	if err != nil {
		return err
	}

	ChecksTable{descriptions, suggestions, c.ui}.Print(opts.Verbose)

	if len(suggestions) > 0 {
		return bosherr.Errorf("Multiple suggestions found")
	}

	return nil
}
