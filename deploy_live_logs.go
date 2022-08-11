package api

import (
	"context"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"gopkg.in/nullstone-io/go-api-client.v0/ws"
)

type DeployLiveLogs struct {
	Client *Client
}

func (l DeployLiveLogs) path(stackId int64, deployId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/deploys/%d/live_logs", l.Client.Config.OrgName, stackId, deployId)
}

func (l DeployLiveLogs) Watch(ctx context.Context, stackId int64, deployId int64, retryFn ws.StreamerRetryFunc) (<-chan types.LiveLogMessage, error) {
	endpoint, headers, err := l.Client.Config.ConstructWsEndpoint(l.path(stackId, deployId))
	if err != nil {
		return nil, err
	}
	return ws.StreamLogs(ctx, endpoint, headers, retryFn), nil
}
