package manifest

import (
	"fmt"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"gopkg.in/yaml.v2"

	check "github.com/bosh-tools/bosh-ext-cli/check"
)

type LintableManifest struct {
	bytes    []byte
	manifest Manifest
	config   Config
}

func NewLintableManifest(bytes []byte, config Config) (LintableManifest, error) {
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

	checks = append(checks, check.NewDashedName(ctx, m.manifest.Name, m.config.ManifestName))
	checks = append(checks, NewRootProperties(ctx, m.manifest.Properties, m.config.RootProperties))
	checks = append(checks, NewTopLevelJobs(ctx, m.manifest.Jobs, m.config.TopLevelJobs))
	checks = append(checks, NewTopLevelNetworks(ctx, m.manifest.Networks, m.config.TopLevelNetworks))
	checks = append(checks, NewYAMLAnchors(ctx, m.bytes, m.config.YAMLAnchors))
	checks = append(checks, NewVarInterpolation(ctx, m.bytes, m.config.VarInterpolation))

	for _, stem := range m.manifest.Stemcells {
		ctx := ctx.Nested(fmt.Sprintf("Stemcell '%s'", stem.Alias))

		checks = append(checks, check.NewDashedName(ctx, stem.Alias, m.config.StemcellAlias))
	}

	for _, ig := range m.manifest.InstanceGroups {
		ctx := ctx.Nested(fmt.Sprintf("Instance group '%s'", ig.Name))

		checks = append(checks, check.NewDashedName(ctx, ig.Name, m.config.IGName))
		checks = append(checks, NewIGAZs(ctx, ig.AZs, m.config.IGAZs))
		checks = append(checks, NewIGStemcell(ctx, ig.Stemcell, m.config.IGStemcell))
		checks = append(checks, NewIGProperties(ctx, ig.Properties, m.config.IGProperties))
		checks = append(checks, NewIGLinks(ctx, ig.Consumes, ig.Provides, m.config.IGLinks))
		checks = append(checks, NewStaticIPs(ctx, ig.Networks, m.config.StaticIPs))
	}

	for _, var_ := range m.manifest.Variables {
		ctx := ctx.Nested(fmt.Sprintf("Variable '%s'", var_.Name))

		checks = append(checks, check.NewDashedName(ctx, var_.Name, m.config.VariableName))
	}

	return checks
}
