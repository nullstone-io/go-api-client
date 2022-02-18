package api

import (
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type DomainsByName struct {
	Client *Client
}

func (s DomainsByName) domainPath(stackName, domainName string) string {
	return fmt.Sprintf("orgs/%s/stacks/%s/domains/%s", s.Client.Config.OrgName, stackName, domainName)
}

// Get - GET /orgs/:orgName/stacks/:stackName/domains/:name
func (s DomainsByName) Get(stackName, domainName string) (*types.Domain, error) {
	res, err := s.Client.Do(http.MethodGet, s.domainPath(stackName, domainName), nil, nil, nil)
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
