package manifest

import (
	check "github.com/bosh-tools/bosh-ext-cli/check"
)

type IGProperties struct {
	context check.Context
	props   interface{}
	check.Config
}

func NewIGProperties(context check.Context, props interface{}, config check.Config) IGProperties {
	return IGProperties{context, props, config}
}

func (c IGProperties) Description() check.Description {
	return check.Description{
		Context_:   c.context,
		Purpose_:   "if instance group properties are present",
		Reasoning_: "It's recommended to specify job level properties instead of instance group properties.",
	}
}

func (c IGProperties) Check() ([]check.Suggestion, error) {
	var sugs []check.Suggestion

	if c.props != nil {
		sugs = append(sugs, check.Simple{
			Context_:    c.context,
			Problem_:    "Instance group properties are specified",
			Resolution_: "Remove instance group properties",
		})
	}

	return sugs, nil
}
