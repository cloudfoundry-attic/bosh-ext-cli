package manifest

import (
	check "github.com/bosh-tools/bosh-ext-cli/check"
)

// todo IGLinks should really be fatally validated in the Director
type IGLinks struct {
	context            check.Context
	consumes, provides interface{}
	check.Config
}

func NewIGLinks(context check.Context, consumes, provides interface{}, config check.Config) IGLinks {
	return IGLinks{context, consumes, provides, config}
}

func (c IGLinks) Description() check.Description {
	return check.Description{
		Context_:   c.context,
		Purpose_:   "if consumes/provides keys are present on an instance group",
		Reasoning_: "It's an error to specify consumes/provides keys at instance group level.",
	}
}

func (c IGLinks) Check() ([]check.Suggestion, error) {
	var sugs []check.Suggestion

	if c.consumes != nil {
		sugs = append(sugs, check.Simple{
			Context_:    c.context,
			Problem_:    "`consumes` key specified on an instance group",
			Resolution_: "Move `consumes` keys into a job",
		})
	}

	if c.provides != nil {
		sugs = append(sugs, check.Simple{
			Context_:    c.context,
			Problem_:    "`provides` key specified on an instance group",
			Resolution_: "Move `provides` keys into a job",
		})
	}

	return sugs, nil
}
