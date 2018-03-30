package release

import (
	"fmt"

	boshrel "github.com/cloudfoundry/bosh-cli/release"

	check "github.com/bosh-tools/bosh-ext-cli/lint/check"
)

type UnusedPackages struct {
	context check.Context
	release boshrel.Release
	check.Config
}

func NewUnusedPackages(context check.Context, release boshrel.Release, config check.Config) UnusedPackages {
	return UnusedPackages{context, release, config}
}

func (c UnusedPackages) Description() check.Description {
	return check.Description{
		Context_: c.context,
		Purpose_: "if there are any unused packages",
	}
}

func (c UnusedPackages) Check() ([]check.Suggestion, error) {
	var sugs []check.Suggestion

	allPkgNames := map[string]struct{}{}

	for _, pkg := range c.release.Packages() {
		allPkgNames[pkg.Name()] = struct{}{}
	}

	for _, pkg := range c.release.Packages() {
		for _, dep := range pkg.Dependencies {
			delete(allPkgNames, dep.Name())
		}
	}

	for _, job := range c.release.Jobs() {
		for _, pkg := range job.Packages {
			delete(allPkgNames, pkg.Name())
		}
	}

	for unusedName, _ := range allPkgNames {
		sugs = append(sugs, check.Simple{
			Context_:    c.context,
			Problem_:    "Unused packages",
			Resolution_: fmt.Sprintf("Delete package '%s'", unusedName),
		})
	}

	return sugs, nil
}
