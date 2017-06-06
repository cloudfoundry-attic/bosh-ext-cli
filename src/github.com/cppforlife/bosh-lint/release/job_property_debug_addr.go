package release

import (
	"strings"

	boshjob "github.com/cloudfoundry/bosh-cli/release/job"

	check "github.com/cppforlife/bosh-lint/check"
)

type JobPropertyDebugAddr struct {
	context check.Context
	name    string
	def     boshjob.PropertyDefinition
	JobPropertyDebugAddrConfig
}

func NewJobPropertyDebugAddr(context check.Context, name string, def boshjob.PropertyDefinition, config JobPropertyDebugAddrConfig) JobPropertyDebugAddr {
	return JobPropertyDebugAddr{context, name, def, config}
}

func (c JobPropertyDebugAddr) Description() check.Description {
	return check.Description{
		Context_: c.context,
		Purpose_: "if property represents a debug address and should not have open address default",
	}
}

func (c JobPropertyDebugAddr) Check() ([]check.Suggestion, error) {
	var sugs []check.Suggestion

	if c.def.Default == nil {
		return nil, nil
	}

	for _, piece := range c.Whitelist {
		if strings.Contains(c.def.Description, piece) {
			return nil, nil
		}
	}

	defaultStr, ok := c.def.Default.(string)
	if !ok {
		return nil, nil
	}

	if strings.Contains(defaultStr, "0.0.0.0") {
		for _, piece := range c.DebugPatterns {
			if strings.Contains(c.name, piece) {
				sugs = append(sugs, check.Simple{
					Context_:    c.context,
					Problem_:    "Property holding debug address should not use '0.0.0.0'",
					Resolution_: "Update default to use '127.0.0.1'",
				})
			}
		}
	}

	return sugs, nil
}
