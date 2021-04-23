package types

type Domain struct {
	IdModel
	DnsName     string `json:"dnsName"`
	OrgName     string `json:"orgName"`
	StackName   string `json:"stackName"`
	Registrar   string `json:"registrar"`
	Certificate string `json:"certificate"`
	BlockId     int    `json:"blockId"`

	Block *Block `json:"block,omitempty"`
}
