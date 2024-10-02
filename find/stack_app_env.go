package find

import (
	"context"
	"gopkg.in/nullstone-io/go-api-client.v0"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

func StackAppEnv(ctx context.Context, cfg api.Config, stackName, appName, envName string) (*types.Stack, *types.Application, *types.Environment, error) {
	app, err := App(ctx, cfg, appName, stackName)
	if err != nil {
		return nil, nil, nil, err
	} else if app == nil {
		return nil, nil, nil, AppDoesNotExistError{AppName: appName}
	}

	if stackName == "" {
		client := api.Client{Config: cfg}
		s, err := client.Stacks().Get(ctx, app.StackId, false)
		if err != nil {
			return nil, nil, nil, StackIdDoesNotExistError{StackId: app.StackId}
		}
		stackName = s.Name
	}
	stack, err := Stack(ctx, cfg, stackName)
	if err != nil {
		return nil, nil, nil, err
	} else if stack == nil {
		return nil, nil, nil, StackDoesNotExistError{StackName: stackName}
	}

	env, err := Env(ctx, cfg, stack.Id, envName)
	if err != nil {
		return nil, nil, nil, err
	} else if env == nil {
		return nil, nil, nil, EnvDoesNotExistError{StackName: stack.Name, EnvName: envName}
	}

	return stack, app, env, nil
}
