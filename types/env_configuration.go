package types

type EnvConfiguration struct {
	Version      string                            `yaml:"version"`
	Subdomains   map[string]SubdomainConfiguration `yaml:"subdomains"`
	Applications map[string]AppConfiguration       `yaml:"apps"`
}

type SubdomainConfiguration struct {
	Name                string            `yaml:"-"`
	DnsName             string            `yaml:"dns_name"`
	ModuleSource        string            `yaml:"module"`
	ModuleSourceVersion *string           `yaml:"module_version"`
	Variables           map[string]any    `yaml:"vars"`
	Connections         ConnectionTargets `yaml:"connections"`
}

type AppConfiguration struct {
	Name                string                   `yaml:"-"`
	ModuleSource        string                   `yaml:"module"`
	ModuleSourceVersion *string                  `yaml:"module_version"`
	Variables           map[string]any           `yaml:"vars"`
	Capabilities        CapabilityConfigurations `yaml:"capabilities"`
	EnvVariables        map[string]string        `yaml:"environment"`
}

type CapabilityConfigurations []CapabilityConfiguration

type CapabilityConfiguration struct {
	ModuleSource        string            `yaml:"module"`
	ModuleSourceVersion *string           `yaml:"module_version"`
	Variables           map[string]any    `yaml:"vars"`
	Connections         ConnectionTargets `yaml:"connections"`
	Namespace           *string           `yaml:"namespace"`
}

type ConnectionTargets map[string]ConnectionTarget
