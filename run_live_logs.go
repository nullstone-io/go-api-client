package api

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"gopkg.in/nullstone-io/go-api-client.v0/websocket"
	"net/http"
	"net/url"
	"strings"
)

type RunLiveLogs struct {
	Client *Client
}

func (l RunLiveLogs) path(stackId int64, runUid uuid.UUID) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/runs/%s/live_logs", l.Client.Config.OrgName, stackId, runUid)
}

func (l RunLiveLogs) Watch(ctx context.Context, stackId int64, runUid uuid.UUID) (<-chan types.LiveLogMessage, error) {
	endpoint, err := url.Parse(l.Client.Config.BaseAddress)
	if err != nil {
		return nil, fmt.Errorf("invalid url: %w", err)
	}
	endpoint.Path = l.path(stackId, runUid)
	endpoint.Scheme = strings.Replace(endpoint.Scheme, "http", "ws", 1)

	headers := http.Header{}
	headers.Set("Authorization", fmt.Sprintf("Bearer %s", l.Client.Config.ApiKey))

	streamer, err := websocket.NewReconnectingStreamer[types.LiveLogMessage](ctx, endpoint.String(), headers)
	if err != nil {
		return nil, err
	}
	return streamer.Stream(ctx), nil
}
