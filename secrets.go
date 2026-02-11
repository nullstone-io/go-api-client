package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

type Secrets struct {
	Client *Client
}

func (s Secrets) basePath(stackId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/envs/%d/secrets", s.Client.Config.OrgName, stackId, envId)
}

func (s Secrets) secretPath(stackId, envId int64, secretNameOrId string) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/envs/%d/secrets/%s", s.Client.Config.OrgName, stackId, envId, url.PathEscape(secretNameOrId))
}

func (s Secrets) List(ctx context.Context, stackId, envId int64, location types.SecretLocation) ([]types.Secret, error) {
	q, err := query.Values(location)
	if err != nil {
		return nil, fmt.Errorf("error encoding request query: %w", err)
	}
	res, err := s.Client.Do(ctx, http.MethodGet, s.basePath(stackId, envId), q, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonVal[[]types.Secret](res)
}

type AddSecretInput struct {
	Identity types.SecretIdentity `json:"identity"`
	Value    string               `json:"value"`
}

func (s Secrets) Add(ctx context.Context, stackId, envId int64, input AddSecretInput) (*types.Secret, error) {
	raw, _ := json.Marshal(input)
	res, err := s.Client.Do(ctx, http.MethodPost, s.basePath(stackId, envId), nil, nil, json.RawMessage(raw))
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.Secret](res)
}

type UpdateSecretInput struct {
	Value string `json:"value"`
}

// Update modifies the secret value in the platform's secret manager
// If secretNameOrId is Name, the id is inferred by using environment's default project/account/region info
// Specify a full Id (see types.SecretIdentity) to specify a different project/account/region
func (s Secrets) Update(ctx context.Context, stackId, envId int64, secretNameOrId string, input UpdateSecretInput) (*types.Secret, error) {
	raw, _ := json.Marshal(input)
	res, err := s.Client.Do(ctx, http.MethodPut, s.secretPath(stackId, envId, secretNameOrId), nil, nil, json.RawMessage(raw))
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.Secret](res)
}
