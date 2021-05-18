package api

import (
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"path"
)

type BlocksByName struct {
	Client *Client
}

// Get - GET /orgs/:orgName/stacks/:stackName/blocks/:name
func (s BlocksByName) Get(stackName, blockName string) (*types.Block, error) {
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
