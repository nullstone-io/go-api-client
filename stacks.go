package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type Stacks struct {
	Client *Client
}

func (s Stacks) basePath() string {
	return "stacks"
}

func (s Stacks) stackPath(stackId int64) string {
	return fmt.Sprintf("stacks/%d", stackId)
}

// List - GET /orgs/:orgName/stacks
func (s Stacks) List() ([]*types.Stack, error) {
	res, err := s.Client.Do(http.MethodGet, s.basePath(), nil, nil, nil)
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

// Get - GET /orgs/:orgName/stacks/:id
func (s Stacks) Get(stackId int64) (*types.Stack, error) {
	res, err := s.Client.Do(http.MethodGet, s.stackPath(stackId), nil, nil, nil)
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
	res, err := s.Client.Do(http.MethodPost, s.basePath(), nil, nil, json.RawMessage(rawPayload))
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

// Update - PUT/PATCH /orgs/:orgName/stacks/:id
func (s Stacks) Update(stackId int64, stack *types.Stack) (*types.Stack, error) {
	rawPayload, _ := json.Marshal(stack)
	res, err := s.Client.Do(http.MethodPut, s.stackPath(stackId), nil, nil, json.RawMessage(rawPayload))
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

// Destroy - DELETE /orgs/:orgName/stacks/:id
func (s Stacks) Destroy(stackId int64) (bool, error) {
	res, err := s.Client.Do(http.MethodDelete, s.stackPath(stackId), nil, nil, nil)
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
