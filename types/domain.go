package types

type Domain struct {
	Block
	DnsName     string `json:"dnsName"`
	Registrar   string `json:"registrar"`
	Certificate string `json:"certificate"`
}
