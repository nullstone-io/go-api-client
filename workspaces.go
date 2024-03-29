package api

import (
	"context"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type Workspaces struct {
	Client *Client
}

func (w Workspaces) basePath(stackId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/workspaces", w.Client.Config.OrgName, stackId)
}

func (w Workspaces) workspacePath(stackId, blockId, envId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/blocks/%d/envs/%d", w.Client.Config.OrgName, stackId, blockId, envId)
}

// List - GET /orgs/:orgName/stacks/:stackId/workspaces
func (w Workspaces) List(ctx context.Context, stackId int64) ([]types.Workspace, error) {
	res, err := w.Client.Do(ctx, http.MethodGet, w.basePath(stackId), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonVal[[]types.Workspace](res)
}

// Get - GET /orgs/:orgName/stacks/:stackId/blocks/:blockId/envs/:envId
func (w Workspaces) Get(ctx context.Context, stackId, blockId, envId int64) (*types.Workspace, error) {
	res, err := w.Client.Do(ctx, http.MethodGet, w.workspacePath(stackId, blockId, envId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var workspace types.Workspace
	if err := response.ReadJson(res, &workspace); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &workspace, nil
}
