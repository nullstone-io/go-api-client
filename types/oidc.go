package types

type OidcDiscoveryDocument struct {
	Issuer                           string   `json:"issuer"`
	JwksUri                          string   `json:"jwks_uri"`
	IdTokenSigningAlgValuesSupported []string `json:"id_token_signing_alg_values_supported"`
	ResponseTypesSupported           []string `json:"response_types_supported"`
	SubjectTypesSupported            []string `json:"subject_types_supported"`
	ClaimsSupported                  []string `json:"claims_supported"`
}

type IssueTokenInput struct {
	SubjectType string `json:"subject_type"`
	OrgName     string `json:"org_name"`
	StackName   string `json:"stack_name,omitempty"`
	BlockName   string `json:"block_name,omitempty"`
	EnvName     string `json:"env_name,omitempty"`
	Audience    string `json:"audience,omitempty"`
	TtlSeconds  int    `json:"ttl_seconds,omitempty"`
}

type IssueTokenOutput struct {
	Token string `json:"token"`
}
