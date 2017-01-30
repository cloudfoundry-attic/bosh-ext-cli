package manifest

import (
	"fmt"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"gopkg.in/yaml.v2"

	check "github.com/cppforlife/bosh-lint/check"
)

type LintableManifest struct {
	bytes    []byte
	manifest Manifest
}

func NewLintableManifest(bytes []byte) (LintableManifest, error) {
	var manifest Manifest

	err := yaml.Unmarshal(bytes, &manifest)
	if err != nil {
		return LintableManifest{}, bosherr.WrapError(err, "Unmarshalling manifest")
	}

	return LintableManifest{bytes, manifest}, nil
}

func (m LintableManifest) Lint() ([]check.Description, []check.Suggestion, error) {
	var descriptions []check.Description
	var suggestions []check.Suggestion
	var errs []error

	for _, check := range m.collectChecks() {
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

func (m LintableManifest) collectChecks() []check.Check {
	var checks []check.Check

	ctx := check.NewRootContext("Manifest")

	checks = append(checks, check.NewUnderscoredName(ctx, m.manifest.Name))
	checks = append(checks, NewRootProperties(ctx, m.manifest.Properties))
	checks = append(checks, NewTopLevelJobs(ctx, m.manifest.Jobs))
	checks = append(checks, NewTopLevelNetworks(ctx, m.manifest.Networks))
	checks = append(checks, NewYAMLAnchors(ctx, m.bytes))
	checks = append(checks, NewVarInterpolation(ctx, m.bytes))

	for _, var_ := range m.manifest.Stemcells {
		ctx := ctx.Nested(fmt.Sprintf("Stemcell '%s'", var_.Alias))

		checks = append(checks, check.NewUnderscoredName(ctx, var_.Alias))
	}

	for _, ig := range m.manifest.InstanceGroups {
		ctx := ctx.Nested(fmt.Sprintf("Instance group '%s'", ig.Name))

		checks = append(checks, check.NewUnderscoredName(ctx, ig.Name))
		checks = append(checks, NewIGAZs(ctx, ig.AZs))
		checks = append(checks, NewIGProperties(ctx, ig.Properties))
		checks = append(checks, NewStaticIPs(ctx, ig.Networks))
	}

	for _, var_ := range m.manifest.Variables {
		ctx := ctx.Nested(fmt.Sprintf("Variable '%s'", var_.Name))

		checks = append(checks, check.NewUnderscoredName(ctx, var_.Name))
	}

	return checks
}
