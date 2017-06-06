package manifest

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"gopkg.in/yaml.v2"

	check "github.com/cppforlife/bosh-lint/check"
)

type Config struct {
	DashedName       check.Config `yaml:"dashed_name"`
	RootProperties   check.Config `yaml:"root_properties"`
	TopLevelJobs     check.Config `yaml:"top_level_jobs"`
	TopLevelNetworks check.Config `yaml:"top_level_networks"`
	YAMLAnchors      check.Config `yaml:"yaml_anchors"`
	VarInterpolation check.Config `yaml:"var_interpolation"`

	IGAZs        check.Config `yaml:"ig_azs"`
	IGProperties check.Config `yaml:"ig_properties"`
	StaticIPs    check.Config `yaml:"static_ips"`
}

func NewConfig(bytes []byte) (Config, error) {
	config := Config{}

	err := yaml.Unmarshal(bytes, &config)
	if err != nil {
		return Config{}, bosherr.WrapError(err, "Unmarshalling config")
	}

	return config, nil
}
