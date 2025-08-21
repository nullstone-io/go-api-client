package types

import (
	"github.com/jinzhu/copier"
)

type WorkspaceConfig struct {
	// Source refers to the module used for this workspace
	Source string `json:"source"`
	// SourceConstraint is a constraint or desired version for the workspace module
	// Once resolved, SourceVersion contains the effective module version
	SourceConstraint string `json:"sourceConstraint"`
	// SourceVersion refers to the effective module version
	// Variables and Connections on this WorkspaceConfig should match the schema for this module version
	SourceVersion string `json:"sourceVersion"`
	// SourceToolName refers to the tool used to execute this SourceVersion module
	// Examples: terraform, opentofu
	SourceToolName string               `json:"sourceToolName"`
	Variables      Variables            `json:"variables"`
	EnvVariables   EnvVariables         `json:"envVariables"`
	Connections    Connections          `json:"connections"`
	Providers      Providers            `json:"providers"`
	Capabilities   CapabilityConfigs    `json:"capabilities"`
	Extra          ExtraWorkspaceConfig `json:"extra"`

	// Dependencies represents a list of workspace references that are necessary for this run
	// This is saved to the run config so that a user can quickly access a list of dependencies
	//   It *should not* be used by the nullfire engine to pull in dependencies
	//   because a cleanup run excludes dependencies that are detached, but not destroyed
	Dependencies Dependencies `json:"dependencies"`

	// This field is used to capture user specific configuration for unlaunched dependencies
	// that require some configuration in order to be launched for the first time.
	// Two examples are:
	// - An application that is connected to an unlaunched "Existing Network"
	// - An application that is connected to an unlaunched Datadog datastore via a capability
	DependencyConfigs DependencyConfigs `json:"dependencyConfigs"`
}

func (c WorkspaceConfig) Clone() (WorkspaceConfig, error) {
	config := WorkspaceConfig{}
	err := copier.CopyWithOption(&config, c, copier.Option{DeepCopy: true})
	return config, err
}

type ExtraWorkspaceConfig struct {
	// SubdomainName refers to the subdomain name for this subdomain
	// Normally, this is equivalent to `<dns-name>[.<env-chunk>]`
	// This is the FQDN without the domain name
	// Valid Blocks: Subdomain
	SubdomainName string `json:"dnsName,omitempty"`
	// DomainName refers to the Subdomain's domain or the Domain's full name
	// Normally, this refers to a second-level domain (SLD).
	// Valid Blocks: Subdomain, Domain
	DomainName string `json:"domainName,omitempty"`
	// Fqdn refers to the fully qualified domain name
	// This is equal to `<subdomain-name>.<domain-name>.`
	// NOTE: This has a trailing "."
	// Valid Blocks: Subdomain, Domain
	Fqdn string `json:"fqdn,omitempty"`
}
