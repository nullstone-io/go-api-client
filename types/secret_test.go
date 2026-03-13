package types

import (
	"testing"
)

func TestParseSecretIdentity(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected SecretIdentity
	}{
		// AWS valid cases
		{
			name:  "AWS ARN - valid format",
			input: "arn:aws:secretsmanager:us-east-1:123456789012:secret:my-secret",
			expected: SecretIdentity{
				SecretLocation: SecretLocation{
					Platform:     SecretLocationPlatformAws,
					AwsRegion:    "us-east-1",
					AwsAccountId: "123456789012",
				},
				Name: "my-secret",
			},
		},
		{
			name:  "AWS ARN - with dashes in secret name",
			input: "arn:aws:secretsmanager:eu-west-2:987654321098:secret:my-database-secret",
			expected: SecretIdentity{
				SecretLocation: SecretLocation{
					Platform:     SecretLocationPlatformAws,
					AwsRegion:    "eu-west-2",
					AwsAccountId: "987654321098",
				},
				Name: "my-database-secret",
			},
		},
		{
			name:  "AWS ARN - different region",
			input: "arn:aws:secretsmanager:ap-southeast-1:111222333444:secret:prod-db-creds",
			expected: SecretIdentity{
				SecretLocation: SecretLocation{
					Platform:     SecretLocationPlatformAws,
					AwsRegion:    "ap-southeast-1",
					AwsAccountId: "111222333444",
				},
				Name: "prod-db-creds",
			},
		},
		// AWS invalid cases
		{
			name:  "AWS ARN - invalid service",
			input: "arn:aws:s3:us-east-1:123456789012:bucket:my-bucket",
			expected: SecretIdentity{
				Name: "arn:aws:s3:us-east-1:123456789012:bucket:my-bucket",
			},
		},
		{
			name:  "AWS ARN - wrong number of parts (too few)",
			input: "arn:aws:secretsmanager:us-east-1:123456789012",
			expected: SecretIdentity{
				Name: "arn:aws:secretsmanager:us-east-1:123456789012",
			},
		},
		{
			name:  "AWS ARN - wrong number of parts (too many)",
			input: "arn:aws:secretsmanager:us-east-1:123456789012:secret:my-secret:extra",
			expected: SecretIdentity{
				Name: "arn:aws:secretsmanager:us-east-1:123456789012:secret:my-secret:extra",
			},
		},
		{
			name:  "AWS ARN - starts with arn:aws: but wrong format",
			input: "arn:aws:invalid:format",
			expected: SecretIdentity{
				Name: "arn:aws:invalid:format",
			},
		},
		// GCP valid cases - 4 parts
		{
			name:  "GCP - 4 parts format",
			input: "projects/my-project/secrets/my-secret",
			expected: SecretIdentity{
				SecretLocation: SecretLocation{
					Platform:     SecretLocationPlatformGcp,
					GcpProjectId: "my-project",
				},
				Name: "my-secret",
			},
		},
		{
			name:  "GCP - 4 parts with numbers in project ID",
			input: "projects/project-123/secrets/api-key",
			expected: SecretIdentity{
				SecretLocation: SecretLocation{
					Platform:     SecretLocationPlatformGcp,
					GcpProjectId: "project-123",
				},
				Name: "api-key",
			},
		},
		// GCP valid cases - 6 parts (locations format without version)
		{
			name:  "GCP - 6 parts format (locations)",
			input: "projects/my-project/locations/us-central1/secrets/my-secret",
			expected: SecretIdentity{
				SecretLocation: SecretLocation{
					Platform:     SecretLocationPlatformGcp,
					GcpProjectId: "my-project",
				},
				Name: "my-secret",
			},
		},
		// GCP valid cases - 8 parts (locations + versions)
		{
			name:  "GCP - 8 parts format (locations + versions)",
			input: "projects/my-project/locations/us-central1/secrets/my-secret/versions/1",
			expected: SecretIdentity{
				SecretLocation: SecretLocation{
					Platform:     SecretLocationPlatformGcp,
					GcpProjectId: "my-project",
				},
				Name: "my-secret",
			},
		},
		{
			name:  "GCP - 8 parts with latest version",
			input: "projects/my-project/locations/us-east1/secrets/db-password/versions/latest",
			expected: SecretIdentity{
				SecretLocation: SecretLocation{
					Platform:     SecretLocationPlatformGcp,
					GcpProjectId: "my-project",
				},
				Name: "db-password",
			},
		},
		// GCP invalid cases
		{
			name:  "GCP - 4 parts invalid keyword (not secrets)",
			input: "projects/my-project/invalid/my-secret",
			expected: SecretIdentity{
				Name: "projects/my-project/invalid/my-secret",
			},
		},
		{
			name:  "GCP - 6 parts invalid keyword (not secrets at index 4)",
			input: "projects/my-project/locations/us-central1/buckets/my-bucket",
			expected: SecretIdentity{
				Name: "projects/my-project/locations/us-central1/buckets/my-bucket",
			},
		},
		{
			name:  "GCP - unsupported number of parts (5)",
			input: "projects/my-project/secrets/my-secret/extra",
			expected: SecretIdentity{
				Name: "projects/my-project/secrets/my-secret/extra",
			},
		},
		{
			name:  "GCP - unsupported number of parts (3)",
			input: "projects/my-project/secrets",
			expected: SecretIdentity{
				Name: "projects/my-project/secrets",
			},
		},
		// Azure valid cases
		{
			name:  "Azure - secret without version",
			input: "https://my-vault.vault.azure.net/secrets/my-secret",
			expected: SecretIdentity{
				SecretLocation: SecretLocation{
					Platform:       SecretLocationPlatformAzure,
					AzureVaultName: "my-vault",
				},
				Name: "my-secret",
			},
		},
		{
			name:  "Azure - secret with version",
			input: "https://my-vault.vault.azure.net/secrets/my-secret/abc123def456",
			expected: SecretIdentity{
				SecretLocation: SecretLocation{
					Platform:           SecretLocationPlatformAzure,
					AzureVaultName:     "my-vault",
					AzureSecretVersion: "abc123def456",
				},
				Name: "my-secret",
			},
		},
		{
			name:  "Azure - vault name with dashes",
			input: "https://my-company-prod-vault.vault.azure.net/secrets/db-connection-string",
			expected: SecretIdentity{
				SecretLocation: SecretLocation{
					Platform:       SecretLocationPlatformAzure,
					AzureVaultName: "my-company-prod-vault",
				},
				Name: "db-connection-string",
			},
		},
		{
			name:  "Azure - secret name with dashes",
			input: "https://vault1.vault.azure.net/secrets/my-api-key-v2",
			expected: SecretIdentity{
				SecretLocation: SecretLocation{
					Platform:       SecretLocationPlatformAzure,
					AzureVaultName: "vault1",
				},
				Name: "my-api-key-v2",
			},
		},
		// Azure invalid cases
		{
			name:     "Azure - not vault.azure.net domain",
			input:    "https://my-vault.blob.azure.net/secrets/my-secret",
			expected: SecretIdentity{Name: "https://my-vault.blob.azure.net/secrets/my-secret"},
		},
		{
			name:     "Azure - http instead of https",
			input:    "http://my-vault.vault.azure.net/secrets/my-secret",
			expected: SecretIdentity{Name: "http://my-vault.vault.azure.net/secrets/my-secret"},
		},
		{
			name:     "Azure - not /secrets/ path",
			input:    "https://my-vault.vault.azure.net/keys/my-key",
			expected: SecretIdentity{Name: "https://my-vault.vault.azure.net/keys/my-key"},
		},
		// Default/fallback cases
		{
			name:     "Simple name",
			input:    "simple-secret-name",
			expected: SecretIdentity{Name: "simple-secret-name"},
		},
		{
			name:     "Empty string",
			input:    "",
			expected: SecretIdentity{Name: ""},
		},
		{
			name:     "Name with special characters",
			input:    "secret-with_123.dash",
			expected: SecretIdentity{Name: "secret-with_123.dash"},
		},
		{
			name:     "Name with slashes but not projects prefix",
			input:    "some/random/path",
			expected: SecretIdentity{Name: "some/random/path"},
		},
		{
			name:     "HTTPS URL but not azure",
			input:    "https://example.com/secrets/foo",
			expected: SecretIdentity{Name: "https://example.com/secrets/foo"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseSecretIdentity(tt.input)
			if result != tt.expected {
				t.Errorf("ParseSecretIdentity(%q) = %+v, want %+v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestSecretIdentityId(t *testing.T) {
	tests := []struct {
		name     string
		identity SecretIdentity
		expected string
	}{
		{
			name: "AWS identity",
			identity: SecretIdentity{
				SecretLocation: SecretLocation{
					Platform:     SecretLocationPlatformAws,
					AwsRegion:    "us-west-2",
					AwsAccountId: "123456789012",
				},
				Name: "database-secret",
			},
			expected: "arn:aws:secretsmanager:us-west-2:123456789012:secret:database-secret",
		},
		{
			name: "AWS identity - different region",
			identity: SecretIdentity{
				SecretLocation: SecretLocation{
					Platform:     SecretLocationPlatformAws,
					AwsRegion:    "ap-southeast-1",
					AwsAccountId: "999888777666",
				},
				Name: "prod-api-key",
			},
			expected: "arn:aws:secretsmanager:ap-southeast-1:999888777666:secret:prod-api-key",
		},
		{
			name: "GCP identity",
			identity: SecretIdentity{
				SecretLocation: SecretLocation{
					Platform:     SecretLocationPlatformGcp,
					GcpProjectId: "my-gcp-project",
				},
				Name: "api-key",
			},
			expected: "projects/my-gcp-project/secrets/api-key",
		},
		{
			name: "Azure identity - without version",
			identity: SecretIdentity{
				SecretLocation: SecretLocation{
					Platform:       SecretLocationPlatformAzure,
					AzureVaultName: "my-vault",
				},
				Name: "my-secret",
			},
			expected: "https://my-vault.vault.azure.net/secrets/my-secret",
		},
		{
			name: "Azure identity - with version",
			identity: SecretIdentity{
				SecretLocation: SecretLocation{
					Platform:           SecretLocationPlatformAzure,
					AzureVaultName:     "my-vault",
					AzureSecretVersion: "abc123",
				},
				Name: "my-secret",
			},
			expected: "https://my-vault.vault.azure.net/secrets/my-secret/abc123",
		},
		{
			name: "Azure identity - empty version treated as no version",
			identity: SecretIdentity{
				SecretLocation: SecretLocation{
					Platform:           SecretLocationPlatformAzure,
					AzureVaultName:     "prod-vault",
					AzureSecretVersion: "",
				},
				Name: "db-password",
			},
			expected: "https://prod-vault.vault.azure.net/secrets/db-password",
		},
		{
			name: "Default/unknown platform",
			identity: SecretIdentity{
				Name: "simple-secret",
			},
			expected: "simple-secret",
		},
		{
			name: "Unknown platform string",
			identity: SecretIdentity{
				SecretLocation: SecretLocation{
					Platform: "unknown",
				},
				Name: "some-secret",
			},
			expected: "some-secret",
		},
		{
			name:     "Empty identity",
			identity: SecretIdentity{},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.identity.Id()
			if result != tt.expected {
				t.Errorf("SecretIdentity.Id() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestParseSecretIdentityRoundTrip(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "AWS roundtrip",
			input: "arn:aws:secretsmanager:us-east-1:123456789012:secret:my-secret",
		},
		{
			name:  "GCP roundtrip",
			input: "projects/my-project/secrets/my-secret",
		},
		{
			name:  "Azure roundtrip without version",
			input: "https://my-vault.vault.azure.net/secrets/my-secret",
		},
		{
			name:  "Azure roundtrip with version",
			input: "https://my-vault.vault.azure.net/secrets/my-secret/v1",
		},
		{
			name:  "Simple name roundtrip",
			input: "plain-secret",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identity := ParseSecretIdentity(tt.input)
			result := identity.Id()
			if result != tt.input {
				t.Errorf("Roundtrip failed: ParseSecretIdentity(%q).Id() = %q", tt.input, result)
			}
		})
	}
}
