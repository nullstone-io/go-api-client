package api

import (
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

	var providers []*types.Provider
	if err := s.Client.ReadJsonResponse(res, &providers); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return providers, nil
}

// Get - GET /orgs/:orgName/providers/:name
func (s Providers) Get(providerName string) (*types.Provider, error) {
	res, err := s.Client.Do(http.MethodGet, path.Join("providers", providerName), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var provider types.Provider
	if err := s.Client.ReadJsonResponse(res, &provider); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &provider, nil
}
