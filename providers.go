package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type Providers struct {
	Client *Client
}

func (s Providers) basePath() string {
	return fmt.Sprintf("orgs/%s/providers", s.Client.Config.OrgName)
}

func (s Providers) providerPath(providerName string) string {
	return fmt.Sprintf("orgs/%s/providers/%s", s.Client.Config.OrgName, providerName)
}

// List - GET /orgs/:orgName/providers
func (s Providers) List() ([]*types.Provider, error) {
	res, err := s.Client.Do(http.MethodGet, s.basePath(), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var providers []*types.Provider
	if err := response.ReadJson(res, &providers); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return providers, nil
}

// Get - GET /orgs/:orgName/providers/:name
func (s Providers) Get(providerName string) (*types.Provider, error) {
	res, err := s.Client.Do(http.MethodGet, s.providerPath(providerName), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var provider types.Provider
	if err := response.ReadJson(res, &provider); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &provider, nil
}

// Create - POST /orgs/:orgName/providers
func (s Providers) Create(provider *types.Provider) (*types.Provider, error) {
	rawPayload, _ := json.Marshal(provider)
	res, err := s.Client.Do(http.MethodPost, s.basePath(), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updatedProvider types.Provider
	if err := response.ReadJson(res, &updatedProvider); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedProvider, nil
}

// Update - PUT/PATCH /orgs/:orgName/providers/:name
func (s Providers) Update(providerName string, provider *types.Provider) (*types.Provider, error) {
	rawPayload, _ := json.Marshal(provider)
	res, err := s.Client.Do(http.MethodPut, s.providerPath(providerName), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updatedProvider types.Provider
	if err := response.ReadJson(res, &updatedProvider); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedProvider, nil
}

// Destroy - DELETE /orgs/:orgName/providers/:name
func (s Providers) Destroy(providerName string) (bool, error) {
	res, err := s.Client.Do(http.MethodDelete, s.providerPath(providerName), nil, nil, nil)
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
