package api

import (
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type AutogenSubdomain struct {
	Client *Client
}

func (d AutogenSubdomain) path(subdomainId int64, envName string) string {
	return fmt.Sprintf("subdomains/%d/envs/%s/autogen_subdomain", subdomainId, envName)
}

// Get - GET /orgs/:orgName/subdomains/:subdomainId/envs/:envName/autogen_subdomain
func (a AutogenSubdomain) Get(subdomainId int64, envName string) (*types.AutogenSubdomain, error) {
	res, err := a.Client.Do(http.MethodGet, a.path(subdomainId, envName), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var autogenSubdomain types.AutogenSubdomain
	if err := a.Client.ReadJsonResponse(res, &autogenSubdomain); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &autogenSubdomain, nil
}

// Create - POST /orgs/:orgName/subdomains/:subdomainId/envs/:envName/autogen_subdomain
func (a AutogenSubdomain) Create(subdomainId int64, envName string) (*types.AutogenSubdomain, error) {
	res, err := a.Client.Do(http.MethodPost, a.path(subdomainId, envName), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var autogenSubdomain types.AutogenSubdomain
	if err := a.Client.ReadJsonResponse(res, &autogenSubdomain); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &autogenSubdomain, nil
}

// Destroy - DELETE /orgs/:orgName/subdomains/:subdomainId/envs/:envName/autogen_subdomain
func (a AutogenSubdomain) Destroy(subdomainId int64, envName string) (bool, error) {
	res, err := a.Client.Do(http.MethodDelete, a.path(subdomainId, envName), nil, nil, nil)
	if err != nil {
		return false, err
	}
	if err := a.Client.VerifyResponse(res); IsNotFoundError(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
