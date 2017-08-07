package cmd

import (
	"fmt"

	boshcmd "github.com/cloudfoundry/bosh-cli/cmd"
	boshrel "github.com/cloudfoundry/bosh-cli/release"
	boshreldir "github.com/cloudfoundry/bosh-cli/releasedir"
	boshui "github.com/cloudfoundry/bosh-cli/ui"
)

type Cmd struct {
	BoshOpts BoshOpts
	Opts     interface{}

	deps BasicDeps
}

func NewCmd(boshOpts BoshOpts, opts interface{}, deps BasicDeps) Cmd {
	return Cmd{boshOpts, opts, deps}
}

type cmdConveniencePanic struct {
	Err error
}

func (c Cmd) Execute() (cmdErr error) {
	// Catch convenience panics from panicIfErr
	defer func() {
		if r := recover(); r != nil {
			if cp, ok := r.(cmdConveniencePanic); ok {
				cmdErr = cp.Err
			} else {
				panic(r)
			}
		}
	}()

	c.configureUI()
	c.configureFS()

	deps := c.deps

	switch opts := c.Opts.(type) {
	case *LintReleaseOpts:
		_, relDirProv := c.releaseProviders()

		releaseDirFactory := func(dir boshcmd.DirOrCWDArg) (boshrel.Reader, boshreldir.ReleaseDir) {
			releaseReader := relDirProv.NewReleaseReader(dir.Path)
			releaseDir := relDirProv.NewFSReleaseDir(dir.Path)
			return releaseReader, releaseDir
		}

		return NewLintReleaseCmd(releaseDirFactory, deps.FS, deps.UI).Run(*opts)

	case *LintCPIConfigOpts:
		return fmt.Errorf("Not implemented yet")

	case *LintCloudConfigOpts:
		return fmt.Errorf("Not implemented yet")

	case *LintRuntimeConfigOpts:
		return fmt.Errorf("Not implemented yet")

	case *LintManifestOpts:
		return NewLintManifestCmd(deps.UI).Run(*opts)

	case *DebugTaskOpts:
		return NewDebugTaskCmd(deps.UI).Run(*opts)

	case *VisualizeEventsOpts:
		return NewVisualizeEventsCmd(deps.CmdRunner, deps.UI, deps.Logger).Run()

	case *MessageOpts:
		deps.UI.PrintBlock(opts.Message)
		return nil

	default:
		return fmt.Errorf("Unhandled command: %#v", c.Opts)
	}
}

func (c Cmd) configureUI() {
	c.deps.UI.EnableTTY(c.BoshOpts.TTYOpt)

	if !c.BoshOpts.NoColorOpt {
		c.deps.UI.EnableColor()
	}

	if c.BoshOpts.JSONOpt {
		c.deps.UI.EnableJSON()
	}

	c.deps.UI.EnableNonInteractive()
}

func (c Cmd) configureFS() {
	tmpDirPath, err := c.deps.FS.ExpandPath("~/.bosh/tmp")
	c.panicIfErr(err)

	err = c.deps.FS.ChangeTempRoot(tmpDirPath)
	c.panicIfErr(err)
}

func (c Cmd) releaseProviders() (boshrel.Provider, boshreldir.Provider) {
	indexReporter := boshui.NewIndexReporter(c.deps.UI)
	blobsReporter := boshui.NewBlobsReporter(c.deps.UI)
	releaseIndexReporter := boshui.NewReleaseIndexReporter(c.deps.UI)

	releaseProvider := boshrel.NewProvider(
		c.deps.CmdRunner, c.deps.Compressor, c.deps.SHA1Calc, c.deps.FS, c.deps.Logger)

	releaseDirProvider := boshreldir.NewProvider(
		indexReporter, releaseIndexReporter, blobsReporter, releaseProvider,
		c.deps.SHA1Calc, c.deps.CmdRunner, c.deps.UUIDGen, c.deps.Time, c.deps.FS, c.deps.Logger)

	return releaseProvider, releaseDirProvider
}

func (c Cmd) releaseDir(dir boshcmd.DirOrCWDArg) boshreldir.ReleaseDir {
	_, relDirProv := c.releaseProviders()
	return relDirProv.NewFSReleaseDir(dir.Path)
}

func (c Cmd) panicIfErr(err error) {
	if err != nil {
		panic(cmdConveniencePanic{err})
	}
}
