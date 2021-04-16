package api

import (
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"path"
)

type AutogenSubdomains struct {
	Client *Client
}

// GET /orgs/:orgName/autogen_subdomains/:subdomainName
func (a AutogenSubdomains) Get(subdomainName string) (*types.AutogenSubdomain, error) {
	res, err := a.Client.Do(http.MethodGet, path.Join("autogen_subdomains", subdomainName), nil, nil)
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

// POST /orgs/:orgName/autogen_subdomains
func (a AutogenSubdomains) Create() (*types.AutogenSubdomain, error) {
	res, err := a.Client.Do(http.MethodPost, path.Join("autogen_subdomains"), nil, nil)
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

// DELETE /orgs/:orgName/autogen_subdomains
func (a AutogenSubdomains) Destroy(subdomainName string) (bool, error) {
	res, err := a.Client.Do(http.MethodDelete, path.Join("autogen_subdomains", subdomainName), nil, nil)
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
