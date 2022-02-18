package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"net/http"
)

type ProviderCredentials struct {
	Client *Client
}

func (s ProviderCredentials) path(providerName string) string {
	return fmt.Sprintf("orgs/%s/providers/%s/credentials", s.Client.Config.OrgName, providerName)
}

// Get - GET /orgs/:orgName/providers/:name/credentials
func (s ProviderCredentials) Get(providerName string) (*json.RawMessage, error) {
	res, err := s.Client.Do(http.MethodGet, s.path(providerName), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var creds json.RawMessage
	if err := response.ReadJson(res, &creds); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &creds, nil
}

// Update - PUT /orgs/:orgName/providers/:name/credentials
func (s ProviderCredentials) Update(providerName string, credentials interface{}) (*json.RawMessage, error) {
	rawPayload, _ := json.Marshal(credentials)
	res, err := s.Client.Do(http.MethodPut, s.path(providerName), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updatedCreds json.RawMessage
	if err := response.ReadJson(res, &updatedCreds); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedCreds, nil
}
