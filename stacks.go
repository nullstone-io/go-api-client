package api

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"net/url"
	"strconv"
)

type Stacks struct {
	Client *Client
}

func (s Stacks) basePath() string {
	return fmt.Sprintf("orgs/%s/stacks", s.Client.Config.OrgName)
}

func (s Stacks) stackPath(stackId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d", s.Client.Config.OrgName, stackId)
}

// List - GET /orgs/:orgName/stacks
func (s Stacks) List(ctx context.Context) ([]*types.Stack, error) {
	res, err := s.Client.Do(ctx, http.MethodGet, s.basePath(), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var stacks []*types.Stack
	if err := response.ReadJson(res, &stacks); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return stacks, nil
}

// Get - GET /orgs/:orgName/stacks/:id
func (s Stacks) Get(ctx context.Context, stackId int64, includeArchived bool) (*types.Stack, error) {
	q := url.Values{
		"include_archived": []string{strconv.FormatBool(includeArchived)},
	}
	res, err := s.Client.Do(ctx, http.MethodGet, s.stackPath(stackId), q, nil, nil)
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

// Create - POST /orgs/:orgName/stacks
func (s Stacks) Create(ctx context.Context, stack *types.Stack) (*types.Stack, error) {
	rawPayload, _ := json.Marshal(stack)
	res, err := s.Client.Do(ctx, http.MethodPost, s.basePath(), nil, nil, json.RawMessage(rawPayload))
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

// Update - PUT/PATCH /orgs/:orgName/stacks/:id
func (s Stacks) Update(ctx context.Context, stackId int64, stack *types.Stack) (*types.Stack, error) {
	rawPayload, _ := json.Marshal(stack)
	res, err := s.Client.Do(ctx, http.MethodPut, s.stackPath(stackId), nil, nil, json.RawMessage(rawPayload))
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

// Destroy - DELETE /orgs/:orgName/stacks/:id
func (s Stacks) Destroy(ctx context.Context, stackId int64) (bool, error) {
	res, err := s.Client.Do(ctx, http.MethodDelete, s.stackPath(stackId), nil, nil, nil)
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
