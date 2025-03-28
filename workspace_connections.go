package api

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type WorkspaceConnections struct {
	Client *Client
}

func (wc WorkspaceConnections) basePath(stackId, blockId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/blocks/%d/envs/%d/connections", wc.Client.Config.OrgName, stackId, blockId, envId)
}

func (wc WorkspaceConnections) Update(ctx context.Context, stackId, blockId, envId int64, input []types.ConnectionInput) (*http.Response, error) {
	raw, _ := json.Marshal(input)
	res, err := wc.Client.Do(ctx, http.MethodPut, wc.basePath(stackId, blockId, envId), nil, nil, json.RawMessage(raw))
	if err != nil {
		return nil, err
	}

	return res, response.Verify(res)
}
