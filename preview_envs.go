package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type PreviewEnvironments struct {
	Client *Client
}

func (pe PreviewEnvironments) basePath(stackId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/preview_envs", pe.Client.Config.OrgName, stackId)
}

// Create - POST /orgs/:orgName/stacks/:stack_id/envs
func (pe PreviewEnvironments) Create(stackId int64, env *types.Environment) (*types.Environment, error) {
	rawPayload, _ := json.Marshal(env)
	res, err := pe.Client.Do(http.MethodPost, pe.basePath(stackId), nil, nil, json.RawMessage(rawPayload))
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
