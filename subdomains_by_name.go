package api

import (
	"gopkg.in/nullstone-io/go-api-client.v0/response"
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
	return response.Json[types.Subdomain](res)
}
