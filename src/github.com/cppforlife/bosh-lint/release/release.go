package release

import (
	"fmt"

	boshrel "github.com/cloudfoundry/bosh-cli/release"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"

	check "github.com/cppforlife/bosh-lint/check"
)

type LintableRelease struct {
	release boshrel.Release
}

func NewLintableRelease(release boshrel.Release) LintableRelease {
	return LintableRelease{release}
}

func (r LintableRelease) Lint() ([]check.Description, []check.Suggestion, error) {
	var descriptions []check.Description
	var suggestions []check.Suggestion
	var errs []error

	for _, check := range r.collectChecks() {
		descriptions = append(descriptions, check.Description())

		sugs, err := check.Check()
		if err != nil {
			errs = append(errs, err)
		}

		suggestions = append(suggestions, sugs...)
	}

	if len(errs) > 0 {
		return descriptions, suggestions, bosherr.NewMultiError(errs...)
	}

	return descriptions, suggestions, nil
}

func (r LintableRelease) collectChecks() []check.Check {
	var checks []check.Check

	ctx := check.NewRootContext("Release")

	checks = append(checks, NewReleaseName(ctx, r.release.Name()))
	checks = append(checks, NewMissingLicense(ctx, r.release.License()))
	checks = append(checks, NewMissingJobs(ctx, r.release))
	checks = append(checks, NewUnusedPackages(ctx, r.release))

	for _, job := range r.release.Jobs() {
		ctx := ctx.Nested(fmt.Sprintf("Job '%s'", job.Name()))

		checks = append(checks, check.NewDashedName(ctx, job.Name()))
		checks = append(checks, NewJobPropertiesSyslogDaemonConfig(ctx, job))
		checks = append(checks, NewJobPropertiesCertificate(ctx, job))
		checks = append(checks, NewJobTemplatesCtl(ctx, job))

		for propName, propDef := range job.Properties {
			ctx := ctx.Nested(fmt.Sprintf("Property '%s'", propName))

			checks = append(checks, NewJobProperty(ctx, propName, propDef))
			checks = append(checks, NewJobPropertySecret(ctx, propName, propDef))
			checks = append(checks, NewJobPropertySkipSSL(ctx, propName))
			checks = append(checks, NewJobPropertyDeprecated(ctx, propDef))
			checks = append(checks, NewJobPropertyNamespace(ctx, propName, job))
			checks = append(checks, NewJobPropertyDebugAddr(ctx, propName, propDef))
			checks = append(checks, check.NewTodo(ctx, propDef.Description))
		}
	}

	for _, pkg := range r.release.Packages() {
		ctx := ctx.Nested(fmt.Sprintf("Package '%s'", pkg.Name()))

		checks = append(checks, NewPackageName(ctx, pkg.Name()))
	}

	return checks
}
