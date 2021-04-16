package types

type AutogenSubdomainDelegation struct {
	Nameservers Nameservers `json:"nameservers"`
}

type Nameservers []string
