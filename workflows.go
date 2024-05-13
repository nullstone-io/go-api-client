package api

import (
	"context"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"net/url"
)

type WorkspaceWorkflows struct {
	Client *Client
}

func (w WorkspaceWorkflows) basePath(stackId, blockId, envId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/blocks/%d/envs/%d/workspace_workflows", w.Client.Config.OrgName, stackId, blockId, envId)
}

func (w WorkspaceWorkflows) List(ctx context.Context, stackId, blockId, envId int64, page, pageSize *int64) ([]types.WorkspaceWorkflowStatus, error) {
	query := url.Values{}
	if page != nil {
		query.Set("page", fmt.Sprint(*page))
	}
	if pageSize != nil {
		query.Set("pageSize", fmt.Sprint(*pageSize))
	}
	res, err := w.Client.Do(ctx, http.MethodGet, w.basePath(stackId, blockId, envId), query, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonVal[[]types.WorkspaceWorkflowStatus](res)
}
