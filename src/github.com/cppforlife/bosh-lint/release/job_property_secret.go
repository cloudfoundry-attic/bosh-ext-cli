package release

import (
	"strings"

	boshjob "github.com/cloudfoundry/bosh-cli/release/job"

	check "github.com/cppforlife/bosh-lint/check"
)

type JobPropertySecret struct {
	context check.Context
	name    string
	def     boshjob.PropertyDefinition
}

func NewJobPropertySecret(context check.Context, name string, def boshjob.PropertyDefinition) JobPropertySecret {
	return JobPropertySecret{context, name, def}
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

	for _, piece := range []string{"secret", "password", "token", "passphrase", "key"} {
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
