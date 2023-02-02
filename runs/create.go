package runs

import (
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

func Create(cfg api.Config, workspace types.Workspace, isApproved *bool, isDestroy bool, destroyDeps string) (*types.Run, error) {
	input := types.CreateRunInput{
		IsDestroy:           isDestroy,
		DestroyDependencies: destroyDeps,
		IsApproved:          isApproved,
	}

	client := api.Client{Config: cfg}
	newRun, err := client.Runs().Create(workspace.StackId, workspace.Uid, input)
	if err != nil {
		return nil, fmt.Errorf("error creating run: %w", err)
	}
	return newRun, nil
}
