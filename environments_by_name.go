package api

import (
	"encoding/json"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"path"
)

type EnvironmentsByName struct {
	Client *Client
}

// List - GET /orgs/:orgName/stacks/:stackName/envs
func (s EnvironmentsByName) List(stackName string) ([]*types.Environment, error) {
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

// Get - GET /orgs/:orgName/stacks/:stack_name/envs/:name
func (s EnvironmentsByName) Get(stackName, envName string) (*types.Environment, error) {
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

// Upsert - PUT/PATCH /orgs/:orgName/stacks/:stack_name/envs/:name
func (s EnvironmentsByName) Upsert(stackName, envName string, env *types.Environment) (*types.Environment, error) {
	rawPayload, _ := json.Marshal(env)
	res, err := s.Client.Do(http.MethodPut, path.Join("stacks", stackName, "envs", envName), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updatedEnv types.Environment
	if err := s.Client.ReadJsonResponse(res, &updatedEnv); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedEnv, nil
}
