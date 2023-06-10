package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

// when creating a preview environment, we must have a name
type CreatePreviewEnvInput struct {
	Name string `json:"name"`
}

// when updating, we don't have to pass a name
// if you don't pass any fields, the api will simply make sure the preview environment is active
type UpdatePreviewEnvInput struct {
	Name *string `json:"name,omitempty"`
}

type PreviewEnvs struct {
	Client *Client
}

func (pe PreviewEnvs) basePath(stackId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/preview_envs", pe.Client.Config.OrgName, stackId)
}

func (pe PreviewEnvs) envPath(stackId, envId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/preview_envs/%d", pe.Client.Config.OrgName, stackId, envId)
}

// Create - POST /orgs/:orgName/stacks/:stack_id/preview_envs
func (pe PreviewEnvs) Create(stackId int64, env *CreatePreviewEnvInput) (*types.Environment, error) {
	rawPayload, _ := json.Marshal(env)
	res, err := pe.Client.Do(http.MethodPost, pe.basePath(stackId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	return response.ReadJsonPtr[types.Environment](res)
}

// Update - PUT/PATCH /orgs/:orgName/stacks/:stack_id/preview_envs/:id
func (pe PreviewEnvs) Update(stackId, envId int64, env *UpdatePreviewEnvInput) (*types.Environment, error) {
	rawPayload, _ := json.Marshal(env)
	res, err := pe.Client.Do(http.MethodPut, pe.envPath(stackId, envId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	return response.ReadJsonPtr[types.Environment](res)
}
