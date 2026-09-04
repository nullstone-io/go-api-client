package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
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

func (s Environments) envActivityPath(stackId, envId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/envs/%d/activity", s.Client.Config.OrgName, stackId, envId)
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
// Returns active environments only.
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

// FindEnvironmentsInput narrows Find. Filters AND together; a zero value matches every
// active environment in the stack, which is what List returns.
type FindEnvironmentsInput struct {
	// Types matches any one of these environment types (OR).
	Types []types.EnvironmentType
	// Status selects active (the default) or archived environments.
	Status types.EnvStatus
	// IsProd matches on the prod flag; nil matches either.
	IsProd *bool
	// Search matches the name case-insensitively: as a whole-name pattern when it contains
	// * (any run of characters) or ? (one character), otherwise as a substring.
	Search string
	// Tags requires every key: a non-empty value must equal the environment's tag; an empty
	// value matches environments where the key is absent or present as "" — the way to find
	// environments nobody has tagged yet.
	Tags map[string]string
}

func (input FindEnvironmentsInput) query() url.Values {
	q := url.Values{}
	if len(input.Types) > 0 {
		envTypeStrings := make([]string, 0, len(input.Types))
		for _, envType := range input.Types {
			envTypeStrings = append(envTypeStrings, string(envType))
		}
		q.Set("type", strings.Join(envTypeStrings, ","))
	}
	if input.Status != "" {
		q.Set("status", string(input.Status))
	}
	if input.IsProd != nil {
		q.Set("is_prod", strconv.FormatBool(*input.IsProd))
	}
	if input.Search != "" {
		q.Set("search", input.Search)
	}
	// sorted so the same filters always produce the same URL
	keys := make([]string, 0, len(input.Tags))
	for key := range input.Tags {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		q.Add("tag", key+"="+input.Tags[key])
	}
	return q
}

// Find - GET /orgs/:orgName/stacks/:stackId/envs?type=&status=&is_prod=&search=&tag=KEY=VALUE
// Filtering happens server-side; List is Find with no filters.
func (s Environments) Find(ctx context.Context, stackId int64, input FindEnvironmentsInput) ([]*types.Environment, error) {
	res, err := s.Client.Do(ctx, http.MethodGet, s.basePath(stackId), input.query(), nil, nil)
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

// UpdateEnvironmentMetadataInput is a partial update of an environment's
// descriptive metadata. Every field is a pointer so a caller can update a single
// field without clearing the others.
type UpdateEnvironmentMetadataInput struct {
	// Description updates the environment description: nil leaves it untouched,
	// an empty string clears it, any other value sets it.
	Description *string `json:"description,omitempty"`
}

// ApplyTo merges the provided fields onto existing metadata, leaving untouched
// any field whose pointer is nil.
func (i UpdateEnvironmentMetadataInput) ApplyTo(existing types.EnvironmentMetadata) types.EnvironmentMetadata {
	if i.Description != nil {
		existing.Description = *i.Description
	}
	return existing
}

type UpdateEnvironmentInput struct {
	Name           *string               `json:"name,omitempty"`
	IsProd         *bool                 `json:"isProd,omitempty"`
	PipelineOrder  *int                  `json:"pipelineOrder,omitempty"`
	ProviderConfig *types.ProviderConfig `json:"providerConfig,omitempty"`
	// Metadata is a partial update; omitting it leaves the stored metadata unchanged.
	Metadata *UpdateEnvironmentMetadataInput `json:"metadata,omitempty"`
}

// Update - PUT/PATCH /orgs/:orgName/stacks/:stack_id/envs/:id
func (s Environments) Update(ctx context.Context, stackId, envId int64, input UpdateEnvironmentInput) (*types.Environment, error) {
	rawPayload, _ := json.Marshal(input)
	res, err := s.Client.Do(ctx, http.MethodPut, s.envPath(stackId, envId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.Environment](res)
}

// UpdateActivity - PUT /orgs/:orgName/stacks/:stack_id/envs/:id/activity
func (s Environments) UpdateActivity(ctx context.Context, stackId, envId int64) (*types.Environment, error) {
	res, err := s.Client.Do(ctx, http.MethodPut, s.envActivityPath(stackId, envId), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.Environment](res)
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
