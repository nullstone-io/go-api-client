package api

import (
	"context"
	"net/http"

	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

type NullstoneAgent struct {
	Client *Client
}

func (a NullstoneAgent) Get(ctx context.Context) (*types.NullstoneAgentInfo, error) {
	res, err := a.Client.Do(ctx, http.MethodGet, "nullstone_agent", nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.NullstoneAgentInfo](res)
}
