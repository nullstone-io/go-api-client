package find

import (
	"context"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"strings"
)

// ConnectionTarget finds the target of a connection from a source stack
// A connection target can be specified in one of three ways:
//
//	{stack}.{env}.{block}
//	{stack}.{block}
//	{block}
func ConnectionTarget(ctx context.Context, cfg api.Config, sourceStackName, raw string) (*types.ConnectionTarget, error) {
	tokens := strings.SplitN(raw, ".", 3)
	switch len(tokens) {
	case 1: // {block}
		return connectionTargetByStackBlock(ctx, cfg, sourceStackName, tokens[0], raw)
	case 2: // {stack}.{block}
		return connectionTargetByStackBlock(ctx, cfg, tokens[0], tokens[1], raw)
	case 3: // {stack}.{env}.{block}
		return connectionTargetByStackEnvBlock(ctx, cfg, tokens[0], tokens[1], tokens[2], raw)
	default:
		return nil, fmt.Errorf("invalid connection target %q", raw)
	}
}

func connectionTargetByStackEnvBlock(ctx context.Context, cfg api.Config, stackName, envName, blockName, raw string) (*types.ConnectionTarget, error) {
	ct, err := connectionTargetByStackBlock(ctx, cfg, stackName, blockName, raw)
	if err != nil {
		return nil, err
	}

	client := api.Client{Config: cfg}
	env, err := client.EnvironmentsByName().Get(ctx, stackName, envName)
	if err != nil {
		return nil, fmt.Errorf("error searching for environment %q in stack %q: %w", envName, stackName, err)
	} else if env == nil {
		return nil, fmt.Errorf("environment %q does not exist in stack %q for mapping %q", envName, stackName, raw)
	}
	ct.EnvId = &env.Id

	return ct, nil
}

func connectionTargetByStackBlock(ctx context.Context, cfg api.Config, stackName, blockName, raw string) (*types.ConnectionTarget, error) {
	targetStack, targetBlock, err := blockByStackAndBlockName(ctx, cfg, stackName, blockName)
	if err != nil {
		return nil, err
	} else if targetStack == nil {
		return nil, fmt.Errorf("stack %q does not exist for mapping %q", stackName, raw)
	} else if targetBlock == nil {
		return nil, fmt.Errorf("block %q does not exist in stack %q for mapping %q", blockName, stackName, raw)
	}
	return &types.ConnectionTarget{
		StackId:   targetStack.Id,
		BlockId:   targetBlock.Id,
		BlockName: blockName,
		EnvId:     nil,
	}, nil
}
