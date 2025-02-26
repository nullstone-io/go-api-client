package api

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"net/url"
	"strconv"
)

type Blocks struct {
	Client *Client
}

func (s Blocks) basePath(stackId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/blocks", s.Client.Config.OrgName, stackId)
}

func (s Blocks) blockPath(stackId, blockId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/blocks/%d", s.Client.Config.OrgName, stackId, blockId)
}

// List - GET /orgs/:orgName/stacks/:stack_id/blocks
func (s Blocks) List(ctx context.Context, stackId int64, includeCapabilities bool) ([]types.Block, error) {
	var q url.Values
	if includeCapabilities {
		q = url.Values{"include_capabilities": []string{"true"}}
	}
	res, err := s.Client.Do(ctx, http.MethodGet, s.basePath(stackId), q, nil, nil)
	if err != nil {
		return nil, err
	}

	return response.ReadJsonVal[[]types.Block](res)
}

// Get - GET /orgs/:orgName/stacks/:stack_id/blocks/:id
func (s Blocks) Get(ctx context.Context, stackId, blockId int64, includeArchived bool) (*types.Block, error) {
	q := url.Values{
		"include_archived": []string{strconv.FormatBool(includeArchived)},
	}
	res, err := s.Client.Do(ctx, http.MethodGet, s.blockPath(stackId, blockId), q, nil, nil)
	if err != nil {
		return nil, err
	}

	var block types.Block
	if err := response.ReadJson(res, &block); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &block, nil
}

// Create - POST /orgs/:orgName/stacks/:stack_id/blocks
func (s Blocks) Create(ctx context.Context, stackId int64, block *types.Block) (*types.Block, error) {
	rawPayload, _ := json.Marshal(block)
	res, err := s.Client.Do(ctx, http.MethodPost, s.basePath(stackId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updatedBlock types.Block
	if err := response.ReadJson(res, &updatedBlock); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedBlock, nil
}

// Update - PUT/PATCH /orgs/:orgName/stacks/:stack_id/blocks/:id
func (s Blocks) Update(ctx context.Context, stackId, blockId int64, block *types.Block) (*types.Block, error) {
	rawPayload, _ := json.Marshal(block)
	res, err := s.Client.Do(ctx, http.MethodPut, s.blockPath(stackId, blockId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updatedBlock types.Block
	if err := response.ReadJson(res, &updatedBlock); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedBlock, nil
}

// Destroy - DELETE /orgs/:orgName/stacks/:stack_id/blocks/:id
func (s Blocks) Destroy(ctx context.Context, stackId, blockId int64) (bool, error) {
	res, err := s.Client.Do(ctx, http.MethodDelete, s.blockPath(stackId, blockId), nil, nil, nil)
	if err != nil {
		return false, err
	}
	if err := response.Verify(res); response.IsNotFoundError(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
