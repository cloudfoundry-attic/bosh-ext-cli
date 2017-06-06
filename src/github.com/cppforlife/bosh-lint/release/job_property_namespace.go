package release

import (
	"fmt"
	"strings"

	boshjob "github.com/cloudfoundry/bosh-cli/release/job"

	check "github.com/cppforlife/bosh-lint/check"
)

type JobPropertyNamespace struct {
	context check.Context
	name    string
	job     *boshjob.Job
	check.CheckConfig
}

func NewJobPropertyNamespace(context check.Context, name string, job *boshjob.Job, config check.CheckConfig) JobPropertyNamespace {
	return JobPropertyNamespace{context, name, job, config}
}

func (c JobPropertyNamespace) Description() check.Description {
	return check.Description{
		Context_: c.context,
		Purpose_: "if property is unnecessarily namespaced",
	}
}

func (c JobPropertyNamespace) Check() ([]check.Suggestion, error) {
	var sugs []check.Suggestion

	ns := c.job.Name() + "."

	if strings.Contains(c.name, ns) {
		sugs = append(sugs, check.Simple{
			Context_:    c.context,
			Problem_:    "Unnecessarily namespaced with job name",
			Resolution_: fmt.Sprintf("Remove namespace '%s'", ns),
		})
	}

	return sugs, nil
}
