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
	SecretLocationPlatformAws   = "aws"
	SecretLocationPlatformGcp   = "gcp"
	SecretLocationPlatformAzure = "azure"
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
	} else if strings.HasPrefix(input, "https://") && strings.Contains(input, ".vault.azure.net/secrets/") {
		// Parse Azure Key Vault secret URI: https://{vaultName}.vault.azure.net/secrets/{secretName}[/{secretVersion}]
		// Remove the protocol prefix
		uriWithoutProtocol := strings.TrimPrefix(input, "https://")
		// Split by .vault.azure.net to get vault name
		parts := strings.SplitN(uriWithoutProtocol, ".vault.azure.net/secrets/", 2)
		if len(parts) == 2 {
			vaultName := parts[0]
			secretPath := parts[1]
			// Split the secret path to get the secret name and optional version
			secretParts := strings.SplitN(secretPath, "/", 2)
			secretName := secretParts[0]
			var secretVersion string
			if len(secretParts) == 2 {
				secretVersion = secretParts[1]
			}
			return SecretIdentity{
				SecretLocation: SecretLocation{
					Platform:           SecretLocationPlatformAzure,
					AzureVaultName:     vaultName,
					AzureSecretVersion: secretVersion,
				},
				Name: secretName,
			}
		}
	}
	return SecretIdentity{Name: input}
}

// SecretIdentity contains all metadata to uniquely identify a secret in platform's secrets manager
// AWS => `arn:aws:secretsmanager:{region}:{accountId}:secret:{secretName}`
// GCP => `projects/{projectId}/secrets/{secretName}`
// Azure => `https://{vaultName}.vault.azure.net/secrets/{secretName}[/{secretVersion}]`
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
	case SecretLocationPlatformAzure:
		if i.AzureSecretVersion != "" {
			return fmt.Sprintf("https://%s.vault.azure.net/secrets/%s/%s", i.AzureVaultName, i.Name, i.AzureSecretVersion)
		}
		return fmt.Sprintf("https://%s.vault.azure.net/secrets/%s", i.AzureVaultName, i.Name)
	default:
		return i.Name
	}
}

type SecretLocation struct {
	// Platform identifies the secrets manager being used
	Platform string `json:"platform" url:"platform,omitempty"`

	// AWS-specific
	AwsRegion    string `json:"awsRegion,omitempty" url:"aws_region,omitempty"`
	AwsAccountId string `json:"awsAccountId,omitempty" url:"aws_account_id,omitempty"`

	// GCP-specific
	GcpProjectId string `json:"gcpProjectId,omitempty" url:"gcp_project_id,omitempty"`

	// Azure-specific
	AzureVaultName     string `json:"azureVaultName,omitempty" url:"azure_vault_name,omitempty"`
	AzureSecretVersion string `json:"azureSecretVersion,omitempty" url:"azure_secret_version,omitempty"`
}
