package api

import (
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"path"
)

type SubdomainsByName struct {
	Client *Client
}

// Get - GET /orgs/:orgName/stacks/:stackName/subdomains/:name
func (s SubdomainsByName) Get(stackName string, subdomainName string) (*types.Subdomain, error) {
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
