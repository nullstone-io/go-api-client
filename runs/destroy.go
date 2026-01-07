package runs

import (
	"context"
	"fmt"
	"time"

	"gopkg.in/nullstone-io/go-api-client.v0"
)

func Destroy(ctx context.Context, cfg api.Config, stackId, appId, envId int64, commitSha string, latestUpdateAt time.Time, deps string, approve bool) (*api.RunCreateResult, error) {
	client := api.Client{Config: cfg}
	workspace, err := client.Workspaces().Get(ctx, stackId, appId, envId)
	if err != nil {
		return nil, fmt.Errorf("error looking for workspace: %w", err)
	} else if workspace == nil {
		return nil, nil
	}

	var isApproved *bool
	if approve {
		isApproved = &approve
	}
	return Create(ctx, cfg, *workspace, commitSha, isApproved, latestUpdateAt, true, deps)
}
