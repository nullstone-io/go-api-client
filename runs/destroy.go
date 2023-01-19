package runs

import (
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

func Destroy(cfg api.Config, stackId, appId, envId int64, deps string, approve bool) (*types.Run, error) {
	client := api.Client{Config: cfg}
	workspace, err := client.Workspaces().Get(stackId, appId, envId)
	if err != nil {
		return nil, fmt.Errorf("error looking for workspace: %w", err)
	} else if workspace == nil {
		return nil, nil
	}

	runConfig, err := client.RunConfigs().GetLatest(workspace.StackId, workspace.Uid)
	if err != nil {
		return nil, fmt.Errorf("error retrieving existing config: %w", err)
	}
	fillRunConfig(runConfig)
	var isApproved *bool
	if approve {
		isApproved = &approve
	}
	return Create(cfg, *workspace, runConfig, isApproved, true, deps)
}
