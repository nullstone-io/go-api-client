package api

import (
	"context"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type WorkspaceConfigs struct {
	Client *Client
}

func (w WorkspaceConfigs) currentPath(stackId, blockId, envId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/blocks/%d/envs/%d/configs/current", w.Client.Config.OrgName, stackId, blockId, envId)
}

func (w WorkspaceConfigs) effectivePath(stackId, blockId, envId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/blocks/%d/envs/%d/configs/effective", w.Client.Config.OrgName, stackId, blockId, envId)
}

// GetCurrent - GET /orgs/:orgName/stacks/:stackId/blocks/:blockId/envs/:envId/configs/current
func (w WorkspaceConfigs) GetCurrent(ctx context.Context, stackId, blockId, envId int64) (*types.WorkspaceConfig, error) {
	res, err := w.Client.Do(ctx, http.MethodGet, w.currentPath(stackId, blockId, envId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return response.ReadJsonPtr[types.WorkspaceConfig](res)
}

// GetEffective - GET /orgs/:orgName/stacks/:stackId/blocks/:blockId/envs/:envId/configs/effective
func (w WorkspaceConfigs) GetEffective(ctx context.Context, stackId, blockId, envId int64) (*types.WorkspaceConfig, error) {
	res, err := w.Client.Do(ctx, http.MethodGet, w.currentPath(stackId, blockId, envId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return response.ReadJsonPtr[types.WorkspaceConfig](res)
}
