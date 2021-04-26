package api

import (
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"path"
	"strconv"
)

type AutogenSubdomains struct {
	Client *Client
}

// GET /orgs/:orgName/subdomains/:subdomainId/envs/:envName/autogen_subdomains
func (a AutogenSubdomains) Get(subdomainId int, envName string) (*types.AutogenSubdomain, error) {
	res, err := a.Client.Do(http.MethodGet, path.Join("subdomains", strconv.Itoa(subdomainId), "envs", envName, "autogen_subdomain"), nil, nil, nil)
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

// POST /orgs/:orgName/subdomains/:subdomainId/envs/:envName/autogen_subdomain
func (a AutogenSubdomains) Create(subdomainId int, envName string) (*types.AutogenSubdomain, error) {
	res, err := a.Client.Do(http.MethodPost, path.Join("subdomains", strconv.Itoa(subdomainId), "envs", envName, "autogen_subdomain"), nil, nil, nil)
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

// DELETE /orgs/:orgName/subdomains/:subdomainId/envs/:envName/autogen_subdomains
func (a AutogenSubdomains) Destroy(subdomainId int, envName string) (bool, error) {
	res, err := a.Client.Do(http.MethodDelete, path.Join("subdomains", strconv.Itoa(subdomainId), "envs", envName, "autogen_subdomain"), nil, nil, nil)
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
