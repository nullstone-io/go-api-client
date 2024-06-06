package api

import (
	"context"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type WorkspaceDetails struct {
	Client *Client
}

func (d WorkspaceDetails) basePath(stackId, blockId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/blocks/%d/envs/%d/workspace_details", d.Client.Config.OrgName, stackId, blockId, envId)
}

func (d WorkspaceDetails) Get(ctx context.Context, stackId, blockId, envId int64) (*types.WorkspaceDetails, error) {
	res, err := d.Client.Do(ctx, http.MethodGet, d.basePath(stackId, blockId, envId), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.WorkspaceDetails](res)
}
