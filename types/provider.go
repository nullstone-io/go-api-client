package types

import (
	"encoding/json"
)

type Provider struct {
	IdModel
	Name         string `json:"name"`
	OrgName      string `json:"orgName"`
	ProviderType string `json:"providerType"`

	// ProviderId represents a single namespace for the provider
	//   AWS: Account ID
	//   GCP: Project ID
	ProviderId string `json:"providerId"`

	Credentials json.RawMessage `json:"credentials,omitempty"`
}
