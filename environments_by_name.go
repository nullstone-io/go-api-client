package api

import (
	"encoding/json"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"path"
)

type EnvironmentsByName struct {
	Client *Client
}

// Get - GET /orgs/:orgName/stacks_by_name/:stack_name/envs/:name
func (s EnvironmentsByName) Get(stackName, envName string) (*types.Environment, error) {
	res, err := s.Client.Do(http.MethodGet, path.Join("stacks_by_name", stackName, "envs", envName), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.Json[types.Environment](res)
}

// Upsert - PUT/PATCH /orgs/:orgName/stacks_by_name/:stack_name/envs/:name
func (s EnvironmentsByName) Upsert(stackName, envName string, env *types.Environment) (*types.Environment, error) {
	rawPayload, _ := json.Marshal(env)
	res, err := s.Client.Do(http.MethodPut, path.Join("stacks_by_name", stackName, "envs", envName), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	return response.Json[types.Environment](res)
}
