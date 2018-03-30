package manifest

import (
	check "github.com/bosh-tools/bosh-ext-cli/lint/check"
)

type IGAZs struct {
	context check.Context
	azs     *[]string
	check.Config
}

func NewIGAZs(context check.Context, azs *[]string, config check.Config) IGAZs {
	return IGAZs{context, azs, config}
}

func (c IGAZs) Description() check.Description {
	return check.Description{
		Context_:   c.context,
		Purpose_:   "if instance group AZs are present",
		Reasoning_: "It's recommended to use AZs with instance groups.",
	}
}

func (c IGAZs) Check() ([]check.Suggestion, error) {
	var sugs []check.Suggestion

	if c.azs == nil {
		sugs = append(sugs, check.Simple{
			Context_:    c.context,
			Problem_:    "Instance group AZs are not specified",
			Resolution_: "Specify instance group AZs",
		})
	}

	return sugs, nil
}
