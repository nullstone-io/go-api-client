package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type Domains struct {
	Client *Client
}

func (s Domains) globalPath() string {
	return fmt.Sprintf("orgs/%s/domains", s.Client.Config.OrgName)
}

func (s Domains) basePath(stackId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/domains", s.Client.Config.OrgName, stackId)
}

func (s Domains) domainPath(stackId, domainId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/domains/%d", s.Client.Config.OrgName, stackId, domainId)
}

// List - GET /orgs/:orgName/domains
func (s Domains) List() ([]types.Domain, error) {
	res, err := s.Client.Do(http.MethodGet, s.globalPath(), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var domains []types.Domain
	if err := response.ReadJson(res, &domains); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return domains, nil
}

// Get - GET /orgs/:orgName/stacks/:stackId/domains/:id
func (s Domains) Get(stackId, domainId int64) (*types.Domain, error) {
	res, err := s.Client.Do(http.MethodGet, s.domainPath(stackId, domainId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var domain types.Domain
	if err := response.ReadJson(res, &domain); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &domain, nil
}

// Create - POST /orgs/:orgName/stacks/:stackId/domains
func (s Domains) Create(stackId int64, domain *types.Domain) (*types.Domain, error) {
	rawPayload, _ := json.Marshal(domain)
	res, err := s.Client.Do(http.MethodPost, s.basePath(stackId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updatedDomain types.Domain
	if err := response.ReadJson(res, &updatedDomain); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedDomain, nil
}

// Update - PUT/PATCH /orgs/:orgName/stacks/:stackId/domains/:id
func (s Domains) Update(stackId, domainId int64, domain *types.Domain) (*types.Domain, error) {
	rawPayload, _ := json.Marshal(domain)
	res, err := s.Client.Do(http.MethodPut, s.domainPath(stackId, domainId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updatedDomain types.Domain
	if err := response.ReadJson(res, &updatedDomain); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedDomain, nil
}

// Destroy - DELETE /orgs/:orgName/stacks/:stackId/domains/:id
func (s Domains) Destroy(stackId, domainId int64) (bool, error) {
	res, err := s.Client.Do(http.MethodDelete, s.domainPath(stackId, domainId), nil, nil, nil)
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
