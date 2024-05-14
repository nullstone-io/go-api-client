package api

import (
	"context"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"gopkg.in/nullstone-io/go-api-client.v0/ws"
)

type BuildLogs struct {
	Client *Client
}

func (l BuildLogs) path(stackId int64, buildId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/builds/%d/logs", l.Client.Config.OrgName, stackId, buildId)
}

func (l BuildLogs) Watch(ctx context.Context, stackId int64, buildId int64, retryFn ws.StreamerRetryFunc) (<-chan types.Message, error) {
	endpoint, headers, err := l.Client.Config.ConstructWsEndpoint(ctx, l.path(stackId, buildId))
	if err != nil {
		return nil, err
	}
	return ws.StreamLogs(ctx, endpoint, headers, retryFn), nil
}
