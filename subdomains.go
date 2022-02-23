package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type Subdomains struct {
	Client *Client
}

func (s Subdomains) globalPath() string {
	return fmt.Sprintf("orgs/%s/subdomains", s.Client.Config.OrgName)
}

func (s Subdomains) globalSubdomainPath(subdomainId int64) string {
	return fmt.Sprintf("orgs/%s/subdomains/%d", s.Client.Config.OrgName, subdomainId)
}

func (s Subdomains) basePath(stackId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/subdomains", s.Client.Config.OrgName, stackId)
}

func (s Subdomains) subdomainPath(stackId, subdomainId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/subdomains/%d", s.Client.Config.OrgName, stackId, subdomainId)
}

// List - GET /orgs/:orgName/subdomains
func (s Subdomains) List() ([]types.Subdomain, error) {
	res, err := s.Client.Do(http.MethodGet, s.globalPath(), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var subdomains []types.Subdomain
	if err := response.ReadJson(res, &subdomains); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return subdomains, nil
}

// Get - GET /orgs/:orgName/subdomains/:id
func (s Subdomains) GlobalGet(subdomainId int64) (*types.Subdomain, error) {
	res, err := s.Client.Do(http.MethodGet, s.globalSubdomainPath(subdomainId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var subdomain types.Subdomain
	if err := response.ReadJson(res, &subdomain); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &subdomain, nil
}

// Get - GET /orgs/:orgName/stacks/:stackId/subdomains/:id
func (s Subdomains) Get(stackId, subdomainId int64) (*types.Subdomain, error) {
	res, err := s.Client.Do(http.MethodGet, s.subdomainPath(stackId, subdomainId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var subdomain types.Subdomain
	if err := response.ReadJson(res, &subdomain); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &subdomain, nil
}

// Create - POST /orgs/:orgName/stacks/:stackId/subdomains
func (s Subdomains) Create(stackId int64, subdomain *types.Subdomain) (*types.Subdomain, error) {
	rawPayload, _ := json.Marshal(subdomain)
	res, err := s.Client.Do(http.MethodPost, s.basePath(stackId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updatedDomain types.Subdomain
	if err := response.ReadJson(res, &updatedDomain); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedDomain, nil
}

// Update - PUT/PATCH /orgs/:orgName/stacks/:stackId/subdomains/:id
func (s Subdomains) Update(stackId, subdomainId int64, subdomain *types.Subdomain) (*types.Subdomain, error) {
	rawPayload, _ := json.Marshal(subdomain)
	res, err := s.Client.Do(http.MethodPut, s.subdomainPath(stackId, subdomainId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updatedDomain types.Subdomain
	if err := response.ReadJson(res, &updatedDomain); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedDomain, nil
}

// Destroy - DELETE /orgs/:orgName/stacks/:stackId/subdomains/:id
func (s Subdomains) Destroy(stackId, subdomainId int64) (bool, error) {
	res, err := s.Client.Do(http.MethodDelete, s.subdomainPath(stackId, subdomainId), nil, nil, nil)
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
