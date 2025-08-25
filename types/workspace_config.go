package types

import (
	"fmt"
	"strings"

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
	Domain    *ExtraDomainConfig    `json:"domain,omitempty"`
	Subdomain *ExtraSubdomainConfig `json:"subdomain,omitempty"`
}

type ExtraDomainConfig struct {
	// DomainNameTemplate is a template for configuring DomainName
	// This allows for interpolating the following template variables:
	//  - {{ NULLSTONE_ORG }}
	DomainNameTemplate string `json:"domainNameTemplate,omitempty"`

	// DomainName refers to the Domain's full name
	// Normally, this refers to a second-level domain (SLD).
	DomainName string `json:"domainName,omitempty"`

	// Fqdn refers to the fully qualified domain name
	// This is equal to `<domain-name>.`
	// NOTE: This has a trailing "."
	Fqdn string `json:"fqdn,omitempty"`
}

func (c ExtraDomainConfig) Equal(other ExtraDomainConfig) bool {
	return c.DomainNameTemplate == other.DomainNameTemplate &&
		c.DomainName == other.DomainName &&
		c.Fqdn == other.Fqdn
}

type ExtraSubdomainConfig struct {
	// SubdomainNameTemplate is a template for configuring SubdomainName
	// This allows for interpolating the following template variables:
	//  - {{ random() }}
	//  - {{ NULLSTONE_ORG }}
	//  - {{ NULLSTONE_STACK }}
	//  - {{ NULLSTONE_ENV }}
	SubdomainNameTemplate string `json:"subdomainNameTemplate,omitempty"`

	// SubdomainName refers to the subdomain name for this subdomain
	// Normally, this is equivalent to `<dns-name>[.<env-chunk>]`
	// This is the FQDN without the domain name
	SubdomainName string `json:"subdomainName,omitempty"`

	// DomainName refers to the Subdomain's domain or the Domain's full name
	// Normally, this refers to a second-level domain (SLD).
	DomainName string `json:"domainName,omitempty"`

	// Fqdn refers to the fully qualified domain name
	// This is equal to `<subdomain-name>.<domain-name>.`
	// NOTE: This has a trailing "."
	Fqdn string `json:"fqdn,omitempty"`
}

func (c ExtraSubdomainConfig) CalculateFqdn() string {
	fqdn := fmt.Sprintf("%s.%s.", c.SubdomainName, c.DomainName)
	fqdn = strings.TrimSuffix(fqdn, "..")
	fqdn = strings.TrimPrefix(fqdn, ".")
	return fqdn
}

func (c ExtraSubdomainConfig) Equal(other ExtraSubdomainConfig) bool {
	return c.SubdomainNameTemplate == other.SubdomainNameTemplate &&
		c.SubdomainName == other.SubdomainName &&
		c.DomainName == other.DomainName &&
		c.Fqdn == other.Fqdn
}
