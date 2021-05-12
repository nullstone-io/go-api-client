package api

import (
	"encoding/json"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"path"
	"strconv"
)

type Stacks struct {
	Client *Client
}

// List - GET /orgs/:orgName/stacks
func (s Stacks) List() ([]*types.Stack, error) {
	res, err := s.Client.Do(http.MethodGet, path.Join("stacks"), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var stacks []*types.Stack
	if err := s.Client.ReadJsonResponse(res, &stacks); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return stacks, nil
}

// Get - GET /orgs/:orgName/stacks/:name
func (s Stacks) Get(stackName string) (*types.Stack, error) {
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

// Create - POST /orgs/:orgName/stacks
func (s Stacks) Create(stack *types.Stack) (*types.Stack, error) {
	rawPayload, _ := json.Marshal(stack)
	res, err := s.Client.Do(http.MethodPost, path.Join("stacks"), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updatedStack types.Stack
	if err := s.Client.ReadJsonResponse(res, &updatedStack); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedStack, nil
}

// Update - PUT/PATCH /orgs/:orgName/envs/:id
func (s Stacks) Update(stackId int, stack *types.Stack) (*types.Stack, error) {
	rawPayload, _ := json.Marshal(stack)
	endpoint := path.Join("envs", strconv.Itoa(stackId))
	res, err := s.Client.Do(http.MethodPut, endpoint, nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updatedStack types.Stack
	if err := s.Client.ReadJsonResponse(res, &updatedStack); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedStack, nil
}

// Upsert - PUT/PATCH /orgs/:orgName/stacks/:name
func (s Stacks) Upsert(stackName string, stack *types.Stack) (*types.Stack, error) {
	rawPayload, _ := json.Marshal(stack)
	endpoint := path.Join("stacks", stackName)
	res, err := s.Client.Do(http.MethodPut, endpoint, nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updatedStack types.Stack
	if err := s.Client.ReadJsonResponse(res, &updatedStack); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedStack, nil
}

// Destroy - DELETE /orgs/:orgName/stacks/:name
func (s Stacks) Destroy(stackName string) (bool, error) {
	res, err := s.Client.Do(http.MethodDelete, path.Join("stacks", stackName), nil, nil, nil)
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
