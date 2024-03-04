package api

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

// UpdatePreviewEnvInput
// when updating, we don't have to pass a name
// if you don't pass any fields, the api will simply make sure the preview environment is active
type UpdatePreviewEnvInput struct {
	Name *string `json:"name,omitempty"`
}

type PreviewEnvs struct {
	Client *Client
}

func (pe PreviewEnvs) envPath(stackId, envId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/preview_envs/%d", pe.Client.Config.OrgName, stackId, envId)
}

// Get - GET /orgs/:orgName/stacks/:stack_id/preview_envs/:id
func (pe PreviewEnvs) Get(ctx context.Context, stackId, envId int64) (*types.Environment, error) {
	res, err := pe.Client.Do(ctx, http.MethodGet, pe.envPath(stackId, envId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return response.ReadJsonPtr[types.Environment](res)
}

// Update - PUT/PATCH /orgs/:orgName/stacks/:stack_id/preview_envs/:id
func (pe PreviewEnvs) Update(ctx context.Context, stackId, envId int64, env *UpdatePreviewEnvInput) (*types.Environment, error) {
	rawPayload, _ := json.Marshal(env)
	res, err := pe.Client.Do(ctx, http.MethodPut, pe.envPath(stackId, envId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	return response.ReadJsonPtr[types.Environment](res)
}
