package api

import (
	"encoding/json"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"path"
	"strconv"
)

type Subdomains struct {
	Client *Client
}

// List - GET /orgs/:orgName/subdomains
func (s Subdomains) List() ([]types.Subdomain, error) {
	res, err := s.Client.Do(http.MethodGet, path.Join("subdomains"), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var subdomains []types.Subdomain
	if err := s.Client.ReadJsonResponse(res, &subdomains); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return subdomains, nil
}

// GetByName - GET /orgs/:orgName/stacks/:stackName/subdomains/:name
func (s Subdomains) GetByName(stackName string, subdomainName string) (*types.Subdomain, error) {
	res, err := s.Client.Do(http.MethodGet, path.Join("stacks", stackName, "subdomains", subdomainName), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var subdomain types.Subdomain
	if err := s.Client.ReadJsonResponse(res, &subdomain); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &subdomain, nil
}

// Get - GET /orgs/:orgName/subdomains/:id
func (s Subdomains) Get(subdomainId int) (*types.Subdomain, error) {
	res, err := s.Client.Do(http.MethodGet, path.Join("subdomains", strconv.Itoa(subdomainId)), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var subdomain types.Subdomain
	if err := s.Client.ReadJsonResponse(res, &subdomain); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &subdomain, nil
}

// Create - POST /orgs/:orgName/subdomains
func (s Subdomains) Create(subdomain *types.Subdomain) (*types.Subdomain, error) {
	rawPayload, _ := json.Marshal(subdomain)
	res, err := s.Client.Do(http.MethodPost, path.Join("subdomains"), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updatedDomain types.Subdomain
	if err := s.Client.ReadJsonResponse(res, &updatedDomain); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedDomain, nil
}

// Update - PUT/PATCH /orgs/:orgName/subdomains/:id
func (s Subdomains) Update(subdomainId int, subdomain *types.Subdomain) (*types.Subdomain, error) {
	rawPayload, _ := json.Marshal(subdomain)
	endpoint := path.Join("subdomains", strconv.Itoa(subdomainId))
	res, err := s.Client.Do(http.MethodPut, endpoint, nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updatedDomain types.Subdomain
	if err := s.Client.ReadJsonResponse(res, &updatedDomain); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedDomain, nil
}

// Destroy - DELETE /orgs/:orgName/subdomains/:id
func (s Subdomains) Destroy(subdomainId int) (bool, error) {
	res, err := s.Client.Do(http.MethodDelete, path.Join("subdomains", strconv.Itoa(subdomainId)), nil, nil, nil)
	if err != nil {
		return false, err
	}
	if err := s.Client.VerifyResponse(res); IsNotFoundError(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
