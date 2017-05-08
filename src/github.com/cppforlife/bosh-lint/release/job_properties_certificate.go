package release

import (
	"fmt"
	"strings"

	boshjob "github.com/cloudfoundry/bosh-cli/release/job"

	check "github.com/cppforlife/bosh-lint/check"
)

type JobPropertiesCertificate struct {
	context check.Context
	job     *boshjob.Job
	check.CheckConfig
}

func NewJobPropertiesCertificate(context check.Context, job *boshjob.Job, config check.CheckConfig) JobPropertiesCertificate {
	return JobPropertiesCertificate{context, job, config}
}

func (c JobPropertiesCertificate) Description() check.Description {
	return check.Description{
		Context_: c.context,
		Purpose_: "if job can use certificate type properties",

		Reasoning_: markdown(`
It's recommended to replace multiple related certificate properties with a single @type: certificate@ property. For example:

Before:

@@@yaml
properties:
  ca_cert:
    description: "Trusted CA certificate for clients"
  server_cert:
    description: "Server certificate for TLS"
  server_key:
    description: "Server key for TLS"
@@@

After:

@@@yaml
properties:
  cert:
    type: certificate
    description: "Server certificate"
@@@

And use following accessors within ERB templates:

<%= p("policy_server.cert.ca") %>
<%= p("policy_server.cert.certificate") %>
<%= p("policy_server.cert.private_key") %>

Above syntax is compatible with certificates either explicitly provided by the operator or generated via config server API.
`),
	}
}

func (c JobPropertiesCertificate) Check() ([]check.Suggestion, error) {
	var sugs []check.Suggestion

	for propName, _ := range c.job.Properties {
		if strings.Contains(propName, "ca_cert") {
			sugs = append(sugs, check.Simple{
				Context_:    c.context.Nested(fmt.Sprintf("Property '%s'", propName)),
				Problem_:    "Asks for a certificate via multiple properties",
				Resolution_: "Replace related properties with a single `type: certificate` property",
			})
		}
	}

	return sugs, nil
}
