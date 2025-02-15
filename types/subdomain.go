package types

import "github.com/google/uuid"

type Subdomain struct {
	Block       `json:",inline"`
	Certificate string `json:"certificate"`
}

type SubdomainWorkspace struct {
	// WorkspaceUid refers to the uid of the workspace (stack/block/env)
	WorkspaceUid uuid.UUID `json:"workspaceUid"`

	// DnsName refers to the first token in the full subdomain
	// This is configured by the user
	// If DnsName==".", this subdomain acts as a passthrough to the domain
	DnsName string `json:"dnsName"`

	// SubdomainName refers to the subdomain name for this subdomain
	// If Nullstone-managed, this is equal to `<dns-name>[.<env-chunk>]`
	// This should be the FQDN without the domain name
	SubdomainName string `json:"subdomainName"`

	// DomainName refers to the domain name for this subdomain
	// Normally, this refers to a second-level domain (SLD).
	// Example: "acme.com"
	DomainName string `json:"domainName"`

	// Fqdn refers to the fully-qualified domain name
	// This is equal to `<subdomain-name>.<domain-name>.`
	// NOTE: This has a trailing "."
	Fqdn string `json:"fqdn"`
}
