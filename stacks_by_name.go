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

func (s StacksByName) stackPath(stackName string) string {
	return path.Join("orgs", s.Client.Config.OrgName, "stacks_by_name", stackName)
}

// Get - GET /orgs/:orgName/stacks_by_name/:name
func (s StacksByName) Get(stackName string) (*types.Stack, error) {
	res, err := s.Client.Do(http.MethodGet, s.stackPath(stackName), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var stack types.Stack
	if err := response.ReadJson(res, &stack); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &stack, nil
}

// Upsert - PUT/PATCH /orgs/:orgName/stacks_by_name/:name
func (s StacksByName) Upsert(stackName string, stack *types.Stack) (*types.Stack, error) {
	rawPayload, _ := json.Marshal(stack)
	res, err := s.Client.Do(http.MethodPut, s.stackPath(stackName), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updatedStack types.Stack
	if err := response.ReadJson(res, &updatedStack); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedStack, nil
}
