package api

import (
	"encoding/json"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"path"
)

type Providers struct {
	Client *Client
}

// List - GET /orgs/:orgName/providers
func (s Providers) List() ([]*types.Provider, error) {
	res, err := s.Client.Do(http.MethodGet, path.Join("providers"), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.JsonArray[*types.Provider](res)
}

// Get - GET /orgs/:orgName/providers/:name
func (s Providers) Get(providerName string) (*types.Provider, error) {
	res, err := s.Client.Do(http.MethodGet, path.Join("providers", providerName), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.Json[types.Provider](res)
}

// Create - POST /orgs/:orgName/providers
func (s Providers) Create(provider *types.Provider) (*types.Provider, error) {
	rawPayload, _ := json.Marshal(provider)
	res, err := s.Client.Do(http.MethodPost, path.Join("providers"), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	return response.Json[types.Provider](res)
}

// Update - PUT/PATCH /orgs/:orgName/providers/:name
func (s Providers) Update(providerName string, provider *types.Provider) (*types.Provider, error) {
	rawPayload, _ := json.Marshal(provider)
	endpoint := path.Join("providers", providerName)
	res, err := s.Client.Do(http.MethodPut, endpoint, nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	return response.Json[types.Provider](res)
}

// Destroy - DELETE /orgs/:orgName/providers/:name
func (s Providers) Destroy(providerName string) (bool, error) {
	res, err := s.Client.Do(http.MethodDelete, path.Join("providers", providerName), nil, nil, nil)
	if err != nil {
		return false, err
	}
	if err := response.Verify(res); response.IsNotFoundError(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
