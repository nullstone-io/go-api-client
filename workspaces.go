package api

import (
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"path"
)

type Workspaces struct {
	Client *Client
}

func (w Workspaces) Get(stackName, blockName, envName string) (*types.Workspace, error) {
	res, err := w.Client.Do(http.MethodGet, path.Join("stacks", stackName, "blocks", blockName, "envs", envName), nil, nil)
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
