package api

import (
	"context"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/live_logs"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"net/http"
	"net/url"
	"strings"
)

type DeployLiveLogs struct {
	Client *Client
}

func (l DeployLiveLogs) path(stackId int64, deployId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/deploys/%d/live_logs", l.Client.Config.OrgName, stackId, deployId)
}

func (l DeployLiveLogs) Watch(ctx context.Context, stackId int64, deployId int64) (<-chan types.LiveLogMessage, error) {
	endpoint, err := url.Parse(l.Client.Config.BaseAddress)
	if err != nil {
		return nil, fmt.Errorf("invalid url: %w", err)
	}
	endpoint.Path = l.path(stackId, deployId)
	endpoint.Scheme = strings.Replace(endpoint.Scheme, "http", "ws", 1)

	headers := http.Header{}
	headers.Set("Authorization", fmt.Sprintf("Bearer %s", l.Client.Config.ApiKey))

	streamer, err := live_logs.NewReconnectingStreamer(ctx, endpoint.String(), headers)
	if err != nil {
		return nil, err
	}
	return streamer.Stream(ctx), nil
}
