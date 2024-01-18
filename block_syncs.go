package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type BlockSyncs struct {
	Client *Client
}

type BlockSyncPayload struct {
	Blocks []types.Block `json:"blocks"`
}

func (s BlockSyncs) basePath(stackId, envId int64, repo string) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/envs/%d/repos/%s/block_syncs", s.Client.Config.OrgName, stackId, envId, repo)
}

// Create - POST /orgs/:orgName/stacks/:stack_id/repos/:repo_name/block_syncs
func (s BlockSyncs) Create(stackId, envId int64, repo string, payload BlockSyncPayload) ([]types.Block, error) {
	rawPayload, _ := json.Marshal(payload)
	res, err := s.Client.Do(http.MethodPost, s.basePath(stackId, envId, repo), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	return response.ReadJsonVal[[]types.Block](res)
}
