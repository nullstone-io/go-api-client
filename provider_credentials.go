package api

import (
	"encoding/json"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"net/http"
	"path"
)

type ProviderCredentials struct {
	Client *Client
}

// Get - GET /orgs/:orgName/providers/:name/credentials
func (s ProviderCredentials) Get(providerName string) (*json.RawMessage, error) {
	res, err := s.Client.Do(http.MethodGet, path.Join("providers", providerName, "credentials"), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.Json[json.RawMessage](res)
}

// Update - PUT /orgs/:orgName/providers/:name/credentials
func (s ProviderCredentials) Update(providerName string, credentials interface{}) (*json.RawMessage, error) {
	rawPayload, _ := json.Marshal(credentials)
	endpoint := path.Join("providers", providerName, "credentials")
	res, err := s.Client.Do(http.MethodPut, endpoint, nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	return response.Json[json.RawMessage](res)
}
