package api

import (
	"context"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"net/url"
	"strconv"
)

type WorkspaceWorkflows struct {
	Client *Client
}

func (ww WorkspaceWorkflows) basePath(stackId, blockId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/blocks/%d/envs/%d/workspace_workflows", ww.Client.Config.OrgName, stackId, blockId, envId)
}

func (ww WorkspaceWorkflows) List(ctx context.Context, stackId, blockId, envId int64, page, perPage int) ([]types.WorkspaceWorkflow, error) {
	q := url.Values{}
	if page > 0 {
		q.Set("page", strconv.Itoa(page))
	}
	if perPage > 0 {
		q.Set("perPage", strconv.Itoa(perPage))
	}
	res, err := ww.Client.Do(ctx, http.MethodGet, ww.basePath(stackId, blockId, envId), q, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonVal[[]types.WorkspaceWorkflow](res)
}
