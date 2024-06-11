package api

import (
	"context"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type Connections struct {
	Client *Client
}

func (c Connections) basePath(stackId, blockId, envId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/blocks/%d/envs/%d/connections", c.Client.Config.OrgName, stackId, blockId, envId)
}

// List - GET /orgs/:orgName/stacks/:stackId/blocks/:block_id/envs/:env_id/connections
func (c Connections) List(ctx context.Context, stackId, blockId, envId int64) (map[string]types.ConnectionTarget, error) {
	res, err := c.Client.Do(ctx, http.MethodGet, c.basePath(stackId, blockId, envId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return response.ReadJsonVal[map[string]types.ConnectionTarget](res)
}
