package api

import (
	"context"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type Builds struct {
	Client *Client
}

func (d Builds) basePath(stackId, appId, envId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/apps/%d/envs/%d/builds", d.Client.Config.OrgName, stackId, appId, envId)
}

func (d Builds) path(stackId, appId, envId, buildId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/apps/%d/envs/%d/builds/%d", d.Client.Config.OrgName, stackId, appId, envId, buildId)
}

func (d Builds) List(ctx context.Context, stackId, appId, envId int64) ([]*types.Build, error) {
	res, err := d.Client.Do(ctx, http.MethodGet, d.basePath(stackId, appId, envId), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonVal[[]*types.Build](res)
}

func (d Builds) Get(ctx context.Context, stackId, appId, envId, buildId int64) (*types.Build, error) {
	res, err := d.Client.Do(ctx, http.MethodGet, d.path(stackId, appId, envId, buildId), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.Build](res)
}
