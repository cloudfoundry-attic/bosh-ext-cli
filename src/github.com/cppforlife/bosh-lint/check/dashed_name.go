package check

import (
	"fmt"
	"regexp"
)

var (
	dashedName = regexp.MustCompile("^[a-z0-9]+(([a-z0-9]+\\-?)+[a-z0-9]+)?$")
)

type DashedName struct {
	context Context
	name    string
}

func NewDashedName(context Context, name string) DashedName {
	return DashedName{context, name}
}

func (c DashedName) Description() Description {
	return Description{
		Context_: c.context,
		Purpose_: fmt.Sprintf("if name matches suggested regexp '%s'", dashedName),
	}
}

func (c DashedName) Check() ([]Suggestion, error) {
	var sugs []Suggestion

	if !dashedName.MatchString(c.name) {
		sugs = append(sugs, Simple{
			Context_:    c.context,
			Problem_:    fmt.Sprintf("Name does not match suggested regexp '%s'", dashedName),
			Resolution_: "Rename",
		})
	}

	return sugs, nil
}
