package find

import (
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

type StackResolver struct {
	ApiClient           *api.Client
	Stack               types.Stack
	PreviewsSharedEnvId int64
	EnvsById            map[int64]types.Environment
	EnvsByName          map[string]types.Environment
	BlocksById          map[int64]types.Block
	BlocksByName        map[string]types.Block
}

func (r *StackResolver) ResolveEnv(ct types.ConnectionTarget, curEnvId int64) (types.Environment, error) {
	if ct.EnvName != "" {
		return r.resolveEnvByName(ct.EnvName)
	}
	if ct.EnvId == nil {
		ct.EnvId = &curEnvId
	}
	return r.resolveEnvById(*ct.EnvId)
}

func (r *StackResolver) resolveEnvByName(envName string) (types.Environment, error) {
	if env, ok := r.EnvsByName[envName]; ok {
		return env, nil
	}
	if err := r.loadEnvs(); err != nil {
		return types.Environment{}, err
	}
	if env, ok := r.EnvsByName[envName]; ok {
		return env, nil
	}
	return types.Environment{}, EnvDoesNotExistError{StackName: r.Stack.Name, EnvName: envName}
}

func (r *StackResolver) resolveEnvById(envId int64) (types.Environment, error) {
	if env, ok := r.EnvsById[envId]; ok {
		return env, nil
	}
	if err := r.loadEnvs(); err != nil {
		return types.Environment{}, err
	}
	if env, ok := r.EnvsById[envId]; ok {
		return env, nil
	}
	return types.Environment{}, EnvIdDoesNotExistError{StackName: r.Stack.Name, EnvId: envId}
}

func (r *StackResolver) loadEnvs() error {
	envs, err := r.ApiClient.Environments().List(r.Stack.Id)
	if err != nil {
		return fmt.Errorf("unable to fetch environments (%s/%d): %w", r.Stack.OrgName, r.Stack.Id, err)
	}
	r.EnvsById = map[int64]types.Environment{}
	r.EnvsByName = map[string]types.Environment{}
	for _, env := range envs {
		r.EnvsById[env.Id] = *env
		r.EnvsByName[env.Name] = *env
		if env.Type == types.EnvTypePreviewsShared {
			r.PreviewsSharedEnvId = env.Id
		}
	}
	return nil
}

func (r *StackResolver) Blocks() ([]types.Block, error) {
	if len(r.BlocksById) == 0 {
		if err := r.loadBlocks(); err != nil {
			return nil, err
		}
	}
	result := make([]types.Block, 0, len(r.BlocksById))
	for _, block := range r.BlocksById {
		result = append(result, block)
	}
	return result, nil
}

func (r *StackResolver) ResolveBlock(ct types.ConnectionTarget) (types.Block, error) {
	if ct.BlockName != "" {
		return r.resolveBlockByName(ct.BlockName)
	}
	return r.resolveBlockById(ct.BlockId)
}

func (r *StackResolver) resolveBlockByName(blockName string) (types.Block, error) {
	if block, ok := r.BlocksByName[blockName]; ok {
		return block, nil
	}
	if err := r.loadBlocks(); err != nil {
		return types.Block{}, err
	}
	if block, ok := r.BlocksByName[blockName]; ok {
		return block, nil
	}
	return types.Block{}, BlockDoesNotExistError{StackName: r.Stack.Name, BlockName: blockName}
}

func (r *StackResolver) resolveBlockById(blockId int64) (types.Block, error) {
	if block, ok := r.BlocksById[blockId]; ok {
		return block, nil
	}
	if err := r.loadBlocks(); err != nil {
		return types.Block{}, err
	}
	if block, ok := r.BlocksById[blockId]; ok {
		return block, nil
	}
	return types.Block{}, BlockIdDoesNotExistError{StackName: r.Stack.Name, BlockId: blockId}
}

func (r *StackResolver) loadBlocks() error {
	blocks, err := r.ApiClient.Blocks().List(r.Stack.Id)
	if err != nil {
		return fmt.Errorf("unable to fetch blocks (%s/%d): %w", r.Stack.OrgName, r.Stack.Id, err)
	}
	r.BlocksById = map[int64]types.Block{}
	r.BlocksByName = map[string]types.Block{}
	for _, block := range blocks {
		r.BlocksById[block.Id] = block
		r.BlocksByName[block.Name] = block
	}
	return nil
}
