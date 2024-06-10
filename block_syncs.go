package api

import (
	"context"
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
	Blocks   []types.BlockSync `json:"blocks"`
	RepoName string            `json:"repoName"`
}

func (s BlockSyncs) basePath(stackId, envId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/envs/%d/block_syncs", s.Client.Config.OrgName, stackId, envId)
}

// Create - POST /orgs/:orgName/stacks/:stack_id/block_syncs
func (s BlockSyncs) Create(ctx context.Context, stackId, envId int64, payload BlockSyncPayload) ([]types.Block, error) {
	rawPayload, _ := json.Marshal(payload)
	res, err := s.Client.Do(ctx, http.MethodPost, s.basePath(stackId, envId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	return response.ReadJsonVal[[]types.Block](res)
}
