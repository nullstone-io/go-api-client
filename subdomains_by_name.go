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

func (s SubdomainsByName) subdomainPath(stackName, subdomainName string) string {
	return path.Join("orgs", s.Client.Config.OrgName, "stacks", stackName, "subdomains", subdomainName)
}

// Get - GET /orgs/:orgName/stacks/:stackName/subdomains/:name
func (s SubdomainsByName) Get(stackName string, subdomainName string) (*types.Subdomain, error) {
	res, err := s.Client.Do(http.MethodGet, s.subdomainPath(stackName, subdomainName), nil, nil, nil)
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
