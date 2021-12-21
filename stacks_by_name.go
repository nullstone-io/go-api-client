package api

import (
	"encoding/json"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"path"
)

type StacksByName struct {
	Client *Client
}

// Get - GET /orgs/:orgName/stacks_by_name/:name
func (s StacksByName) Get(stackName string) (*types.Stack, error) {
	res, err := s.Client.Do(http.MethodGet, path.Join("stacks_by_name", stackName), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.Json[types.Stack](res)
}

// Upsert - PUT/PATCH /orgs/:orgName/stacks_by_name/:name
func (s StacksByName) Upsert(stackName string, stack *types.Stack) (*types.Stack, error) {
	rawPayload, _ := json.Marshal(stack)
	endpoint := path.Join("stacks_by_name", stackName)
	res, err := s.Client.Do(http.MethodPut, endpoint, nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	return response.Json[types.Stack](res)
}
