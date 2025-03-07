package api

import (
	"context"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type WorkspaceChanges struct {
	Client *Client
}

func (wc WorkspaceChanges) basePath(stackId, blockId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/blocks/%d/envs/%d/changes", wc.Client.Config.OrgName, stackId, blockId, envId)
}

func (wc WorkspaceChanges) changePath(stackId, blockId, envId, changeId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/blocks/%d/envs/%d/changes/%d", wc.Client.Config.OrgName, stackId, blockId, envId, changeId)
}

func (wc WorkspaceChanges) List(ctx context.Context, stackId, blockId, envId int64) (*types.WorkspaceChangeset, error) {
	res, err := wc.Client.Do(ctx, http.MethodGet, wc.basePath(stackId, blockId, envId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return response.ReadJsonPtr[types.WorkspaceChangeset](res)
}

func (wc WorkspaceChanges) Destroy(ctx context.Context, stackId, blockId, envId, changeId int64) error {
	res, err := wc.Client.Do(ctx, http.MethodDelete, wc.changePath(stackId, blockId, envId, changeId), nil, nil, nil)
	if err != nil {
		return err
	}

	return response.Verify(res)
}
