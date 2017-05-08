package check

type CheckConfig struct {
	Disable bool `yaml:"disable"`
}

func (pc CheckConfig) IsEnabled() bool {
	return !pc.Disable
}
