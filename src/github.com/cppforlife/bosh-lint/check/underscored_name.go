package check

import (
	"fmt"
	"regexp"
)

var (
	underscoredName = regexp.MustCompile("^[a-z0-9]+(([a-z0-9]+\\_?)+[a-z0-9]+)?$")
)

type UnderscoredName struct {
	context Context
	name    string
}

func NewUnderscoredName(context Context, name string) UnderscoredName {
	return UnderscoredName{context, name}
}

func (c UnderscoredName) Description() Description {
	return Description{
		Context_: c.context,
		Purpose_: fmt.Sprintf("if name matches suggested regexp '%s'", underscoredName),
	}
}

func (c UnderscoredName) Check() ([]Suggestion, error) {
	var sugs []Suggestion

	if !underscoredName.MatchString(c.name) {
		sugs = append(sugs, Simple{
			Context_:    c.context,
			Problem_:    fmt.Sprintf("Name does not match suggested regexp '%s'", underscoredName),
			Resolution_: "Rename",
		})
	}

	return sugs, nil
}
