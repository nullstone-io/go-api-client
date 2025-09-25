package types

type WorkspaceTemplate struct {
	OrgName   string                  `json:"orgName"`
	StackId   int64                   `json:"stackId"`
	BlockId   int64                   `json:"blockId"`
	BlockType string                  `json:"blockType"`
	Config    WorkspaceTemplateConfig `json:"config"`
}

type WorkspaceTemplateConfig struct {
	Module           string            `json:"module"`
	ModuleConstraint string            `json:"moduleConstraint"`
	Connections      ConnectionTargets `json:"connections"`

	// Capabilities provide a template list of Capabilities (including Module, ModuleConstraint, and Connections) for an Application Workspace
	Capabilities []WorkspaceCapabilityTemplateConfig `json:"capabilities"`

	// DomainName provides a domain name on a Domain Workspace
	DomainName string `json:"domainName"`

	// SubdomainNameTemplate provides a template for creating a subdomain name on a Subdomain Workspace
	SubdomainNameTemplate string `json:"subdomainNameTemplate"`
}

type WorkspaceCapabilityTemplateConfig struct {
	Name             string                      `json:"name"`
	Namespace        string                      `json:"namespace"`
	Module           string                      `json:"module"`
	ModuleConstraint string                      `json:"moduleConstraint"`
	Connections      map[string]ConnectionTarget `json:"connections"`
}
