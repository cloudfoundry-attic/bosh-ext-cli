package manifest

import (
	check "github.com/cppforlife/bosh-lint/check"
)

type IGStemcell struct {
	context  check.Context
	stemcell interface{}
	check.Config
}

func NewIGStemcell(context check.Context, stemcell interface{}, config check.Config) IGStemcell {
	return IGStemcell{context, stemcell, config}
}

func (c IGStemcell) Description() check.Description {
	return check.Description{
		Context_:   c.context,
		Purpose_:   "if instance group stemcell is not present",
		Reasoning_: "It's mandatory to specify stemcell under a instance group.",
	}
}

func (c IGStemcell) Check() ([]check.Suggestion, error) {
	var sugs []check.Suggestion

	if c.stemcell == nil {
		sugs = append(sugs, check.Simple{
			Context_:    c.context,
			Problem_:    "Instance group stemcell is not specified",
			Resolution_: "Add instance group stemcell",
		})
	}

	return sugs, nil
}
