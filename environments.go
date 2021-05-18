package api

import (
	"encoding/json"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"path"
	"strconv"
)

type Environments struct {
	Client *Client
}

// List - GET /orgs/:orgName/stacks/:stackId/envs
func (s Environments) List(stackId int64) ([]*types.Environment, error) {
	res, err := s.Client.Do(http.MethodGet, path.Join("stacks", strconv.FormatInt(stackId, 10), "envs"), nil, nil, nil)
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

// ListByName - GET /orgs/:orgName/stacks/:stackName/envs
func (s Environments) ListByName(stackName string) ([]*types.Environment, error) {
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

// Get - GET /orgs/:orgName/envs/:id
func (s Environments) Get(envId int64) (*types.Environment, error) {
	res, err := s.Client.Do(http.MethodGet, path.Join("envs", strconv.FormatInt(envId, 10)), nil, nil, nil)
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

// Create - POST /orgs/:orgName/stacks/:stackName/envs
func (s Environments) Create(stackName string, env *types.Environment) (*types.Environment, error) {
	rawPayload, _ := json.Marshal(env)
	res, err := s.Client.Do(http.MethodPost, path.Join("stacks", stackName, "envs"), nil, nil, json.RawMessage(rawPayload))
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

// Update - PUT/PATCH /orgs/:orgName/envs/:id
func (s Environments) Update(envId int, env *types.Environment) (*types.Environment, error) {
	rawPayload, _ := json.Marshal(env)
	endpoint := path.Join("envs", strconv.Itoa(envId))
	res, err := s.Client.Do(http.MethodPut, endpoint, nil, nil, json.RawMessage(rawPayload))
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

// Upsert - PUT/PATCH /orgs/:orgName/stacks/:stackName/envs/:name
func (s Environments) Upsert(stackName, envName string, env *types.Environment) (*types.Environment, error) {
	rawPayload, _ := json.Marshal(env)
	endpoint := path.Join("stacks", stackName, "envs", envName)
	res, err := s.Client.Do(http.MethodPut, endpoint, nil, nil, json.RawMessage(rawPayload))
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

// Destroy - DELETE /orgs/:orgName/stacks/:stackName/envs/:name
func (s Environments) Destroy(stackName, envName string) (bool, error) {
	res, err := s.Client.Do(http.MethodDelete, path.Join("stacks", stackName, "envs", envName), nil, nil, nil)
	if err != nil {
		return false, err
	}
	if err := s.Client.VerifyResponse(res); IsNotFoundError(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
