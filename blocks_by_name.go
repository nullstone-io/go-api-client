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

func (s BlocksByName) blockPath(stackName, blockName string) string {
	return path.Join("orgs", s.Client.Config.OrgName, "stacks_by_name", stackName, "blocks", blockName)
}

// Get - GET /orgs/:orgName/stacks_by_name/:stack_name/blocks/:name
func (s BlocksByName) Get(stackName, blockName string) (*types.Block, error) {
	res, err := s.Client.Do(http.MethodGet, s.blockPath(stackName, blockName), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	var env types.Block
	if err := response.ReadJson(res, &env); response.IsNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &env, nil
}
