package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type PipelineBlockSyncs struct {
	Client *Client
}

func (p PipelineBlockSyncs) basePath(stackId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/block_syncs", p.Client.Config.OrgName, stackId)
}

// Create - POST /orgs/:orgName/stacks/:stack_id/envs/:env_id/block_syncs
func (p PipelineBlockSyncs) Create(stackId int64, payload BlockSyncPayload) ([]types.Block, error) {
	rawPayload, _ := json.Marshal(payload)
	res, err := p.Client.Do(http.MethodPost, p.basePath(stackId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	return response.ReadJsonVal[[]types.Block](res)
}
