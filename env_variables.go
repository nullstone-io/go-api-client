package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type EnvVariables struct {
	Client *Client
}

func (ev EnvVariables) basePath(stackId, blockId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/blocks/%d/envs/%d/env-variables", ev.Client.Config.OrgName, stackId, blockId, envId)
}

func (ev EnvVariables) envVarPath(stackId, blockId, envId int64, key string) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/blocks/%d/envs/%d/env-variables/%s", ev.Client.Config.OrgName, stackId, blockId, envId, key)
}

func (ev EnvVariables) Create(stackId, blockId, envId int64, input []types.EnvVariableInput) (*types.WorkspaceChangeset, error) {
	raw, _ := json.Marshal(input)
	res, err := ev.Client.Do(http.MethodPost, ev.basePath(stackId, blockId, envId), nil, nil, json.RawMessage(raw))
	if err != nil {
		return nil, err
	}

	var changeset *types.WorkspaceChangeset
	if err := response.ReadJson(res, changeset); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return changeset, nil
}

func (ev EnvVariables) Destroy(stackId, blockId, envId int64, key string) (*types.WorkspaceChangeset, error) {
	res, err := ev.Client.Do(http.MethodDelete, ev.envVarPath(stackId, blockId, envId, key), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var changeset *types.WorkspaceChangeset
	if err := response.ReadJson(res, changeset); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return changeset, nil
}
