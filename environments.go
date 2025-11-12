package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

type Environments struct {
	Client *Client
}

func (s Environments) orgPath() string {
	return fmt.Sprintf("orgs/%s/envs", s.Client.Config.OrgName)
}

func (s Environments) basePath(stackId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/envs", s.Client.Config.OrgName, stackId)
}

func (s Environments) envPath(stackId, envId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/envs/%d", s.Client.Config.OrgName, stackId, envId)
}

// GlobalList - GET /orgs/:orgName/envs
func (s Environments) GlobalList(ctx context.Context, envTypes []types.EnvironmentType, page, limit int, search string) (*Paginated[types.EnvironmentWithStack], error) {
	q := url.Values{}
	if page > 0 {
		q.Set("page", strconv.Itoa(page))
	}
	if limit > 0 {
		q.Set("limit", strconv.Itoa(limit))
	}
	if len(envTypes) > 0 {
		envTypeStrings := make([]string, 0)
		for _, envType := range envTypes {
			envTypeStrings = append(envTypeStrings, string(envType))
		}
		q.Set("type", strings.Join(envTypeStrings, ","))
	}
	if search != "" {
		q.Set("search", search)
	}
	res, err := s.Client.Do(ctx, http.MethodGet, s.orgPath(), q, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[Paginated[types.EnvironmentWithStack]](res)
}

// List - GET /orgs/:orgName/stacks/:stackId/envs
func (s Environments) List(ctx context.Context, stackId int64) ([]*types.Environment, error) {
	res, err := s.Client.Do(ctx, http.MethodGet, s.basePath(stackId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var envs []*types.Environment
	if err := response.ReadJson(res, &envs); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return envs, nil
}

// Get - GET /orgs/:orgName/stacks/:stack_id/envs/:id
func (s Environments) Get(ctx context.Context, stackId, envId int64, includeArchived bool) (*types.Environment, error) {
	q := url.Values{
		"include_archived": []string{strconv.FormatBool(includeArchived)},
	}
	res, err := s.Client.Do(ctx, http.MethodGet, s.envPath(stackId, envId), q, nil, nil)
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

// Create - POST /orgs/:orgName/stacks/:stack_id/envs
func (s Environments) Create(ctx context.Context, stackId int64, env *types.Environment) (*types.Environment, error) {
	rawPayload, _ := json.Marshal(env)
	res, err := s.Client.Do(ctx, http.MethodPost, s.basePath(stackId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	return response.ReadJsonPtr[types.Environment](res)
}

// Update - PUT/PATCH /orgs/:orgName/stacks/:stack_id/envs/:id
func (s Environments) Update(ctx context.Context, stackId, envId int64, env *types.Environment) (*types.Environment, error) {
	rawPayload, _ := json.Marshal(env)
	res, err := s.Client.Do(ctx, http.MethodPut, s.envPath(stackId, envId), nil, nil, json.RawMessage(rawPayload))
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

// Destroy - DELETE /orgs/:orgName/stacks/:stack_id/envs/:id
func (s Environments) Destroy(ctx context.Context, stackId, envId int64) (bool, error) {
	res, err := s.Client.Do(ctx, http.MethodDelete, s.envPath(stackId, envId), nil, nil, nil)
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
