package api

import (
	"context"
	"fmt"
	"net/http"

	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

type PreviewApps struct {
	Client *Client
}

func (p PreviewApps) basePath(stackId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/envs/%d/preview_apps", p.Client.Config.OrgName, stackId, envId)
}

// List - GET /orgs/{orgName}/stacks/{stackId}/envs/{envId}/preview_apps
func (p PreviewApps) List(ctx context.Context, stackId, envId int64) ([]types.PreviewApp, error) {
	res, err := p.Client.Do(ctx, http.MethodGet, p.basePath(stackId, envId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return response.ReadJsonVal[[]types.PreviewApp](res)
}
