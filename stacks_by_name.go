package api

import (
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"path"
)

type StacksByName struct {
	Client *Client
}

// Get - GET /orgs/:orgName/stacks/:name
func (s StacksByName) Get(stackName string) (*types.Stack, error) {
	res, err := s.Client.Do(http.MethodGet, path.Join("stacks", stackName), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var stack types.Stack
	if err := s.Client.ReadJsonResponse(res, &stack); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &stack, nil
}
