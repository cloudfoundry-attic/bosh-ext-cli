package release

import (
	boshlic "github.com/cloudfoundry/bosh-cli/release/license"

	check "github.com/cppforlife/bosh-lint/check"
)

type MissingLicense struct {
	context check.Context
	license *boshlic.License
	check.CheckConfig
}

func NewMissingLicense(context check.Context, license *boshlic.License, config check.CheckConfig) MissingLicense {
	return MissingLicense{context, license, config}
}

func (c MissingLicense) Description() check.Description {
	return check.Description{
		Context_: c.context,
		Purpose_: "if license is present",
	}
}

func (c MissingLicense) Check() ([]check.Suggestion, error) {
	var sugs []check.Suggestion

	if c.license == nil {
		sugs = append(sugs, check.Simple{
			Context_:    c.context,
			Problem_:    "Missing license",
			Resolution_: "Add LICENSE and NOTICE files",
		})
	}

	return sugs, nil
}
