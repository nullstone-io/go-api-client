package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type Domains struct {
	Client *Client
}

func (s Domains) basePath() string {
	return fmt.Sprintf("domains")
}

func (s Domains) domainPath(domainId int64) string {
	return fmt.Sprintf("domains/%d", domainId)
}

// List - GET /orgs/:orgName/domains
func (s Domains) List() ([]types.Domain, error) {
	res, err := s.Client.Do(http.MethodGet, s.basePath(), nil, nil, nil)
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
func (s Domains) Get(domainId int64) (*types.Domain, error) {
	res, err := s.Client.Do(http.MethodGet, s.domainPath(domainId), nil, nil, nil)
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
	res, err := s.Client.Do(http.MethodPost, s.basePath(), nil, nil, json.RawMessage(rawPayload))
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
func (s Domains) Update(domainId int64, domain *types.Domain) (*types.Domain, error) {
	rawPayload, _ := json.Marshal(domain)
	res, err := s.Client.Do(http.MethodPut, s.domainPath(domainId), nil, nil, json.RawMessage(rawPayload))
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
func (s Domains) Destroy(domainId int64) (bool, error) {
	res, err := s.Client.Do(http.MethodDelete, s.domainPath(domainId), nil, nil, nil)
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
