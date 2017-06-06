package release

import (
	"fmt"
	"regexp"

	check "github.com/cppforlife/bosh-lint/check"
)

var (
	preferredPackageName = regexp.MustCompile("^[a-z0-9]+(([a-z0-9]+\\_?)+[a-z0-9]+)?$")
)

type PackageName struct {
	context check.Context
	name    string
	check.CheckConfig
}

func NewPackageName(context check.Context, name string, config check.CheckConfig) PackageName {
	return PackageName{context, name, config}
}

func (c PackageName) Description() check.Description {
	return check.Description{
		Context_: c.context,
		Purpose_: fmt.Sprintf("if name matches suggested regexp '%s'", preferredPackageName),
	}
}

func (c PackageName) Check() ([]check.Suggestion, error) {
	var sugs []check.Suggestion

	if !preferredPackageName.MatchString(c.name) {
		sugs = append(sugs, check.Simple{
			Context_:    c.context,
			Problem_:    fmt.Sprintf("Name does not match suggested regexp '%s'", preferredPackageName),
			Resolution_: "Rename",
		})
	}

	return sugs, nil
}
