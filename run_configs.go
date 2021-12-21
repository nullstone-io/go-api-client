package api

import (
	"fmt"
	"github.com/google/uuid"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type RunConfigs struct {
	Client *Client
}

func (c RunConfigs) GetLatest(stackId int64, workspaceUid uuid.UUID) (*types.RunConfig, error) {
	endpoint := fmt.Sprintf("stacks/%d/workspaces/%s/run-configs/latest", stackId, workspaceUid)
	res, err := c.Client.Do(http.MethodGet, endpoint, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.Json[types.RunConfig](res)
}
