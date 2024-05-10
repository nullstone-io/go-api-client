package api

import (
	"context"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type WorkspaceWorkflows struct {
	Client *Client
}

func (w WorkspaceWorkflows) basePath(orgName string, stackId, blockId, envId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/blocks/%d/envs/%d/workspace_workflows", orgName, stackId, blockId, envId)
}

func (w WorkspaceWorkflows) List(ctx context.Context, orgName string, stackId, blockId, envId int64) ([]types.WorkspaceWorkflowStatus, error) {
	res, err := w.Client.Do(ctx, http.MethodGet, w.basePath(orgName, stackId, blockId, envId), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonVal[[]types.WorkspaceWorkflowStatus](res)
}
