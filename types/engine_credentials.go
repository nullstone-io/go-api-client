package types

import "encoding/json"

type EnvEngineCredentials struct {
	Providers []ProviderEngineCredentials `json:"providers"`
	Warnings  []string                    `json:"warnings,omitempty"`
}

type ProviderEngineCredentials struct {
	Provider        Provider          `json:"provider"`
	CredentialsType string            `json:"credentialsType"`
	RawCredentials  json.RawMessage   `json:"rawCredentials"`
	Data            map[string]string `json:"data"`
}
