package find

import (
	"context"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"strings"
)

func newErrMultipleAppsFound(apps []types.Application, stacks []*types.Stack) *ErrMultipleAppsFound {
	stackNames := make([]string, 0)
	for _, app := range apps {
		for _, stack := range stacks {
			if stack.Id == app.StackId {
				stackNames = append(stackNames, stack.Name)
			}
		}
	}
	return &ErrMultipleAppsFound{
		AppName:    apps[0].Name,
		StackNames: stackNames,
	}
}

type ErrMultipleAppsFound struct {
	AppName    string
	StackNames []string
}

func (e ErrMultipleAppsFound) Error() string {
	return fmt.Sprintf("found multiple applications named %q located in the following stacks: %s\n use the stack param to select a specific application", e.AppName, strings.Join(e.StackNames, ","))
}

// App searches for an app by app name
// If only 1 app is found, returns that app
// If many are found, will return an error with matched app stack names
func App(ctx context.Context, cfg api.Config, appName, stackName string) (*types.Application, error) {
	client := api.Client{Config: cfg}
	stackId := int64(0)
	if stackName != "" {
		stack, err := client.StacksByName().Get(ctx, stackName)
		if err != nil {
			return nil, fmt.Errorf("failed to find stack %q: %w", stackName, err)
		} else if stack == nil {
			return nil, StackDoesNotExistError{StackName: stackName}
		}
		stackId = stack.Id
	}
	allApps, err := client.Apps().List(ctx)
	if err != nil {
		return nil, fmt.Errorf("error listing applications: %w", err)
	}

	matched := make([]types.Application, 0)
	for _, app := range allApps {
		if app.Name == appName && (stackId == 0 || app.StackId == stackId) {
			matched = append(matched, app)
		}
	}

	if len(matched) == 0 {
		return nil, nil
	} else if len(matched) > 1 {
		stacks, err := client.Stacks().List(ctx)
		if err != nil {
			return nil, fmt.Errorf("error listing stacks: %w", err)
		}
		return nil, newErrMultipleAppsFound(matched, stacks)
	}

	return &matched[0], nil
}
