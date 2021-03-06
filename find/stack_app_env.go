package find

import (
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

func StackAppEnv(cfg api.Config, stackName, appName, envName string) (*types.Stack, *types.Application, *types.Environment, error) {
	app, err := App(cfg, appName, stackName)
	if err != nil {
		return nil, nil, nil, err
	} else if app == nil {
		return nil, nil, nil, fmt.Errorf("application %q does not exist", appName)
	}

	stack, err := Stack(cfg, app.StackName)
	if err != nil {
		return nil, nil, nil, err
	} else if stack == nil {
		return nil, nil, nil, fmt.Errorf("stack %q does not exist", app.StackName)
	}

	env, err := Env(cfg, stack.Id, envName)
	if err != nil {
		return nil, nil, nil, err
	} else if env == nil {
		return nil, nil, nil, fmt.Errorf("environment %s/%s does not exist", stack.Name, envName)
	}

	return stack, app, env, nil
}
