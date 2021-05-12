package api

import (
	"encoding/json"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"path"
	"strconv"
)

type Domains struct {
	Client *Client
}

// List - GET /orgs/:orgName/domains
func (s Domains) List() ([]types.Domain, error) {
	res, err := s.Client.Do(http.MethodGet, path.Join("domains"), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var domains []types.Domain
	if err := s.Client.ReadJsonResponse(res, &domains); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return domains, nil
}

// Get - GET /orgs/:orgName/domains/:id
func (s Domains) Get(domainId int) (*types.Domain, error) {
	res, err := s.Client.Do(http.MethodGet, path.Join("domains", strconv.Itoa(domainId)), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var domain types.Domain
	if err := s.Client.ReadJsonResponse(res, &domain); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &domain, nil
}

// GetByName - GET /orgs/:orgName/stacks/:stackName/domains/:name
func (s Domains) GetByName(stackName, domainName string) (*types.Domain, error) {
	res, err := s.Client.Do(http.MethodGet, path.Join("stacks", stackName, "domains", domainName), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var domain types.Domain
	if err := s.Client.ReadJsonResponse(res, &domain); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &domain, nil
}

// Create - POST /orgs/:orgName/domains
func (s Domains) Create(domain *types.Domain) (*types.Domain, error) {
	rawPayload, _ := json.Marshal(domain)
	res, err := s.Client.Do(http.MethodPost, path.Join("domains"), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updatedDomain types.Domain
	if err := s.Client.ReadJsonResponse(res, &updatedDomain); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedDomain, nil
}

// Update - PUT/PATCH /orgs/:orgName/domains/:id
func (s Domains) Update(domainId int, domain *types.Domain) (*types.Domain, error) {
	rawPayload, _ := json.Marshal(domain)
	endpoint := path.Join("domains", strconv.Itoa(domainId))
	res, err := s.Client.Do(http.MethodPut, endpoint, nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updatedDomain types.Domain
	if err := s.Client.ReadJsonResponse(res, &updatedDomain); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedDomain, nil
}

// Destroy - DELETE /orgs/:orgName/domains/:id
func (s Domains) Destroy(domainId int) (bool, error) {
	res, err := s.Client.Do(http.MethodDelete, path.Join("domains", strconv.Itoa(domainId)), nil, nil, nil)
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
