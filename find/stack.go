package find

import (
	"context"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

func Stack(ctx context.Context, cfg api.Config, stackName string) (*types.Stack, error) {
	client := api.Client{Config: cfg}
	stacks, err := client.Stacks().List(ctx)
	if err != nil {
		return nil, fmt.Errorf("error retrieving stacks: %w", err)
	}
	for _, stack := range stacks {
		if stack.Name == stackName {
			return stack, nil
		}
	}
	return nil, nil
}
