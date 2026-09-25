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

	// AWS cost reporting reads Cost Explorer by default. A FOCUS 1.2 Data Export
	// (s3://<bucket>/<prefix>/<export-name>) adds list/contracted cost and resource-level detail
	// for the months it covers. The bucket region is discovered when the cost provider is tested.
	// Both are empty for GCP/Azure and for AWS providers on the Cost Explorer tier.
	DataExportPath   string `json:"dataExportPath,omitempty"`
	DataExportRegion string `json:"dataExportRegion,omitempty"`
}

// CostProviderInput is what a user supplies to create or test a cost provider: the cloud
// provider to read costs for plus the provider-type-specific export settings.
type CostProviderInput struct {
	ProviderId int64 `json:"providerId"`
	// GCP: the BigQuery billing export dataset (required).
	BillingExportDataset string `json:"billingExportDataset,omitempty"`
	// AWS: the Data Export location (optional).
	DataExportPath string `json:"dataExportPath,omitempty"`
}
