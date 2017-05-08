package release

import (
	"strings"

	boshjob "github.com/cloudfoundry/bosh-cli/release/job"

	check "github.com/cppforlife/bosh-lint/check"
)

type JobPropertyDeprecated struct {
	context check.Context
	def     boshjob.PropertyDefinition
	check.CheckConfig
}

func NewJobPropertyDeprecated(context check.Context, def boshjob.PropertyDefinition, config check.CheckConfig) JobPropertyDeprecated {
	return JobPropertyDeprecated{context, def, config}
}

func (c JobPropertyDeprecated) Description() check.Description {
	return check.Description{
		Context_: c.context,
		Purpose_: "if property is deprecated",
	}
}

func (c JobPropertyDeprecated) Check() ([]check.Suggestion, error) {
	var sugs []check.Suggestion

	if strings.Contains(strings.ToLower(c.def.Description), "deprecated") {
		sugs = append(sugs, check.Simple{
			Context_:    c.context,
			Problem_:    "Deprecated property",
			Resolution_: "Remove",
		})
	}

	return sugs, nil
}
