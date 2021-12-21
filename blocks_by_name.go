package api

import (
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"path"
)

type BlocksByName struct {
	Client *Client
}

// Get - GET /orgs/:orgName/stacks_by_name/:stack_name/blocks/:name
func (s BlocksByName) Get(stackName, blockName string) (*types.Block, error) {
	res, err := s.Client.Do(http.MethodGet, path.Join("stacks_by_name", stackName, "blocks", blockName), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.Json[types.Block](res)
}
