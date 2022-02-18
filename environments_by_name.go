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

func (s EnvironmentsByName) path(stackName, envName string) string {
	return path.Join("orgs", s.Client.Config.OrgName, "stacks_by_name", stackName, "envs", envName)
}

// Get - GET /orgs/:orgName/stacks_by_name/:stack_name/envs/:name
func (s EnvironmentsByName) Get(stackName, envName string) (*types.Environment, error) {
	res, err := s.Client.Do(http.MethodGet, s.path(stackName, envName), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var env types.Environment
	if err := response.ReadJson(res, &env); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &env, nil
}

// Upsert - PUT/PATCH /orgs/:orgName/stacks_by_name/:stack_name/envs/:name
func (s EnvironmentsByName) Upsert(stackName, envName string, env *types.Environment) (*types.Environment, error) {
	rawPayload, _ := json.Marshal(env)
	res, err := s.Client.Do(http.MethodPut, s.path(stackName, envName), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updatedEnv types.Environment
	if err := response.ReadJson(res, &updatedEnv); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedEnv, nil
}
