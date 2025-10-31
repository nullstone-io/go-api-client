package types

type CostProvider struct {
	Id               int64     `json:"id"`
	OrgName          string    `json:"orgName"`
	ProviderId       int64     `json:"providerId"`
	Provider         *Provider `json:"provider"`
	IsConfigured     bool      `json:"isConfigured"`
	IncludedAccounts []string  `json:"includedAccounts"`
}
