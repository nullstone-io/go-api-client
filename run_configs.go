package api

import (
	"fmt"
	"github.com/google/uuid"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type RunConfigs struct {
	Client *Client
}

func (c RunConfigs) runConfigPath(stackId int64, workspaceUid uuid.UUID) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/workspaces/%s/run-configs/latest", c.Client.Config.OrgName, stackId, workspaceUid)
}

func (c RunConfigs) GetLatest(stackId int64, workspaceUid uuid.UUID) (*types.RunConfig, error) {
	res, err := c.Client.Do(http.MethodGet, c.runConfigPath(stackId, workspaceUid), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var runConfig types.RunConfig
	if err := response.ReadJson(res, &runConfig); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &runConfig, nil
}
