package find

import (
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"strings"
)

func newErrMultipleAppsFound(apps []types.Application) *ErrMultipleAppsFound {
	stackNames := make([]string, 0)
	for _, app := range apps {
		stackNames = append(stackNames, app.StackName)
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
	return fmt.Sprintf("found multiple applications named %q located in the following stacks: %s", e.AppName, strings.Join(e.StackNames, ","))
}


// App searches for an app by app name
// If only 1 app is found, returns that app
// If many are found, will return an error with matched app stack names
func App(cfg api.Config, appName, stackName string) (*types.Application, error) {
	client := api.Client{Config: cfg}
	allApps, err := client.Apps().List()
	if err != nil {
		return nil, fmt.Errorf("error listing applications: %w", err)
	}

	matched := make([]types.Application, 0)
	for _, app := range allApps {
		if app.Name == appName && (stackName == "" || app.StackName == stackName) {
			matched = append(matched, app)
		}
	}

	if len(matched) == 0 {
		return nil, nil
	} else if len(matched) > 1 {
		return nil, newErrMultipleAppsFound(matched)
	}

	return &matched[0], nil
}

