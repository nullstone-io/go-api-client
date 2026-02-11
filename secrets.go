package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

type Secrets struct {
	Client *Client
}

func (s Secrets) basePath(stackId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/envs/%d/secrets", s.Client.Config.OrgName, stackId, envId)
}

func (s Secrets) secretPath(stackId, envId, secretId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/envs/%d/secrets/%s", s.Client.Config.OrgName, stackId, envId, url.PathEscape(fmt.Sprintf("%d", secretId)))
}

type AddSecretInput struct {
	Id    string `json:"id"`
	Value string `json:"value"`
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
	Id    *string `json:"id,omitempty"`
	Value *string `json:"value,omitempty"`
}

func (s Secrets) Update(ctx context.Context, stackId, envId, secretId int64, input UpdateSecretInput) (*types.Secret, error) {
	raw, _ := json.Marshal(input)
	res, err := s.Client.Do(ctx, http.MethodPut, s.secretPath(stackId, envId, secretId), nil, nil, json.RawMessage(raw))
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.Secret](res)
}
