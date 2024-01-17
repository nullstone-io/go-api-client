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

func (p PipelineBlockSyncs) basePath(stackId, envId int64, repo string) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/envs/%d/repos/%s/block_syncs", p.Client.Config.OrgName, stackId, envId, repo)
}

// Create - POST /orgs/:orgName/stacks/:stack_id/envs/:env_id/repos/:repo_name/block_syncs
func (p PipelineBlockSyncs) Create(stackId, envId int64, repo string, payload BlockSyncPayload) ([]types.Block, error) {
	rawPayload, _ := json.Marshal(payload)
	res, err := p.Client.Do(http.MethodPost, p.basePath(stackId, envId, repo), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	return response.ReadJsonVal[[]types.Block](res)
}
