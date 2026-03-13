package types

type ProviderConfig struct {
	Aws   *AwsProviderConfig   `json:"aws,omitempty"`
	Gcp   *GcpProviderConfig   `json:"gcp,omitempty"`
	Azure *AzureProviderConfig `json:"azure,omitempty"`
}

type AwsProviderConfig struct {
	ProviderName string `json:"providerName"`
	Region       string `json:"region"`
}

type GcpProviderConfig struct {
	ProviderName string `json:"providerName"`
	Region       string `json:"region"`
	Zone         string `json:"zone"`
}

type AzureProviderConfig struct {
	ProviderName string `json:"providerName"`
	Location     string `json:"location"`
}
