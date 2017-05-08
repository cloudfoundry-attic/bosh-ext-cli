package manifest

import (
	"fmt"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"gopkg.in/yaml.v2"

	check "github.com/cppforlife/bosh-lint/check"
)

type ManifestConfig struct {
	DashedName       check.CheckConfig `yaml:"dashed_name"`
	RootProperties   check.CheckConfig `yaml:"root_properties"`
	TopLevelJobs     check.CheckConfig `yaml:"top_level_jobs"`
	TopLevelNetworks check.CheckConfig `yaml:"top_level_networks"`
	YAMLAnchors      check.CheckConfig `yaml:"yaml_anchors"`
	VarInterpolation check.CheckConfig `yaml:"var_interpolation"`

	IGAZs        check.CheckConfig `yaml:"ig_azs"`
	IGProperties check.CheckConfig `yaml:"ig_properties"`
	StaticIPs    check.CheckConfig `yaml:"static_ips"`
}

var DefaultManifestConfig = ManifestConfig{}

type LintableManifest struct {
	bytes    []byte
	manifest Manifest
	config   ManifestConfig
}

func NewLintableManifest(bytes []byte, config ManifestConfig) (LintableManifest, error) {
	var manifest Manifest

	err := yaml.Unmarshal(bytes, &manifest)
	if err != nil {
		return LintableManifest{}, bosherr.WrapError(err, "Unmarshalling manifest")
	}

	return LintableManifest{bytes, manifest, config}, nil
}

func (m LintableManifest) Lint() ([]check.Description, []check.Suggestion, error) {
	var descriptions []check.Description
	var suggestions []check.Suggestion
	var errs []error

	for _, check := range m.collectChecks() {
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

func (m LintableManifest) collectChecks() []check.Check {
	var checks []check.Check

	ctx := check.NewRootContext("Manifest")

	checks = append(checks, check.NewDashedName(ctx, m.manifest.Name, m.config.DashedName))
	checks = append(checks, NewRootProperties(ctx, m.manifest.Properties, m.config.RootProperties))
	checks = append(checks, NewTopLevelJobs(ctx, m.manifest.Jobs, m.config.TopLevelJobs))
	checks = append(checks, NewTopLevelNetworks(ctx, m.manifest.Networks, m.config.TopLevelNetworks))
	checks = append(checks, NewYAMLAnchors(ctx, m.bytes, m.config.YAMLAnchors))
	checks = append(checks, NewVarInterpolation(ctx, m.bytes, m.config.VarInterpolation))

	for _, var_ := range m.manifest.Stemcells {
		ctx := ctx.Nested(fmt.Sprintf("Stemcell '%s'", var_.Alias))

		checks = append(checks, check.NewDashedName(ctx, var_.Alias, m.config.DashedName))
	}

	for _, ig := range m.manifest.InstanceGroups {
		ctx := ctx.Nested(fmt.Sprintf("Instance group '%s'", ig.Name))

		checks = append(checks, check.NewDashedName(ctx, ig.Name, m.config.DashedName))
		checks = append(checks, NewIGAZs(ctx, ig.AZs, m.config.IGAZs))
		checks = append(checks, NewIGProperties(ctx, ig.Properties, m.config.IGProperties))
		checks = append(checks, NewStaticIPs(ctx, ig.Networks, m.config.StaticIPs))
	}

	for _, var_ := range m.manifest.Variables {
		ctx := ctx.Nested(fmt.Sprintf("Variable '%s'", var_.Name))

		checks = append(checks, check.NewDashedName(ctx, var_.Name, m.config.DashedName))
	}

	return checks
}
