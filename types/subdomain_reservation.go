package types

type SubdomainReservation struct {
	IsRandom      bool   `json:"isRandom"`
	SubdomainName string `json:"subdomainName"`
	DomainName    string `json:"domainName"`
}
