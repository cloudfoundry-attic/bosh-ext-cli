package cmd

import (
	"code.cloudfoundry.org/clock"
	bicrypto "github.com/cloudfoundry/bosh-cli/crypto"
	boshui "github.com/cloudfoundry/bosh-cli/ui"
	boshcrypto "github.com/cloudfoundry/bosh-utils/crypto"
	boshcmd "github.com/cloudfoundry/bosh-utils/fileutil"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	boshuuid "github.com/cloudfoundry/bosh-utils/uuid"
)

type BasicDeps struct {
	FS     boshsys.FileSystem
	UI     *boshui.ConfUI
	Logger boshlog.Logger

	UUIDGen    boshuuid.Generator
	CmdRunner  boshsys.CmdRunner
	Compressor boshcmd.Compressor

	DigestCreationAlgorithms []boshcrypto.Algorithm
	DigestCalculator         bicrypto.DigestCalculator

	Time clock.Clock
}

func NewBasicDeps(ui *boshui.ConfUI, logger boshlog.Logger) BasicDeps {
	return NewBasicDepsWithFS(ui, boshsys.NewOsFileSystemWithStrictTempRoot(logger), logger)
}

func NewBasicDepsWithFS(ui *boshui.ConfUI, fs boshsys.FileSystem, logger boshlog.Logger) BasicDeps {
	cmdRunner := boshsys.NewExecCmdRunner(logger)
	digestCreationAlgorithms := []boshcrypto.Algorithm{boshcrypto.DigestAlgorithmSHA1}
	digestCalculator := bicrypto.NewDigestCalculator(fs, digestCreationAlgorithms)

	return BasicDeps{
		FS:     fs,
		UI:     ui,
		Logger: logger,

		UUIDGen:    boshuuid.NewGenerator(),
		CmdRunner:  cmdRunner,
		Compressor: boshcmd.NewTarballCompressor(cmdRunner, fs),

		DigestCreationAlgorithms: digestCreationAlgorithms,
		DigestCalculator:         digestCalculator,

		Time: clock.NewClock(),
	}
}
