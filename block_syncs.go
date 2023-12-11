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
	owningRepo string
	blocks     []types.Block
}

func (s BlockSyncs) basePath(stackId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/block_syncs", s.Client.Config.OrgName, stackId)
}

// Create - POST /orgs/:orgName/stacks/:stack_id/block_syncs
func (s BlockSyncs) Create(stackId int64, payload BlockSyncPayload) ([]types.Block, error) {
	rawPayload, _ := json.Marshal(payload)
	res, err := s.Client.Do(http.MethodPost, s.basePath(stackId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	return response.ReadJsonVal[[]types.Block](res)
}
