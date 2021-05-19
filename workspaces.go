package api

import (
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type Workspaces struct {
	Client *Client
}

func (w Workspaces) Get(stackId, blockId, envId int64) (*types.Workspace, error) {
	endpoint := fmt.Sprintf("stacks/%d/blocks/%d/envs/%d", stackId, blockId, envId)
	res, err := w.Client.Do(http.MethodGet, endpoint, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var workspace types.Workspace
	if err := w.Client.ReadJsonResponse(res, &workspace); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &workspace, nil
}
