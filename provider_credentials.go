package api

import (
	"encoding/json"
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
	if err := s.Client.ReadJsonResponse(res, &creds); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &creds, nil
}
