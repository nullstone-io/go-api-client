package types

import "fmt"

const (
	GcpAuthTypeServiceAccount             = "serviceAccount"
	GcpAuthTypeWorkloadIdentityFederation = "workloadIdentityFederation"
)

type GcpCredentials struct {
	AuthType          string               `json:"authType"`
	ServiceAccountKey GcpServiceAccountKey `json:"serviceAccount"`
	WorkloadIdentity  GcpWorkloadIdentity  `json:"workloadIdentity"`
}

type GcpServiceAccountKey struct {
	Type                    string `json:"type"`
	ProjectId               string `json:"project_id"`
	PrivateKeyId            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientId                string `json:"client_id"`
	AuthUri                 string `json:"auth_uri"`
	TokenUri                string `json:"token_uri"`
	AuthProviderX509CertUrl string `json:"auth_provider_x509_cert_url"`
	ClientX509CertUrl       string `json:"client_x509_cert_url"`
}

type GcpWorkloadIdentity struct {
	ProjectNumber       string `json:"projectNumber"`
	ProjectId           string `json:"projectId"`
	ServiceAccountEmail string `json:"serviceAccountEmail"`

	IdentityPoolId         string `json:"identityPoolId"`
	IdentityPoolProviderId string `json:"identityPoolProviderId"`
}

// Audience should be a reference to the user's Workload Identity Pool Provider
func (i GcpWorkloadIdentity) Audience() string {
	return fmt.Sprintf("//iam.googleapis.com/projects/%s/locations/global/workloadIdentityPools/%s/providers/%s", i.ProjectNumber, i.IdentityPoolId, i.IdentityPoolProviderId)
}
