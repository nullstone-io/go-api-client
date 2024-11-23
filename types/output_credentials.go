package types

const (
	ProviderAws   = "aws"
	ProviderGcp   = "gcp"
	ProviderAzure = "azure"
)

const (
	CredentialsTypeAwsAssumeRole      = "aws-assume-role"
	CredentialsTypeAwsGetSessionToken = "aws-get-session-token"
)

type OutputCredentials struct {
	// Provider refers to the cloud provider (e.g. aws, gcp)
	Provider string `json:"provider"`

	// CredentialsType
	CredentialsType string `json:"credentialsType"`

	// Data contains the credentials information
	Data map[string]string `json:"data"`
}
