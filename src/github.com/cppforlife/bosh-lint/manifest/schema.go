package manifest

type Manifest struct {
	Name string

	Networks []Network
	Jobs     []Job

	Stemcells      []Stemcell
	InstanceGroups []InstanceGroup `yaml:"instance_groups"`

	Properties interface{}
	Variables  []Variable
}

type Network struct{}
type Job struct{}

type Stemcell struct {
	Alias string
}

type InstanceGroup struct {
	Name string
	AZs  *[]string

	Properties interface{}

	Consumes interface{}
	Provides interface{}

	Networks []NetworkAssociation
}

type NetworkAssociation struct {
	Name      string
	Default   []string
	StaticIPs []string `yaml:"static_ips"`
}

type Variable struct {
	Name string
}
