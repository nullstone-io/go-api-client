package api

import (
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"path"
)

type Environments struct {
	Client *Client
}

// List - GET /orgs/:orgName/stacks/:stackName/envs
func (s Environments) List(stackName string) ([]*types.Environment, error) {
	res, err := s.Client.Do(http.MethodGet, path.Join("stacks", stackName, "envs"), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var envs []*types.Environment
	if err := s.Client.ReadJsonResponse(res, &envs); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return envs, nil
}

// Get - GET /orgs/:orgName/stacks/:stackName/envs/:name
func (s Environments) Get(stackName, envName string) (*types.Environment, error) {
	res, err := s.Client.Do(http.MethodGet, path.Join("stacks", stackName, "envs", envName), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var env types.Environment
	if err := s.Client.ReadJsonResponse(res, &env); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &env, nil
}