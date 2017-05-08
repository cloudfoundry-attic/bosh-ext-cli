package release

import (
	"fmt"
	"regexp"
	"strings"

	check "github.com/cppforlife/bosh-lint/check"
)

var (
	preferredReleaseName = regexp.MustCompile("^[a-z0-9]+(([a-z0-9]+\\-?)+[a-z0-9]+)?$")
)

type ReleaseName struct {
	context check.Context
	name    string
	check.CheckConfig
}

func NewReleaseName(context check.Context, name string, config check.CheckConfig) ReleaseName {
	return ReleaseName{context, name, config}
}

func (c ReleaseName) Description() check.Description {
	return check.Description{
		Context_: c.context,
		Purpose_: "if name matches suggested regexp",
	}
}

func (c ReleaseName) Check() ([]check.Suggestion, error) {
	var sugs []check.Suggestion

	for _, suffix := range []string{"boshrelease", "bosh-release", "release"} {
		if strings.Contains(c.name, suffix) {
			sugs = append(sugs, check.Simple{
				Context_:    c.context,
				Problem_:    fmt.Sprintf("Name redundantly ends with '%s'", suffix),
				Resolution_: fmt.Sprintf("Remove suffix '%s'", suffix),
			})

			return sugs, nil
		}
	}

	if !preferredReleaseName.MatchString(c.name) {
		sugs = append(sugs, check.Simple{
			Context_:    c.context,
			Problem_:    fmt.Sprintf("Name does not match suggested regexp '%s'", preferredReleaseName),
			Resolution_: "Rename",
		})
	}

	return sugs, nil
}
