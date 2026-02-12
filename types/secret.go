package types

import (
	"fmt"
	"strings"
)

type Secret struct {
	Identity SecretIdentity `json:"identity"`
	Metadata map[string]any `json:"metadata"`
	Value    string         `json:"value"`
	Redacted bool           `json:"redacted"`
}

const (
	SecretLocationPlatformAws = "aws"
	SecretLocationPlatformGcp = "gcp"
)

func ParseSecretIdentity(input string) SecretIdentity {
	if strings.HasPrefix(input, "arn:aws:") {
		tokens := strings.Split(input, ":")
		if len(tokens) == 7 && tokens[2] == "secretsmanager" {
			return SecretIdentity{
				SecretLocation: SecretLocation{
					Platform:     SecretLocationPlatformAws,
					AwsRegion:    tokens[3],
					AwsAccountId: tokens[4],
				},
				Name: tokens[6],
			}
		}
	} else if strings.HasPrefix(input, "projects/") {
		tokens := strings.Split(input, "/")
		switch len(tokens) {
		case 4:
			if tokens[2] == "secrets" {
				return SecretIdentity{
					SecretLocation: SecretLocation{
						Platform:     SecretLocationPlatformGcp,
						GcpProjectId: tokens[1],
					},
					Name: tokens[3],
				}
			}
		case 6:
			fallthrough
		case 8:
			if tokens[4] == "secrets" {
				return SecretIdentity{
					SecretLocation: SecretLocation{
						Platform:     SecretLocationPlatformGcp,
						GcpProjectId: tokens[1],
					},
					Name: tokens[5],
				}
			}
		}
	}
	return SecretIdentity{Name: input}
}

// SecretIdentity contains all metadata to uniquely identify a secret in platform's secrets manager
// AWS => `arn:aws:secretsmanager:{region}:{accountId}:secret:{secretName}`
// GCP => `projects/{projectId}/secrets/{secretName}`
type SecretIdentity struct {
	SecretLocation `json:",inline"`

	Name string `json:"name"`
}

func (i SecretIdentity) Id() string {
	switch i.Platform {
	case SecretLocationPlatformAws:
		return fmt.Sprintf("arn:aws:secretsmanager:%s:%s:secret:%s", i.AwsRegion, i.AwsAccountId, i.Name)
	case SecretLocationPlatformGcp:
		return fmt.Sprintf("projects/%s/secrets/%s", i.GcpProjectId, i.Name)
	default:
		return i.Name
	}
}

type SecretLocation struct {
	// Platform identifies the secrets manager being used
	Platform string `json:"platform"`

	// AWS-specific
	AwsRegion    string `json:"awsRegion,omitempty" url:"aws_region,omitempty"`
	AwsAccountId string `json:"awsAccountId,omitempty" url:"aws_account_id,omitempty"`

	// GCP-specific
	GcpProjectId string `json:"gcpProjectId,omitempty" url:"gcp_project_id,omitempty"`
}
