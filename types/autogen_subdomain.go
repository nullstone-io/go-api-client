package types

type AutogenSubdomain struct {
	IdModel
	DnsName     string   `json:"dnsName"`
	OrgName     string   `json:"orgName"`
	DomainName  string   `json:"domainName"`
	Certificate string   `json:"certificate"`
	Nameservers []string `json:"nameservers"`
}
