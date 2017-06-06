package manifest

import (
	"fmt"

	boshtpl "github.com/cloudfoundry/bosh-cli/director/template"

	check "github.com/cppforlife/bosh-lint/check"
)

type VarInterpolation struct {
	context check.Context
	bytes   []byte
	check.Config
}

func NewVarInterpolation(context check.Context, bytes []byte, config check.Config) VarInterpolation {
	return VarInterpolation{context, bytes, config}
}

func (c VarInterpolation) Description() check.Description {
	return check.Description{
		Context_:   c.context,
		Purpose_:   "if variables are interpolated multiple times",
		Reasoning_: "Same variable interpolation in multiple locations typically indicates that releases could share information via links.",
	}
}

type trackingVars struct {
	typedNames map[string]struct{}
	askedNames map[string]int
}

func (v *trackingVars) TimesAsked(name string) int {
	if _, found := v.typedNames[name]; found {
		return v.askedNames[name] - 1
	}
	return v.askedNames[name]
}

func (v *trackingVars) Get(def boshtpl.VariableDefinition) (interface{}, bool, error) {
	v.askedNames[def.Name] += 1

	if len(def.Type) > 0 {
		v.typedNames[def.Name] = struct{}{}
	}

	// todo cheating a bit...
	switch {
	case def.Type == "rsa" || def.Type == "ssh":
		return map[interface{}]interface{}{"public_key": "", "public_key_fingerprint": "", "private_key": ""}, true, nil

	case def.Type == "certificate":
		return map[interface{}]interface{}{"ca": "", "certificate": "", "private_key": ""}, true, nil

	default:
		return "", true, nil
	}
}

func (v *trackingVars) List() ([]boshtpl.VariableDefinition, error) { panic("Not implemented") }

var _ boshtpl.Variables = &trackingVars{}

func (c VarInterpolation) Check() ([]check.Suggestion, error) {
	var sugs []check.Suggestion

	vars := &trackingVars{
		typedNames: map[string]struct{}{},
		askedNames: map[string]int{},
	}

	tpl := boshtpl.NewTemplate(c.bytes)

	_, err := tpl.Evaluate(vars, nil, boshtpl.EvaluateOpts{})
	if err != nil {
		return nil, err
	}

	for name, _ := range vars.askedNames {
		if vars.TimesAsked(name) > 1 {
			sugs = append(sugs, check.Simple{
				Context_:    c.context,
				Problem_:    "Variable interpolated more than once",
				Resolution_: fmt.Sprintf("Replace variable '%s' secondary interpolation with links", name),
			})
		}
	}

	return sugs, nil
}
