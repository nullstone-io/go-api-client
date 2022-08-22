package api

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"gopkg.in/nullstone-io/go-api-client.v0/ws"
)

type RunLiveLogs struct {
	Client *Client
}

func (l RunLiveLogs) path(stackId int64, runUid uuid.UUID) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/runs/%s/live_logs", l.Client.Config.OrgName, stackId, runUid)
}

func (l RunLiveLogs) Watch(ctx context.Context, stackId int64, runUid uuid.UUID, retryFn ws.StreamerRetryFunc) (<-chan types.Message, error) {
	endpoint, headers, err := l.Client.Config.ConstructWsEndpoint(l.path(stackId, runUid))
	if err != nil {
		return nil, err
	}
	return ws.StreamLogs(ctx, endpoint, headers, retryFn), nil
}
