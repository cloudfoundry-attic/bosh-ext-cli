package manifest

import (
	check "github.com/cppforlife/bosh-lint/check"
)

type RootProperties struct {
	context check.Context
	props   interface{}
	check.Config
}

func NewRootProperties(context check.Context, props interface{}, config check.Config) RootProperties {
	return RootProperties{context, props, config}
}

func (c RootProperties) Description() check.Description {
	return check.Description{
		Context_:   c.context,
		Purpose_:   "if root properties are present",
		Reasoning_: "It's recommended to specify job level properties instead of root properties.",
	}
}

func (c RootProperties) Check() ([]check.Suggestion, error) {
	var sugs []check.Suggestion

	if c.props != nil {
		sugs = append(sugs, check.Simple{
			Context_:    c.context,
			Problem_:    "Root properties are specified",
			Resolution_: "Remove root properties",
		})
	}

	return sugs, nil
}
