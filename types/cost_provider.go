package types

type CostProvider struct {
	Id               int64     `json:"id"`
	OrgName          string    `json:"orgName"`
	ProviderId       int64     `json:"providerId"`
	Provider         *Provider `json:"provider"`
	IsConfigured     bool      `json:"isConfigured"`
	IncludedAccounts []string  `json:"includedAccounts"`

	// GCP cost reporting reads from a BigQuery billing export. The user supplies the dataset;
	// the host project is the provider's own project, and the export table is discovered when
	// the cost provider is tested. Both are empty for AWS/Azure.
	BillingExportDataset string `json:"billingExportDataset,omitempty"`
	BillingExportTable   string `json:"billingExportTable,omitempty"`
}
