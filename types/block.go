package types

type Block struct {
	IdModel
	Type                string                      `json:"type"`
	OrgName             string                      `json:"orgName"`
	StackId             int64                       `json:"stackId"`
	Reference           string                      `json:"reference"`
	Name                string                      `json:"name"`
	DnsName             string                      `json:"dnsName"`
	ModuleSource        string                      `json:"moduleSource"`
	ModuleSourceVersion string                      `json:"moduleSourceVersion"`
	Connections         map[string]ConnectionTarget `json:"connections"`
}
