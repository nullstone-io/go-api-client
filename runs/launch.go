package runs

import (
	"context"
	"fmt"
	"time"

	"gopkg.in/nullstone-io/go-api-client.v0"
)

func Launch(ctx context.Context, cfg api.Config, stackId, appId, envId int64, commitSha string, latestUpdateAt time.Time, approve bool) (*api.RunCreateResult, error) {
	client := api.Client{Config: cfg}
	workspace, err := client.Workspaces().Get(ctx, stackId, appId, envId)
	if err != nil {
		return nil, fmt.Errorf("error looking for workspace: %w", err)
	}

	var isApproved *bool
	if approve {
		isApproved = &approve
	}
	return Create(ctx, cfg, *workspace, commitSha, isApproved, latestUpdateAt, false, "")
}
