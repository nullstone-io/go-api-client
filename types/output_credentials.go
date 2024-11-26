package types

import "time"

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

	// Aws contains aws credentials
	Aws *OutputCredentialsAws `json:"aws,omitempty"`

	// Gcp contains gcp credentials
	Gcp *OutputCredentialsGcp `json:"gcp,omitempty"`

	// Data contains additional credential information
	Data map[string]string `json:"data"`
}

type OutputCredentialsAws struct {
	// AWS Access key ID
	AccessKeyID string `json:"accessKeyID"`

	// AWS Secret Access Key
	SecretAccessKey string `json:"secretAccessKey"`

	// AWS Session Token
	SessionToken string `json:"sessionToken"`

	// Source of the credentials
	Source string `json:"source"`

	// States if the credentials can expire or not.
	CanExpire bool `json:"canExpire"`

	// The time the credentials will expire at. Should be ignored if CanExpire
	// is false.
	Expires time.Time `json:"expires"`
}

type OutputCredentialsGcp struct {
	// AccessToken is the token that authorizes and authenticates the requests.
	AccessToken string `json:"accessToken"`

	// TokenType is the type of token.
	// The Type method returns either this or "Bearer", the default.
	TokenType string `json:"tokenType"`

	// RefreshToken is a token that's used by the application
	// (as opposed to the user) to refresh the access token
	// if it expires.
	RefreshToken string `json:"refreshToken"`

	// Expiry is the optional expiration time of the access token.
	Expiry time.Time `json:"expiry"`
}
