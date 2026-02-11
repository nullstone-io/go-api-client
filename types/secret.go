package types

import (
	"fmt"
)

type Secret struct {
	Identity SecretIdentity `json:"identity"`
	Metadata map[string]any `json:"metadata"`
	Value    string         `json:"value"`
	Redacted bool           `json:"redacted"`
}

const (
	SecretIdentityPlatformAws = "aws"
	SecretIdentityPlatformGcp = "gcp"
)

// SecretIdentity contains all metadata to uniquely identify a secret in platform's secrets manager
// AWS => `arn:aws:secretsmanager:{region}:{accountId}:secret:{secretName}`
// GCP => `projects/{projectId}/secrets/{secretName}`
type SecretIdentity struct {
	// Platform identifies the secrets manager being used
	Platform string `json:"platform"`

	Name string `json:"name"`

	// AWS-specific
	AwsRegion    string `json:"awsRegion,omitempty"`
	AwsAccountId string `json:"awsAccountId,omitempty"`

	// GCP-specific
	GcpProjectId string `json:"gcpProjectId,omitempty"`
}

func (i SecretIdentity) Id() string {
	switch i.Platform {
	case SecretIdentityPlatformAws:
		return fmt.Sprintf("arn:aws:secretsmanager:%s:%s:secret:%s", i.AwsRegion, i.AwsAccountId, i.Name)
	case SecretIdentityPlatformGcp:
		return fmt.Sprintf("projects/%s/secrets/%s", i.GcpProjectId, i.Name)
	default:
		return i.Name
	}
}

type SecretLocation struct {
	// AWS-specific
	AwsRegion    string `json:"awsRegion,omitempty" url:"aws_region,omitempty"`
	AwsAccountId string `json:"awsAccountId,omitempty" url:"aws_account_id,omitempty"`

	// GCP-specific
	GcpProjectId string `json:"gcpProjectId,omitempty" url:"gcp_project_id,omitempty"`
}
