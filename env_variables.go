package api

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type EnvVariables struct {
	Client *Client
}

func (ev EnvVariables) basePath(stackId int64, workspaceUid uuid.UUID) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/workspaces/%s/env-variables", ev.Client.Config.OrgName, stackId, workspaceUid)
}

func (ev EnvVariables) envVarPath(stackId int64, workspaceUid uuid.UUID, key string) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/workspaces/%s/env-variables/%s", ev.Client.Config.OrgName, stackId, workspaceUid, key)
}

func (ev EnvVariables) Create(stackId int64, workspaceUid uuid.UUID, input []types.EnvVariableInput) ([]types.WorkspaceChange, error) {
	raw, _ := json.Marshal(input)
	res, err := ev.Client.Do(http.MethodPost, ev.basePath(stackId, workspaceUid), nil, nil, json.RawMessage(raw))
	if err != nil {
		return nil, err
	}

	var changes []types.WorkspaceChange
	if err := response.ReadJson(res, &changes); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return changes, nil
}

func (ev EnvVariables) Destroy(stackId int64, workspaceUid uuid.UUID, key string) ([]types.WorkspaceChange, error) {
	res, err := ev.Client.Do(http.MethodDelete, ev.envVarPath(stackId, workspaceUid, key), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var changes []types.WorkspaceChange
	if err := response.ReadJson(res, &changes); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return changes, nil
}
