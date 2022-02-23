package api

import (
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type AutogenSubdomain struct {
	Client *Client
}

func (a AutogenSubdomain) path(stackId, subdomainId, envId int64) string {
	return fmt.Sprintf("orgs/%s/subdomains/%d/envs/%d/autogen_subdomain", a.Client.Config.OrgName, subdomainId, envId)
}

// Get - GET /orgs/:orgName/subdomains/:subdomainId/envs/:envId/autogen_subdomain
func (a AutogenSubdomain) Get(stackId, subdomainId, envId int64) (*types.AutogenSubdomain, error) {
	res, err := a.Client.Do(http.MethodGet, a.path(stackId, subdomainId, envId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var autogenSubdomain types.AutogenSubdomain
	if err := response.ReadJson(res, &autogenSubdomain); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &autogenSubdomain, nil
}

// Create - POST /orgs/:orgName/subdomains/:subdomainId/envs/:envId/autogen_subdomain
func (a AutogenSubdomain) Create(stackId, subdomainId, envId int64) (*types.AutogenSubdomain, error) {
	res, err := a.Client.Do(http.MethodPost, a.path(stackId, subdomainId, envId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var autogenSubdomain types.AutogenSubdomain
	if err := response.ReadJson(res, &autogenSubdomain); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &autogenSubdomain, nil
}

// Destroy - DELETE /orgs/:orgName/subdomains/:subdomainId/envs/:envId/autogen_subdomain
func (a AutogenSubdomain) Destroy(stackId, subdomainId, envId int64) (bool, error) {
	res, err := a.Client.Do(http.MethodDelete, a.path(stackId, subdomainId, envId), nil, nil, nil)
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
