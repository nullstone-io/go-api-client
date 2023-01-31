package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type WorkspaceModuleVersion struct {
	Client *Client
}

func (wmv WorkspaceModuleVersion) basePath(stackId, blockId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/blocks/%d/envs/%d/module-version", wmv.Client.Config.OrgName, stackId, blockId, envId)
}

func (wmv WorkspaceModuleVersion) Update(stackId, blockId, envId int64, input []types.VariableInput) (*types.WorkspaceChangeset, error) {
	raw, _ := json.Marshal(input)
	res, err := wmv.Client.Do(http.MethodPut, wmv.basePath(stackId, blockId, envId), nil, nil, json.RawMessage(raw))
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
