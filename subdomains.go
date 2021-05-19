package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type Subdomains struct {
	Client *Client
}

func (s Subdomains) basePath() string {
	return fmt.Sprintf("subdomains")
}

func (s Subdomains) subdomainPath(subdomainId int64) string {
	return fmt.Sprintf("subdomains/%d", subdomainId)
}

// List - GET /orgs/:orgName/subdomains
func (s Subdomains) List() ([]types.Subdomain, error) {
	res, err := s.Client.Do(http.MethodGet, s.basePath(), nil, nil, nil)
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

// Get - GET /orgs/:orgName/subdomains/:id
func (s Subdomains) Get(subdomainId int64) (*types.Subdomain, error) {
	res, err := s.Client.Do(http.MethodGet, s.subdomainPath(subdomainId), nil, nil, nil)
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
	res, err := s.Client.Do(http.MethodPost, s.basePath(), nil, nil, json.RawMessage(rawPayload))
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
func (s Subdomains) Update(subdomainId int64, subdomain *types.Subdomain) (*types.Subdomain, error) {
	rawPayload, _ := json.Marshal(subdomain)
	res, err := s.Client.Do(http.MethodPut, s.subdomainPath(subdomainId), nil, nil, json.RawMessage(rawPayload))
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
func (s Subdomains) Destroy(subdomainId int64) (bool, error) {
	res, err := s.Client.Do(http.MethodDelete, s.subdomainPath(subdomainId), nil, nil, nil)
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
