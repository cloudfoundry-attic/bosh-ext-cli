package release

import (
	"fmt"
	"strings"

	boshjob "github.com/cloudfoundry/bosh-cli/release/job"

	check "github.com/cppforlife/bosh-lint/check"
)

type JobTemplatesCtl struct {
	context check.Context
	job     *boshjob.Job
	check.Config
}

func NewJobTemplatesCtl(context check.Context, job *boshjob.Job, config check.Config) JobTemplatesCtl {
	return JobTemplatesCtl{context, job, config}
}

func (c JobTemplatesCtl) Description() check.Description {
	return check.Description{
		Context_:   c.context,
		Purpose_:   "if ctl script is unnecessarily namespaced",
		Reasoning_: "Job templates are placed under /var/vcap/jobs/{job}/ hence it's unnecessary to prefix ctl script (unless you have multiple ctl scripts).",
	}
}

func (c JobTemplatesCtl) Check() ([]check.Suggestion, error) {
	var sugs []check.Suggestion

	var ctlPaths []string

	for _, dstPath := range c.job.Templates {
		if strings.Contains(dstPath, "ctl") {
			ctlPaths = append(ctlPaths, dstPath)
		}
	}

	if len(ctlPaths) == 1 && ctlPaths[0] != "bin/ctl" {
		sugs = append(sugs, check.Simple{
			Context_:    c.context,
			Problem_:    "Unnecessarily namespaced ctl script",
			Resolution_: fmt.Sprintf("Rename '%s' to 'bin/ctl'", ctlPaths[0]),
		})
	}

	return sugs, nil
}
