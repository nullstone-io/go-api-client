package api

import (
	"context"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type PipelineConnections struct {
	Client *Client
}

func (pc PipelineConnections) basePath(stackId, blockId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/blocks/%d/connections", pc.Client.Config.OrgName, stackId, blockId)
}

// List - GET /orgs/:orgName/stacks/:stackId/blocks/:block_id/connections
func (pc PipelineConnections) List(ctx context.Context, stackId, blockId int64) (map[string]types.ConnectionTarget, error) {
	res, err := pc.Client.Do(ctx, http.MethodGet, pc.basePath(stackId, blockId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return response.ReadJsonVal[map[string]types.ConnectionTarget](res)
}
