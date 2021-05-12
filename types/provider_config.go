package types

type ProviderConfig struct {
	Aws *AwsProviderConfig `json:"aws"`
}

type AwsProviderConfig struct {
	ProviderName string `json:"providerName"`
	Region       string `json:"region"`
}
