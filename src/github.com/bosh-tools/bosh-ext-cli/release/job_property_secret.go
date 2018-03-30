package release

import (
	"strings"

	boshjob "github.com/cloudfoundry/bosh-cli/release/job"

	check "github.com/bosh-tools/bosh-ext-cli/check"
)

type JobPropertySecret struct {
	context check.Context
	name    string
	def     boshjob.PropertyDefinition
	JobPropertySecretConfig
}

func NewJobPropertySecret(context check.Context, name string, def boshjob.PropertyDefinition, config JobPropertySecretConfig) JobPropertySecret {
	return JobPropertySecret{context, name, def, config}
}

func (c JobPropertySecret) Description() check.Description {
	return check.Description{
		Context_: c.context,
		Purpose_: "if property represents a secret and should not have a default",
	}
}

func (c JobPropertySecret) Check() ([]check.Suggestion, error) {
	var sugs []check.Suggestion

	if c.def.Default == nil {
		return nil, nil
	}

	for _, piece := range c.Whitelist {
		if strings.Contains(c.name, piece) {
			return nil, nil
		}
	}

	for _, piece := range c.SecretPatterns {
		if strings.Contains(c.name, piece) {
			sugs = append(sugs, check.Simple{
				Context_:    c.context,
				Problem_:    "Property holding sensitive value should not have a default",
				Resolution_: "Remove default",
			})

			break
		}
	}

	return sugs, nil
}
