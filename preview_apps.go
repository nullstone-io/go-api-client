package api

import (
	"context"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type PreviewApps struct {
	Client *Client
}

func (pa PreviewApps) basePath(stackId, envId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/envs/%d/preview_apps", pa.Client.Config.OrgName, stackId, envId)
}

// List - GET /orgs/:orgName/stacks/:stack_id/envs/:env_id/preview_apps
func (pa PreviewApps) List(ctx context.Context, stackId, envId int64) ([]*types.PreviewApp, error) {
	res, err := pa.Client.Do(ctx, http.MethodGet, pa.basePath(stackId, envId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return response.ReadJsonVal[[]*types.PreviewApp](res)
}
