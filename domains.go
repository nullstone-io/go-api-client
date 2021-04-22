package api

import (
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"path"
)

type Domains struct {
	Client *Client
}

func (s Domains) Get(stackName string, domainName string) (*types.Domain, error) {
	res, err := s.Client.Do(http.MethodGet, path.Join("stacks", stackName, "domains", domainName), nil, nil, nil)
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
