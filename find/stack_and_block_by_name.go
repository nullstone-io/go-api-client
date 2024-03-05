package find

import (
	"context"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

func StackAndBlockByName(ctx context.Context, cfg api.Config, stackName, blockName string) (*types.Stack, *types.Block, error) {
	if stackName == "" {
		return blockByBlockNameNoStack(ctx, cfg, blockName)
	}
	return blockByStackAndBlockName(ctx, cfg, stackName, blockName)
}

func blockByStackAndBlockName(ctx context.Context, cfg api.Config, stackName, blockName string) (*types.Stack, *types.Block, error) {
	client := api.Client{Config: cfg}
	stack, err := client.StacksByName().Get(ctx, stackName)
	if err != nil {
		return nil, nil, err
	} else if stack == nil {
		return nil, nil, nil
	}
	blocks, err := client.Blocks().List(ctx, stack.Id)
	if err != nil {
		return nil, nil, err
	}
	for _, block := range blocks {
		if block.Name == blockName {
			return stack, &block, nil
		}
	}
	return stack, nil, err
}

func blockByBlockNameNoStack(ctx context.Context, cfg api.Config, blockName string) (*types.Stack, *types.Block, error) {
	client := api.Client{Config: cfg}
	stacks, err := client.Stacks().List(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("error retrieving stacks: %w", err)
	}

	stacksById := map[int64]*types.Stack{}
	foundBlocks := make([]types.Block, 0)
	foundStackNames := make([]string, 0)
	for _, stack := range stacks {
		stacksById[stack.Id] = stack
		blocks, err := client.Blocks().List(ctx, stack.Id)
		if err != nil {
			return nil, nil, fmt.Errorf("error retrieving blocks in stack (%s): %w", stack.Name, err)
		}
		for _, block := range blocks {
			if block.Name == blockName {
				foundBlocks = append(foundBlocks, block)
				foundStackNames = append(foundStackNames, stack.Name)
			}
		}
	}
	if len(foundBlocks) > 1 {
		return nil, nil, ErrMultipleBlocksFound{
			BlockName:  blockName,
			StackNames: foundStackNames,
		}
	} else if len(foundBlocks) < 1 {
		return nil, nil, nil
	}
	block := foundBlocks[0]
	return stacksById[block.StackId], &block, nil
}
