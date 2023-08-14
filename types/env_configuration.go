package types

type EnvConfiguration struct {
	Version      string                            `yaml:"version" json:"version"`
	Subdomains   map[string]SubdomainConfiguration `yaml:"subdomains" json:"subdomains"`
	Applications map[string]AppConfiguration       `yaml:"apps" json:"apps"`
}

type SubdomainConfiguration struct {
	Name                string            `yaml:"-" json:"name"`
	DnsName             string            `yaml:"dns_name" json:"dnsName"`
	ModuleSource        string            `yaml:"module" json:"module"`
	ModuleSourceVersion *string           `yaml:"module_version" json:"moduleVersion"`
	Variables           map[string]any    `yaml:"vars" json:"vars"`
	Connections         ConnectionTargets `yaml:"connections" json:"connections"`
}

type AppConfiguration struct {
	Name                string                   `yaml:"-" json:"name"`
	ModuleSource        string                   `yaml:"module" json:"module"`
	ModuleSourceVersion *string                  `yaml:"module_version" json:"moduleVersion"`
	Variables           map[string]any           `yaml:"vars" json:"vars"`
	Capabilities        CapabilityConfigurations `yaml:"capabilities" json:"capabilities"`
	EnvVariables        map[string]string        `yaml:"environment" json:"envVars"`
}

type CapabilityConfigurations []CapabilityConfiguration

type CapabilityConfiguration struct {
	ModuleSource        string            `yaml:"module" json:"module"`
	ModuleSourceVersion *string           `yaml:"module_version" json:"moduleVersion"`
	Variables           map[string]any    `yaml:"vars" json:"vars"`
	Connections         ConnectionTargets `yaml:"connections" json:"connections"`
	Namespace           *string           `yaml:"namespace" json:"namespace"`
}

type ConnectionTargets map[string]ConnectionTarget
