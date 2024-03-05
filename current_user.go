package api

import (
	"context"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
)

type CurrentUser struct {
	Client *Client
}

func (cu CurrentUser) Get(ctx context.Context) (*types.User, error) {
	res, err := cu.Client.Do(ctx, http.MethodGet, "/current_user", nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return response.ReadJsonPtr[types.User](res)
}
