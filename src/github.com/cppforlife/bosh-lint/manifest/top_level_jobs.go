package manifest

import (
	check "github.com/cppforlife/bosh-lint/check"
)

type TopLevelJobs struct {
	context check.Context
	jobs    []Job
}

func NewTopLevelJobs(context check.Context, jobs []Job) TopLevelJobs {
	return TopLevelJobs{context, jobs}
}

func (c TopLevelJobs) Description() check.Description {
	return check.Description{
		Context_:   c.context,
		Purpose_:   "if top-level jobs key is present",
		Reasoning_: "Top-level jobs are now called instance groups.",
	}
}

func (c TopLevelJobs) Check() ([]check.Suggestion, error) {
	var sugs []check.Suggestion

	if len(c.jobs) > 0 {
		sugs = append(sugs, check.Simple{
			Context_:    c.context,
			Problem_:    "Top-level `jobs` key is present",
			Resolution_: "Rename `jobs` key to `instance_groups`",
		})
	}

	return sugs, nil
}
