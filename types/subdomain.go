package types

type Subdomain struct {
	Block
	DnsName     string `json:"dnsName"`
	Certificate string `json:"certificate"`
}
