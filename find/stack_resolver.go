package find

import (
	"errors"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

var (
	ErrResolvingOtherEnvsNotSupported   = errors.New("Nullstone does not support resolving connections to environments from other stacks yet")
	ErrResolvingOtherBlocksNotSupported = errors.New("Nullstone does not support resolving connections to blocks from other stacks yet")
)

type StackResolver struct {
	Stack        types.Stack
	EnvsById     map[int64]types.Environment
	EnvsByName   map[string]types.Environment
	BlocksById   map[int64]types.Block
	BlocksByName map[string]types.Block
}

func (r *StackResolver) LoadEnvs(apiClient *api.Client, orgName string) error {
	envs, err := apiClient.Environments().List(r.Stack.Id)
	if err != nil {
		return fmt.Errorf("unable to fetch environments (%s/%d): %w", orgName, r.Stack.Id, err)
	}
	r.EnvsById = map[int64]types.Environment{}
	r.EnvsByName = map[string]types.Environment{}
	for _, env := range envs {
		r.EnvsById[env.Id] = *env
		r.EnvsByName[env.Name] = *env
	}
	return nil
}

func (r *StackResolver) LoadBlocks(apiClient *api.Client, orgName string) error {
	blocks, err := apiClient.Blocks().List(r.Stack.Id)
	if err != nil {
		return fmt.Errorf("unable to fetch blocks (%s/%d): %w", orgName, r.Stack.Id, err)
	}
	r.BlocksById = map[int64]types.Block{}
	r.BlocksByName = map[string]types.Block{}
	for _, block := range blocks {
		r.BlocksById[block.Id] = block
		r.BlocksByName[block.Name] = block
	}
	return nil
}

func (r *StackResolver) ResolveEnv(ct types.ConnectionTarget, curEnvId int64) (types.Environment, error) {
	if r.EnvsByName == nil {
		return types.Environment{}, ErrResolvingOtherEnvsNotSupported
	}

	if ct.EnvName != "" {
		env, ok := r.EnvsByName[ct.EnvName]
		if !ok {
			return types.Environment{}, EnvDoesNotExistError{StackName: r.Stack.Name, EnvName: ct.EnvName}
		}
		return env, nil
	}
	if ct.EnvId == nil {
		ct.EnvId = &curEnvId
	}
	env, ok := r.EnvsById[*ct.EnvId]
	if !ok {
		return types.Environment{}, EnvIdDoesNotExistError{StackName: r.Stack.Name, EnvId: *ct.EnvId}
	}
	return env, nil
}

func (r *StackResolver) ResolveBlock(ct types.ConnectionTarget) (types.Block, error) {
	if r.BlocksByName == nil {
		return types.Block{}, ErrResolvingOtherBlocksNotSupported
	}

	if ct.BlockName != "" {
		block, ok := r.BlocksByName[ct.BlockName]
		if !ok {
			return types.Block{}, BlockDoesNotExistError{StackName: r.Stack.Name, BlockName: ct.BlockName}
		}
		return block, nil
	}
	block, ok := r.BlocksById[ct.BlockId]
	if !ok {
		return types.Block{}, BlockIdDoesNotExistError{StackName: r.Stack.Name, BlockId: ct.BlockId}
	}
	return block, nil
}
