package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type StacksByName struct {
	Client *Client
}

func (s StacksByName) basePath() string {
	return "stacks"
}

func (s StacksByName) stackPath(stackName string) string {
	return fmt.Sprintf("stacks/%s", stackName)
}

// List - GET /orgs/:orgName/stacks
func (s StacksByName) List() ([]*types.Stack, error) {
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

// Get - GET /orgs/:orgName/stacks/:name
func (s StacksByName) Get(stackName string) (*types.Stack, error) {
	res, err := s.Client.Do(http.MethodGet, s.stackPath(stackName), nil, nil, nil)
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

// Upsert - PUT/PATCH /orgs/:orgName/stacks/:name
func (s StacksByName) Upsert(stackName string, stack *types.Stack) (*types.Stack, error) {
	rawPayload, _ := json.Marshal(stack)
	res, err := s.Client.Do(http.MethodPut, s.stackPath(stackName), nil, nil, json.RawMessage(rawPayload))
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
