package types

type AutogenSubdomain struct {
	IdModel
	Name        string   `json:"name"`
	OrgName     string   `json:"orgName"`
	DomainName  string   `json:"domainName"`
	Certificate string   `json:"certificate"`
	Nameservers []string `json:"nameservers"`
}
