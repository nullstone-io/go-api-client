package api

type Subdomain struct {
	IdModel
	Name         string `json:"name"`
	OrgName      string `json:"orgName"`
	StackName    string `json:"stackName"`
	ModuleSource string `json:"moduleSource"`
	Certificate  string `json:"certificate"`
	DomainId     int    `json:"domainId"`

	Domain Domain `json:"domain,omitempty"`
}
