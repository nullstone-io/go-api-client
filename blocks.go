package api

import (
	"encoding/json"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"path"
	"strconv"
)

type Blocks struct {
	Client *Client
}

// List - GET /orgs/:orgName/stacks/:stackName/blocks
func (s Blocks) List(stackName string) ([]types.Block, error) {
	res, err := s.Client.Do(http.MethodGet, path.Join("stacks", stackName, "blocks"), nil, nil, nil)
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

// Get - GET /orgs/:orgName/blocks/:id
func (s Blocks) Get(blockId int64) (*types.Block, error) {
	res, err := s.Client.Do(http.MethodGet, path.Join("blocks", strconv.FormatInt(blockId, 10)), nil, nil, nil)
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

// GetByName - GET /orgs/:orgName/stacks/:stackName/blocks/:name
func (s Blocks) GetByName(stackName, blockName string) (*types.Block, error) {
	res, err := s.Client.Do(http.MethodGet, path.Join("stacks", stackName, "blocks", blockName), nil, nil, nil)
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

// Create - POST /orgs/:orgName/stacks/:stackName/blocks
func (s Blocks) Create(stackName string, block *types.Block) (*types.Block, error) {
	rawPayload, _ := json.Marshal(block)
	res, err := s.Client.Do(http.MethodPost, path.Join("stacks", stackName, "blocks"), nil, nil, json.RawMessage(rawPayload))
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

// Update - PUT/PATCH /orgs/:orgName/blocks/:id
func (s Blocks) Update(blockId int, block *types.Block) (*types.Block, error) {
	rawPayload, _ := json.Marshal(block)
	endpoint := path.Join("blocks", strconv.Itoa(blockId))
	res, err := s.Client.Do(http.MethodPut, endpoint, nil, nil, json.RawMessage(rawPayload))
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

// Destroy - DELETE /orgs/:orgName/stacks/:stackName/blocks/:name
func (s Blocks) Destroy(stackName, blockName string) (bool, error) {
	res, err := s.Client.Do(http.MethodDelete, path.Join("stacks", stackName, "blocks", blockName), nil, nil, nil)
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
