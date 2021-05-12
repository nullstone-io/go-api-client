package types

type AutogenSubdomain struct {
	IdModel
	OrgName     string      `json:"orgName"`
	DnsName     string      `json:"dnsName"`
	DomainName  string      `json:"domainName"`
	Fqdn        string      `json:"fqdn"`
	Certificate string      `json:"certificate"`
	Nameservers Nameservers `json:"nameservers"`
}

type Nameservers []string
