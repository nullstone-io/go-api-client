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
			name:  "AWS ARN - invalid service",
			input: "arn:aws:s3:us-east-1:123456789012:bucket:my-bucket",
			expected: SecretIdentity{
				Name: "arn:aws:s3:us-east-1:123456789012:bucket:my-bucket",
			},
		},
		{
			name:  "AWS ARN - wrong number of parts",
			input: "arn:aws:secretsmanager:us-east-1:123456789012",
			expected: SecretIdentity{
				Name: "arn:aws:secretsmanager:us-east-1:123456789012",
			},
		},
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
			name:  "GCP - 6 parts format",
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
			name:  "GCP - with numbers in project ID",
			input: "projects/project-123/secrets/api-key",
			expected: SecretIdentity{
				SecretLocation: SecretLocation{
					Platform:     SecretLocationPlatformGcp,
					GcpProjectId: "project-123",
				},
				Name: "api-key",
			},
		},
		{
			name:  "GCP - invalid format",
			input: "projects/my-project/invalid/my-secret",
			expected: SecretIdentity{
				Name: "projects/my-project/invalid/my-secret",
			},
		},
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
			name:  "AWS ARN - starts with arn:aws: but wrong format",
			input: "arn:aws:invalid:format",
			expected: SecretIdentity{
				Name: "arn:aws:invalid:format",
			},
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
			name: "Simple identity",
			identity: SecretIdentity{
				Name: "simple-secret",
			},
			expected: "simple-secret",
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
