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

	var creds json.RawMessage
	if err := s.Client.ReadJsonResponse(res, &creds); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &creds, nil
}

// Update - PUT /orgs/:orgName/providers/:name/credentials
func (s ProviderCredentials) Update(providerName string, credentials interface{}) (*json.RawMessage, error) {
	rawPayload, _ := json.Marshal(credentials)
	endpoint := path.Join("providers", providerName, "credentials")
	res, err := s.Client.Do(http.MethodPut, endpoint, nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updatedCreds json.RawMessage
	if err := s.Client.ReadJsonResponse(res, &updatedCreds); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedCreds, nil
}
