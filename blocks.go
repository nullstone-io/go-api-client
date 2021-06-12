package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type Blocks struct {
	Client *Client
}

func (s Blocks) basePath(stackId int64) string {
	return fmt.Sprintf("stacks/%d/blocks", stackId)
}

func (s Blocks) blockPath(stackId, blockId int64) string {
	return fmt.Sprintf("stacks/%d/blocks/%d", stackId, blockId)
}

// List - GET /orgs/:orgName/stacks/:stack_id/blocks
func (s Blocks) List(stackId int64) ([]types.Block, error) {
	res, err := s.Client.Do(http.MethodGet, s.basePath(stackId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var blocks []types.Block
	if err := s.Client.ReadJsonResponse(res, &blocks); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return blocks, nil
}

// Get - GET /orgs/:orgName/stacks/:stack_id/blocks/:id
func (s Blocks) Get(stackId, blockId int64) (*types.Block, error) {
	res, err := s.Client.Do(http.MethodGet, s.blockPath(stackId, blockId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var block types.Block
	if err := s.Client.ReadJsonResponse(res, &block); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &block, nil
}

// Create - POST /orgs/:orgName/stacks/:stack_id/blocks
func (s Blocks) Create(stackId int64, block *types.Block) (*types.Block, error) {
	rawPayload, _ := json.Marshal(block)
	res, err := s.Client.Do(http.MethodPost, s.basePath(stackId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updatedBlock types.Block
	if err := s.Client.ReadJsonResponse(res, &updatedBlock); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedBlock, nil
}

// Update - PUT/PATCH /orgs/:orgName/stacks/:stack_id/blocks/:id
func (s Blocks) Update(stackId, blockId int64, block *types.Block) (*types.Block, error) {
	rawPayload, _ := json.Marshal(block)
	res, err := s.Client.Do(http.MethodPut, s.blockPath(stackId, blockId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}

	var updatedBlock types.Block
	if err := s.Client.ReadJsonResponse(res, &updatedBlock); IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &updatedBlock, nil
}

// Destroy - DELETE /orgs/:orgName/stacks/:stack_id/blocks/:id
func (s Blocks) Destroy(stackId, blockId int64) (bool, error) {
	res, err := s.Client.Do(http.MethodDelete, s.blockPath(stackId, blockId), nil, nil, nil)
	if err != nil {
		return false, err
	}
	if err := s.Client.VerifyResponse(res); IsNotFoundError(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
