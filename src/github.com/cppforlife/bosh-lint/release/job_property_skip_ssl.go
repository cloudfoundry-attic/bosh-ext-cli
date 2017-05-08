package release

import (
	"strings"

	check "github.com/cppforlife/bosh-lint/check"
)

type JobPropertySkipSSL struct {
	context check.Context
	name    string
	check.CheckConfig
}

func NewJobPropertySkipSSL(context check.Context, name string, config check.CheckConfig) JobPropertySkipSSL {
	return JobPropertySkipSSL{context, name, config}
}

func (c JobPropertySkipSSL) Description() check.Description {
	return check.Description{
		Context_: c.context,
		Purpose_: "if property allows to skip SSL verification",
	}
}

func (c JobPropertySkipSSL) Check() ([]check.Suggestion, error) {
	var sugs []check.Suggestion

	undesirablePieces := []string{
		"skip_tls",
		"skip_ssl",

		"use_tls",
		"use_ssl",

		"verify_tls",
		"verify_ssl",
		"verify_cert",
		"verify_certificate",

		"tls_verify",
		"ssl_verify",

		"require_tls",
		"require_ssl",

		"skip_tls_verify",
		"skip_ssl_verify",
		"skip_cert_verify",

		"skip_tls_validation",
		"skip_ssl_validation",
		"skip_cert_validation",
	}

	for _, piece := range undesirablePieces {
		if strings.Contains(c.name, piece) {
			sugs = append(sugs, check.Simple{
				Context_:    c.context,
				Problem_:    "Leads to a less secure system",
				Resolution_: "Remove ability to skip SSL verification",
			})
		}
	}

	return sugs, nil
}
