package find

import (
	"context"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"sync"
)

type StackResolver struct {
	ApiClient           *api.Client
	EnvGetter           EnvGetter
	Stack               types.Stack
	PreviewsSharedEnvId int64
	EnvsById            map[int64]types.Environment
	EnvsByName          map[string]types.Environment
	BlocksById          map[int64]types.Block
	BlocksByName        map[string]types.Block

	envsOnce        sync.Once
	envsLoadError   error
	blocksOnce      sync.Once
	blocksLoadError error
}

func (r *StackResolver) Envs(ctx context.Context) (map[int64]types.Environment, error) {
	if err := r.ensureEnvs(ctx); err != nil {
		return nil, err
	}
	return r.EnvsById, nil
}

func (r *StackResolver) ResolveEnv(ctx context.Context, ct types.ConnectionTarget, curEnvId int64) (types.Environment, error) {
	// Prefer EnvId over EnvName -- it's possible for a user to rename the env
	if ct.EnvId != nil {
		return r.ResolveEnvById(ctx, *ct.EnvId)
	}
	if ct.EnvName != "" {
		return r.ResolveEnvByName(ctx, ct.EnvName)
	}
	return r.ResolveEnvById(ctx, curEnvId)
}

func (r *StackResolver) ResolveEnvByName(ctx context.Context, envName string) (types.Environment, error) {
	if env, ok := r.EnvsByName[envName]; ok {
		return env, nil
	}
	if err := r.ensureEnvs(ctx); err != nil {
		return types.Environment{}, err
	}
	if env, ok := r.EnvsByName[envName]; ok {
		return env, nil
	}
	return types.Environment{}, EnvDoesNotExistError{StackName: r.Stack.Name, EnvName: envName}
}

func (r *StackResolver) ResolveEnvById(ctx context.Context, envId int64) (types.Environment, error) {
	if env, ok := r.EnvsById[envId]; ok {
		return env, nil
	}
	if err := r.ensureEnvs(ctx); err != nil {
		return types.Environment{}, err
	}
	if env, ok := r.EnvsById[envId]; ok {
		return env, nil
	}
	return types.Environment{}, EnvIdDoesNotExistError{StackName: r.Stack.Name, EnvId: envId}
}

func (r *StackResolver) loadEnvs(ctx context.Context) error {
	envs, err := r.EnvGetter.ListEnvs(ctx, r.Stack.Id)
	if err != nil {
		return fmt.Errorf("unable to fetch environments (%s/%d): %w", r.Stack.OrgName, r.Stack.Id, err)
	}
	if r.EnvsById == nil {
		r.EnvsById = map[int64]types.Environment{}
	}
	if r.EnvsByName == nil {
		r.EnvsByName = map[string]types.Environment{}
	}
	for _, env := range envs {
		r.EnvsById[env.Id] = env
		r.EnvsByName[env.Name] = env
		if env.Type == types.EnvTypePreviewsShared {
			r.PreviewsSharedEnvId = env.Id
		}
	}
	return nil
}

func (r *StackResolver) ensureEnvs(ctx context.Context) error {
	r.envsOnce.Do(func() {
		r.envsLoadError = r.loadEnvs(ctx)
	})
	return r.envsLoadError

}

func (r *StackResolver) Blocks(ctx context.Context) (map[int64]types.Block, error) {
	if err := r.ensureBlocks(ctx); err != nil {
		return nil, err
	}
	return r.BlocksById, nil
}

func (r *StackResolver) ResolveBlock(ctx context.Context, ct types.ConnectionTarget) (types.Block, error) {
	// Prefer BlockId over BlockName -- it's possible for a user to rename the block
	if ct.BlockId != 0 {
		return r.ResolveBlockById(ctx, ct.BlockId)
	}
	if ct.BlockName != "" {
		return r.ResolveBlockByName(ctx, ct.BlockName)
	}
	return types.Block{}, BlockHasNoIdentifierError{StackName: r.Stack.Name}
}

func (r *StackResolver) ResolveBlockByName(ctx context.Context, blockName string) (types.Block, error) {
	if block, ok := r.BlocksByName[blockName]; ok {
		return block, nil
	}
	if err := r.ensureBlocks(ctx); err != nil {
		return types.Block{}, err
	}
	if block, ok := r.BlocksByName[blockName]; ok {
		return block, nil
	}
	return types.Block{}, BlockDoesNotExistError{StackName: r.Stack.Name, BlockName: blockName}
}

func (r *StackResolver) ResolveBlockById(ctx context.Context, blockId int64) (types.Block, error) {
	if block, ok := r.BlocksById[blockId]; ok {
		return block, nil
	}
	if err := r.ensureBlocks(ctx); err != nil {
		return types.Block{}, err
	}
	if block, ok := r.BlocksById[blockId]; ok {
		return block, nil
	}
	return types.Block{}, BlockIdDoesNotExistError{StackName: r.Stack.Name, BlockId: blockId}
}

func (r *StackResolver) ensureBlocks(ctx context.Context) error {
	r.blocksOnce.Do(func() {
		r.blocksLoadError = r.LoadBlocks(ctx)
	})
	return r.blocksLoadError
}

func (r *StackResolver) LoadBlocks(ctx context.Context) error {
	blocks, err := r.ApiClient.Blocks().List(ctx, r.Stack.Id, false)
	if err != nil {
		return fmt.Errorf("unable to fetch blocks (%s/%d): %w", r.Stack.OrgName, r.Stack.Id, err)
	}
	if r.BlocksById == nil {
		r.BlocksById = map[int64]types.Block{}
	}
	if r.BlocksByName == nil {
		r.BlocksByName = map[string]types.Block{}
	}
	for _, block := range blocks {
		r.BlocksById[block.Id] = block
		r.BlocksByName[block.Name] = block
	}
	return nil
}

// AddBlock adds a block to the stack resolver
// This only returns true if the block was added
// This returns false if there is an error or the block already exists
func (r *StackResolver) AddBlock(ctx context.Context, block types.Block) (bool, error) {
	if err := r.ensureBlocks(ctx); err != nil {
		return false, err
	}

	if block.Id != 0 {
		if _, ok := r.BlocksById[block.Id]; ok {
			return false, nil
		}
	}
	if block.Name != "" {
		if _, ok := r.BlocksByName[block.Name]; ok {
			return false, nil
		}
	}

	if block.Id != 0 {
		r.BlocksById[block.Id] = block
	}
	if block.Name != "" {
		r.BlocksByName[block.Name] = block
	}

	return true, nil
}

// AddEnv adds an env to the stack resolver
// This only returns true if the env was added
// This returns false if there is an error or the env already exists
func (r *StackResolver) AddEnv(ctx context.Context, env types.Environment) (bool, error) {
	if err := r.ensureEnvs(ctx); err != nil {
		return false, err
	}

	if env.Id != 0 {
		if _, ok := r.EnvsById[env.Id]; ok {
			return false, nil
		}
	}
	if env.Name != "" {
		if _, ok := r.EnvsByName[env.Name]; ok {
			return false, nil
		}
	}

	if env.Id != 0 {
		r.EnvsById[env.Id] = env
	}
	if env.Name != "" {
		r.EnvsByName[env.Name] = env
	}

	return true, nil
}
