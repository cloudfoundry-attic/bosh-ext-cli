package release

import (
	"fmt"

	boshrel "github.com/cloudfoundry/bosh-cli/release"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"

	check "github.com/cppforlife/bosh-lint/check"
)

type ReleaseConfig struct {
	ReleaseName    check.CheckConfig `yaml:"release_name"`
	MissingLicense check.CheckConfig `yaml:"missing_license"`
	MissingJobs    check.CheckConfig `yaml:"missing_jobs"`
	UnusedPackages check.CheckConfig `yaml:"unused_packages"`
	PackageName    check.CheckConfig `yaml:"package_name"`

	DashedName                      check.CheckConfig `yaml:"dashed_name"`
	JobPropertiesSyslogDaemonConfig check.CheckConfig `yaml:"job_properties_syslog_daemon_config"`
	JobPropertiesCertificate        check.CheckConfig `yaml:"job_properties_certificate"`
	JobTemplatesCtl                 check.CheckConfig `yaml:"job_templates_ctl"`

	JobProperty           check.CheckConfig          `yaml:"job_property"`
	JobPropertySecret     JobPropertySecretConfig    `yaml:"job_property_secret"`
	JobPropertySkipSSL    check.CheckConfig          `yaml:"job_property_skip_ssl"`
	JobPropertyDeprecated check.CheckConfig          `yaml:"job_property_deprecated"`
	JobPropertyNamespace  check.CheckConfig          `yaml:"job_property_namespace"`
	JobPropertyDebugAddr  JobPropertyDebugAddrConfig `yaml:"job_property_debug_addr"`
	Todo                  check.CheckConfig          `yaml:"todo"`
}

var DefaultReleaseConfig = ReleaseConfig{
	JobPropertySecret:    DefaultJobPropertySecretConfig,
	JobPropertyDebugAddr: DefaultJobPropertyDebugAddrConfig,
}

type LintableRelease struct {
	release boshrel.Release
	config  ReleaseConfig
}

func NewLintableRelease(release boshrel.Release, config ReleaseConfig) LintableRelease {
	return LintableRelease{release, config}
}

func (r LintableRelease) Lint() ([]check.Description, []check.Suggestion, error) {
	var descriptions []check.Description
	var suggestions []check.Suggestion
	var errs []error

	for _, check := range r.collectChecks() {
		if check.IsEnabled() {
			descriptions = append(descriptions, check.Description())

			sugs, err := check.Check()
			if err != nil {
				errs = append(errs, err)
			}

			suggestions = append(suggestions, sugs...)
		}
	}

	if len(errs) > 0 {
		return descriptions, suggestions, bosherr.NewMultiError(errs...)
	}

	return descriptions, suggestions, nil
}

func (r LintableRelease) collectChecks() []check.Check {
	var checks []check.Check

	ctx := check.NewRootContext("Release")

	checks = append(checks, NewReleaseName(ctx, r.release.Name(), r.config.ReleaseName))
	checks = append(checks, NewMissingLicense(ctx, r.release.License(), r.config.MissingLicense))
	checks = append(checks, NewMissingJobs(ctx, r.release, r.config.MissingJobs))
	checks = append(checks, NewUnusedPackages(ctx, r.release, r.config.UnusedPackages))

	for _, job := range r.release.Jobs() {
		ctx := ctx.Nested(fmt.Sprintf("Job '%s'", job.Name()))

		checks = append(checks, check.NewDashedName(ctx, job.Name(), r.config.DashedName))
		checks = append(checks, NewJobPropertiesSyslogDaemonConfig(ctx, job, r.config.JobPropertiesSyslogDaemonConfig))
		checks = append(checks, NewJobPropertiesCertificate(ctx, job, r.config.JobPropertiesCertificate))
		checks = append(checks, NewJobTemplatesCtl(ctx, job, r.config.JobTemplatesCtl))

		for propName, propDef := range job.Properties {
			ctx := ctx.Nested(fmt.Sprintf("Property '%s'", propName))

			checks = append(checks, NewJobProperty(ctx, propName, propDef, r.config.JobProperty))
			checks = append(checks, NewJobPropertySecret(ctx, propName, propDef, r.config.JobPropertySecret))
			checks = append(checks, NewJobPropertySkipSSL(ctx, propName, r.config.JobPropertySkipSSL))
			checks = append(checks, NewJobPropertyDeprecated(ctx, propDef, r.config.JobPropertyDeprecated))
			checks = append(checks, NewJobPropertyNamespace(ctx, propName, job, r.config.JobPropertyNamespace))
			checks = append(checks, NewJobPropertyDebugAddr(ctx, propName, propDef, r.config.JobPropertyDebugAddr))
			checks = append(checks, check.NewTodo(ctx, propDef.Description, r.config.Todo))
		}
	}

	for _, pkg := range r.release.Packages() {
		ctx := ctx.Nested(fmt.Sprintf("Package '%s'", pkg.Name()))

		checks = append(checks, NewPackageName(ctx, pkg.Name(), r.config.PackageName))
	}

	return checks
}
