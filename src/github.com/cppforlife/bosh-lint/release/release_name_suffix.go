package release

import (
	"fmt"
	"strings"

	check "github.com/cppforlife/bosh-lint/check"
)

type ReleaseNameSuffix struct {
	context check.Context
	name    string
	check.Config
}

func NewReleaseNameSuffix(context check.Context, name string, config check.Config) ReleaseNameSuffix {
	return ReleaseNameSuffix{context, name, config}
}

func (c ReleaseNameSuffix) Description() check.Description {
	return check.Description{
		Context_: c.context,
		Purpose_: "if name does not have redundant suffix",
	}
}

func (c ReleaseNameSuffix) Check() ([]check.Suggestion, error) {
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

	return sugs, nil
}
