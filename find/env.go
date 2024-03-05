package find

import (
	"context"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

func Env(ctx context.Context, cfg api.Config, stackId int64, envName string) (*types.Environment, error) {
	client := api.Client{Config: cfg}
	envs, err := client.Environments().List(ctx, stackId)
	if err != nil {
		return nil, fmt.Errorf("error retrieving environments: %w", err)
	}
	for _, env := range envs {
		if env.Name == envName {
			return env, nil
		}
	}
	return nil, nil
}
