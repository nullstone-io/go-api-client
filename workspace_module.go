package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type WorkspaceModule struct {
	Client *Client
}

func (wmv WorkspaceModule) basePath(stackId, blockId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/blocks/%d/envs/%d/module-version", wmv.Client.Config.OrgName, stackId, blockId, envId)
}

func (wmv WorkspaceModule) Update(stackId, blockId, envId int64, moduleInput types.WorkspaceModuleInput) (*types.WorkspaceChangeset, error) {
	raw, _ := json.Marshal(moduleInput)
	res, err := wmv.Client.Do(http.MethodPut, wmv.basePath(stackId, blockId, envId), nil, nil, json.RawMessage(raw))
	if err != nil {
		return nil, err
	}

	return response.ReadJsonPtr[types.WorkspaceChangeset](res)
}
