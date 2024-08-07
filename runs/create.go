package runs

import (
	"context"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

func Create(ctx context.Context, cfg api.Config, workspace types.Workspace, commitSha string, isApproved *bool, isDestroy bool, destroyDeps string) (*api.RunCreateResult, error) {
	input := types.CreateRunInput{
		CommitSha:           commitSha,
		IsDestroy:           isDestroy,
		DestroyDependencies: destroyDeps,
		IsApproved:          isApproved,
	}

	client := api.Client{Config: cfg}
	result, err := client.Runs().Create(ctx, workspace.StackId, workspace.Uid, input)
	if err != nil {
		return nil, fmt.Errorf("error creating run: %w", err)
	}
	return result, nil
}
