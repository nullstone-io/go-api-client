package find

import (
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

func StackAndBlockByName(cfg api.Config, stackName, blockName string) (*types.Stack, *types.Block, error) {
	if stackName == "" {
		return blockByBlockNameNoStack(cfg, blockName)
	}
	return blockByStackAndBlockName(cfg, stackName, blockName)
}

func blockByStackAndBlockName(cfg api.Config, stackName, blockName string) (*types.Stack, *types.Block, error) {
	client := api.Client{Config: cfg}
	stack, err := client.StacksByName().Get(stackName)
	if err != nil {
		return nil, nil, err
	} else if stack == nil {
		return nil, nil, nil
	}
	blocks, err := client.Blocks().List(stack.Id)
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

func blockByBlockNameNoStack(cfg api.Config, blockName string) (*types.Stack, *types.Block, error) {
	client := api.Client{Config: cfg}
	stacks, err := client.Stacks().List()
	if err != nil {
		return nil, nil, fmt.Errorf("error retrieving stacks: %w", err)
	}

	stacksByName := map[string]*types.Stack{}
	foundBlocks := make([]types.Block, 0)
	foundStackNames := make([]string, 0)
	for _, stack := range stacks {
		stacksByName[stack.Name] = stack
		blocks, err := client.Blocks().List(stack.Id)
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
	return stacksByName[block.StackName], &block, nil
}
