package manifest

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"gopkg.in/yaml.v2"

	check "github.com/bosh-tools/bosh-ext-cli/lint/check"
)

type Config struct {
	ManifestName check.Config `yaml:"manifest_name"`

	RootProperties   check.Config `yaml:"root_properties"`
	TopLevelJobs     check.Config `yaml:"top_level_jobs"`
	TopLevelNetworks check.Config `yaml:"top_level_networks"`
	YAMLAnchors      check.Config `yaml:"yaml_anchors"`
	VarInterpolation check.Config `yaml:"var_interpolation"`

	StemcellAlias check.Config `yaml:"stemcell_alias"`

	IGName       check.Config `yaml:"instance_group_name"`
	IGAZs        check.Config `yaml:"instance_group_azs"`
	IGStemcell   check.Config `yaml:"instance_group_stemcell"`
	IGProperties check.Config `yaml:"instance_group_properties"`
	IGLinks      check.Config `yaml:"instance_group_links"`
	StaticIPs    check.Config `yaml:"static_ips"`

	VariableName check.Config `yaml:"variable_name"`
}

func NewConfig(bytes []byte) (Config, error) {
	config := Config{}

	err := yaml.Unmarshal(bytes, &config)
	if err != nil {
		return Config{}, bosherr.WrapError(err, "Unmarshalling config")
	}

	return config, nil
}
