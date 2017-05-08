package release

import (
	"regexp"

	boshjob "github.com/cloudfoundry/bosh-cli/release/job"

	check "github.com/cppforlife/bosh-lint/check"
)

var (
	preferredJobProperty = regexp.MustCompile("^[a-z0-9]+(([a-z0-9]+\\_?)+[a-z0-9]+)?$")
)

type JobProperty struct {
	context check.Context
	name    string
	def     boshjob.PropertyDefinition
	check.CheckConfig
}

func NewJobProperty(context check.Context, name string, def boshjob.PropertyDefinition, config check.CheckConfig) JobProperty {
	return JobProperty{context, name, def, config}
}

func (c JobProperty) Description() check.Description {
	return check.Description{
		Context_: c.context,
		Purpose_: "if description is present",
	}
}

func (c JobProperty) Check() ([]check.Suggestion, error) {
	var sugs []check.Suggestion

	if len(c.def.Description) == 0 {
		sugs = append(sugs, check.Simple{
			Context_:    c.context,
			Problem_:    "Description is missing",
			Resolution_: "Add description",
		})
	}

	return sugs, nil
}
