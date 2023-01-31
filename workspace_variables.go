package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type WorkspaceVariables struct {
	Client *Client
}

func (wv WorkspaceVariables) basePath(stackId, blockId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/blocks/%d/envs/%d/variables", wv.Client.Config.OrgName, stackId, blockId, envId)
}

func (wv WorkspaceVariables) Update(stackId, blockId, envId int64, input []types.VariableInput) (*types.WorkspaceChangeset, error) {
	raw, _ := json.Marshal(input)
	res, err := wv.Client.Do(http.MethodPut, wv.basePath(stackId, blockId, envId), nil, nil, json.RawMessage(raw))
	if err != nil {
		return nil, err
	}

	return response.ReadJsonPtr[types.WorkspaceChangeset](res)
}
