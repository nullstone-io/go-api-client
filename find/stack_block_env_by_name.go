package find

import (
	"context"
	"errors"
	"gopkg.in/nullstone-io/go-api-client.v0"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

var (
	ErrStackDoesNotExist = errors.New("stack does not exist")
	ErrBlockDoesNotExist = errors.New("block does not exist")
	ErrEnvDoesNotExist   = errors.New("env does not exist")
)

type StackBlockEnv struct {
	Block types.Block
	Stack types.Stack
	Env   types.Environment
}

// StackBlockEnvByName looks for a workspace by stackName, blockName, and envName
// If stackName is "", will search for a block across all stacks
// stackName is required if there are multiple blocks with the same name in different stacks
func StackBlockEnvByName(ctx context.Context, cfg api.Config, stackName, blockName, envName string) (*StackBlockEnv, error) {
	stack, block, err := StackAndBlockByName(ctx, cfg, stackName, blockName)
	if err != nil {
		return nil, err
	} else if block == nil {
		return nil, ErrBlockDoesNotExist
	} else if stack == nil {
		return nil, ErrStackDoesNotExist
	}

	client := api.Client{Config: cfg}
	env, err := client.EnvironmentsByName().Get(ctx, stack.Name, envName)
	if err != nil {
		return nil, err
	} else if env == nil {
		return nil, ErrEnvDoesNotExist
	}

	return &StackBlockEnv{
		Block: *block,
		Stack: *stack,
		Env:   *env,
	}, nil
}
