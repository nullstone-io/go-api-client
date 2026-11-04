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

	// Metadata carries extensible governance/descriptive metadata about the
	// block (e.g. data classification). This is the canonical, block-level,
	// environment-invariant declaration that IaC writes and the UI edits; it
	// seeds WorkspaceConfig.Metadata when a workspace is instantiated.
	Metadata WorkspaceMetadata `json:"metadata"`

	// Build is the block-level default build configuration for an Application block.
	// It seeds the per-workspace build config when an app workspace is first read and
	// is deliberately *not* part of WorkspaceConfig: build config takes effect on the
	// next build without an infra run. nil means "no block default"; consumers fall
	// back to DefaultBuildConfig(module).
	Build *BuildConfig `json:"build,omitempty"`
}

type WorkspaceCapabilityTemplateConfig struct {
	Name             string                      `json:"name"`
	Namespace        string                      `json:"namespace"`
	Module           string                      `json:"module"`
	ModuleConstraint string                      `json:"moduleConstraint"`
	Connections      map[string]ConnectionTarget `json:"connections"`
}
