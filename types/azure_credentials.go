package types

const (
	AzureAuthTypeWorkloadIdentityFederation = "workloadIdentityFederation"
)

type AzureCredentials struct {
	AuthType string `json:"authType"`
	TenantId string `json:"tenantId"`
	ClientId string `json:"clientId"`
}
