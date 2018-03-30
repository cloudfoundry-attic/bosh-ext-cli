package check

type Config struct {
	Disable bool `yaml:"disable"`
}

func (c Config) IsEnabled() bool {
	return !c.Disable
}
