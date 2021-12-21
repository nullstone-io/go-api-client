package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
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
	return response.JsonArray[*types.Stack](res)
}

// Get - GET /orgs/:orgName/stacks/:id
func (s Stacks) Get(stackId int64) (*types.Stack, error) {
	res, err := s.Client.Do(http.MethodGet, s.stackPath(stackId), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.Json[types.Stack](res)
}

// Create - POST /orgs/:orgName/stacks
func (s Stacks) Create(stack *types.Stack) (*types.Stack, error) {
	rawPayload, _ := json.Marshal(stack)
	res, err := s.Client.Do(http.MethodPost, s.basePath(), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	return response.Json[types.Stack](res)
}

// Update - PUT/PATCH /orgs/:orgName/stacks/:id
func (s Stacks) Update(stackId int64, stack *types.Stack) (*types.Stack, error) {
	rawPayload, _ := json.Marshal(stack)
	res, err := s.Client.Do(http.MethodPut, s.stackPath(stackId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	return response.Json[types.Stack](res)
}

// Destroy - DELETE /orgs/:orgName/stacks/:id
func (s Stacks) Destroy(stackId int64) (bool, error) {
	res, err := s.Client.Do(http.MethodDelete, s.stackPath(stackId), nil, nil, nil)
	if err != nil {
		return false, err
	}
	if err := response.Verify(res); response.IsNotFoundError(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
