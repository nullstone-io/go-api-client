package api

import (
	"github.com/google/uuid"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"path"
)

type RunConfigs struct {
	Client *Client
}

func (c RunConfigs) GetLatest(stackName string, workspaceUid uuid.UUID) (*types.RunConfig, error) {
	res, err := c.Client.Do(http.MethodGet, path.Join("stacks", stackName, "workspaces", workspaceUid.String(), "run-configs", "latest"), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var runConfig types.RunConfig
	if err := c.Client.ReadJsonResponse(res, &runConfig); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &runConfig, nil
}
