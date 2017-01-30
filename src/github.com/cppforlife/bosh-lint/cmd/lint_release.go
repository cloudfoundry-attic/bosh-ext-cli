package cmd

import (
	boshcmd "github.com/cloudfoundry/bosh-cli/cmd"
	boshrel "github.com/cloudfoundry/bosh-cli/release"
	boshreldir "github.com/cloudfoundry/bosh-cli/releasedir"
	boshui "github.com/cloudfoundry/bosh-cli/ui"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshsys "github.com/cloudfoundry/bosh-utils/system"

	lintrel "github.com/cppforlife/bosh-lint/release"
)

type LintReleaseCmd struct {
	releaseDirFactory func(boshcmd.DirOrCWDArg) (boshrel.Reader, boshreldir.ReleaseDir)

	fs boshsys.FileSystem
	ui boshui.UI
}

func NewLintReleaseCmd(
	releaseDirFactory func(boshcmd.DirOrCWDArg) (boshrel.Reader, boshreldir.ReleaseDir),
	fs boshsys.FileSystem,
	ui boshui.UI,
) LintReleaseCmd {
	return LintReleaseCmd{
		releaseDirFactory: releaseDirFactory,

		fs: fs,
		ui: ui,
	}
}

func (c LintReleaseCmd) Run(opts LintReleaseOpts) error {
	release, err := c.release(opts)
	if err != nil {
		return err
	}

	lintableRelease := lintrel.NewLintableRelease(release)

	descriptions, suggestions, err := lintableRelease.Lint()
	if err != nil {
		return err
	}

	ChecksTable{descriptions, suggestions, c.ui}.Print(opts.Verbose)

	if len(suggestions) > 0 {
		return bosherr.Errorf("Multiple suggestions found")
	}

	return nil
}

func (c LintReleaseCmd) release(opts LintReleaseOpts) (boshrel.Release, error) {
	if c.releaseDirFactory == nil {
		return nil, bosherr.Errorf("releaseDirFactory must be specified")
	}

	_, releaseDir := c.releaseDirFactory(opts.Directory)

	name, err := releaseDir.DefaultName()
	if err != nil {
		return nil, err
	}

	version, err := releaseDir.NextDevVersion(name, false)
	if err != nil {
		return nil, err
	}

	return releaseDir.BuildRelease(name, version, true)
}
