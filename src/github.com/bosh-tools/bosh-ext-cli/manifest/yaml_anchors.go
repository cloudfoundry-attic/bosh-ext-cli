package manifest

import (
	"fmt"
	"regexp"
	"strings"

	check "github.com/bosh-tools/bosh-ext-cli/check"
)

var (
	yamlAnchors = regexp.MustCompile(`&([a-zA-Z_]+)[\r\t\n\s]+`)
)

type YAMLAnchors struct {
	context  check.Context
	contents string
	check.Config
}

func NewYAMLAnchors(context check.Context, bytes []byte, config check.Config) YAMLAnchors {
	return YAMLAnchors{context, string(bytes), config}
}

func (c YAMLAnchors) Description() check.Description {
	return check.Description{
		Context_:   c.context,
		Purpose_:   "if YAML anchors are used",
		Reasoning_: "Usage of YAML anchors typically indicates that releases could share information via links.",
	}
}

func (c YAMLAnchors) Check() ([]check.Suggestion, error) {
	var sugs []check.Suggestion

	var anchors, usedAnchors []string

	for _, matches := range yamlAnchors.FindAllStringSubmatch(c.contents, -1) {
		anchors = append(anchors, matches[1:]...)
	}

	for _, anchor := range anchors {
		if strings.Contains(c.contents, "*"+anchor) {
			usedAnchors = append(usedAnchors, anchor)
		}
	}

	if len(usedAnchors) > 0 {
		for _, anchor := range usedAnchors {
			sugs = append(sugs, check.Simple{
				Context_:    c.context,
				Problem_:    "YAML anchors are present",
				Resolution_: fmt.Sprintf("Replace YAML anchor '%s' with links", anchor),
			})
		}
	}

	return sugs, nil
}
