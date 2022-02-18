package api

import (
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type WorkspaceOutputs struct {
	Client *Client
}

func (w WorkspaceOutputs) workspaceOutputsPath(stackId, blockId, envId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/blocks/%d/envs/%d/outputs/latest", w.Client.Config.OrgName, stackId, blockId, envId)
}

func (w WorkspaceOutputs) GetLatest(stackId, blockId, envId int64) (*types.Outputs, error) {
	res, err := w.Client.Do(http.MethodGet, w.workspaceOutputsPath(stackId, blockId, envId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var outputs types.Outputs
	if err := response.ReadJson(res, &outputs); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &outputs, nil
}
