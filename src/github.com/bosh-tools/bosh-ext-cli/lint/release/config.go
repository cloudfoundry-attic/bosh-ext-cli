package release

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"gopkg.in/yaml.v2"

	check "github.com/bosh-tools/bosh-ext-cli/lint/check"
)

type Config struct {
	ReleaseName       check.Config `yaml:"release_name"`
	ReleaseNameSuffix check.Config `yaml:"release_name_suffix"`
	MissingLicense    check.Config `yaml:"missing_license"`
	MissingJobs       check.Config `yaml:"missing_jobs"`
	UnusedPackages    check.Config `yaml:"unused_packages"`

	JobName                         check.Config `yaml:"job_name"`
	JobPropertiesSyslogDaemonConfig check.Config `yaml:"job_properties_syslog_daemon_config"`
	JobPropertiesCertificate        check.Config `yaml:"job_properties_certificate"`
	JobTemplatesCtl                 check.Config `yaml:"job_templates_ctl"`

	JobProperty           check.Config               `yaml:"job_property"`
	JobPropertySecret     JobPropertySecretConfig    `yaml:"job_property_secret"`
	JobPropertySkipSSL    check.Config               `yaml:"job_property_skip_ssl"`
	JobPropertyDeprecated check.Config               `yaml:"job_property_deprecated"`
	JobPropertyNamespace  check.Config               `yaml:"job_property_namespace"`
	JobPropertyDebugAddr  JobPropertyDebugAddrConfig `yaml:"job_property_debug_addr"`

	PackageName check.Config `yaml:"package_name"`

	Todo check.Config `yaml:"todo"`
}

type JobPropertyDebugAddrConfig struct {
	DebugPatterns []string `yaml:"debug_patterns"`
	Whitelist     []string `yaml:"whitelist"`
	check.Config
}

type JobPropertySecretConfig struct {
	SecretPatterns []string `yaml:"secret_patterns"`
	Whitelist      []string `yaml:"whitelist"`
	check.Config
}

func NewConfig(bytes []byte) (Config, error) {
	config := Config{
		JobPropertySecret: JobPropertySecretConfig{
			SecretPatterns: []string{"secret", "password", "token", "passphrase", "key"},
			Whitelist:      []string{},
		},
		JobPropertyDebugAddr: JobPropertyDebugAddrConfig{
			DebugPatterns: []string{"debug"},
			Whitelist:     []string{},
		},
	}

	err := yaml.Unmarshal(bytes, &config)
	if err != nil {
		return Config{}, bosherr.WrapError(err, "Unmarshalling config")
	}

	return config, nil
}
