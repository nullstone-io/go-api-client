package types

type Subdomain struct {
	IdModel
	DnsName      string `json:"dnsName"`
	OrgName      string `json:"orgName"`
	StackName    string `json:"stackName"`
	Certificate  string `json:"certificate"`
	BlockId   	 int    `json:"blockId"`

	Block 		 *Block `json:"block,omitempty"`
}
