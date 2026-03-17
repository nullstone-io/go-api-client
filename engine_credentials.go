package api

import (
	"context"
	"fmt"
	"net/http"

	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

type EngineCredentials struct {
	Client *Client
}

func (s EngineCredentials) path(stackId, envId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/envs/%d/engine_credentials", s.Client.Config.OrgName, stackId, envId)
}

// Acquire - POST /orgs/:orgName/stacks/:stackId/envs/:envId/engine_credentials
func (s EngineCredentials) Acquire(ctx context.Context, stackId, envId int64) (*types.EnvEngineCredentials, error) {
	res, err := s.Client.Do(ctx, http.MethodPost, s.path(stackId, envId), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.EnvEngineCredentials](res)
}
