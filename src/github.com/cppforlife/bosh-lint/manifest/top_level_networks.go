package manifest

import (
	check "github.com/cppforlife/bosh-lint/check"
)

type TopLevelNetworks struct {
	context  check.Context
	networks []Network
	check.Config
}

func NewTopLevelNetworks(context check.Context, networks []Network, config check.Config) TopLevelNetworks {
	return TopLevelNetworks{context, networks, config}
}

func (c TopLevelNetworks) Description() check.Description {
	return check.Description{
		Context_:   c.context,
		Purpose_:   "if top-level networks key is present",
		Reasoning_: "Top-level networks should be replaced with cloud config's networks.",
	}
}

func (c TopLevelNetworks) Check() ([]check.Suggestion, error) {
	var sugs []check.Suggestion

	if len(c.networks) > 0 {
		sugs = append(sugs, check.Simple{
			Context_:    c.context,
			Problem_:    "Top-level `networks` key is present",
			Resolution_: "Remove `networks` key",
		})
	}

	return sugs, nil
}
