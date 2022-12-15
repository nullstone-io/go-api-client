package runs

import (
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

// GetLatest retrieves the current run configuration on the workspace
func GetLatest(cfg api.Config, workspace types.Workspace) (*types.RunConfig, error) {
	client := api.Client{Config: cfg}
	newRunConfig, err := client.RunConfigs().GetLatest(workspace.StackId, workspace.Uid)
	if err != nil {
		return nil, err
	} else if newRunConfig == nil {
		return nil, fmt.Errorf("run config could not be found")
	}

	fillRunConfig(newRunConfig)
	return newRunConfig, nil
}
