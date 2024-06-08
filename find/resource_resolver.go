package find

import (
	"context"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

// ResourceResolver provides a mechanism to resolve the resulting workspace of a types.ConnectionTarget
type ResourceResolver struct {
	ApiClient    *api.Client
	CurStackId   int64
	CurEnvId     int64
	StacksById   map[int64]*StackResolver
	StacksByName map[string]*StackResolver
}

func NewResourceResolver(apiClient *api.Client, curStackId, curEnvId int64) *ResourceResolver {
	return &ResourceResolver{
		ApiClient:    apiClient,
		CurStackId:   curStackId,
		CurEnvId:     curEnvId,
		StacksById:   map[int64]*StackResolver{},
		StacksByName: map[string]*StackResolver{},
	}
}

func (r *ResourceResolver) Resolve(ctx context.Context, ct types.ConnectionTarget) (types.ConnectionTarget, error) {
	result := ct

	sr, err := r.ResolveStack(ctx, result)
	if err != nil {
		return result, err
	}
	result.StackId = sr.Stack.Id
	result.StackName = sr.Stack.Name

	currentEnv, err := sr.ResolveEnvById(ctx, r.CurEnvId)
	if err != nil {
		return result, err
	}
	configuredEnv, err := sr.ResolveEnv(ctx, result)
	if err != nil {
		return result, err
	}

	block, err := sr.ResolveBlock(ctx, result)
	if err != nil {
		return result, err
	}
	result.BlockId = block.Id
	result.BlockName = block.Name

	sharedEnv := r.resolveSharedEnv(block, sr)
	result = r.setEnv(result, configuredEnv, currentEnv, sharedEnv)

	return result, nil
}

func (r *ResourceResolver) FindBlock(ctx context.Context, ct types.ConnectionTarget) (types.Block, error) {
	result := ct
	sr, err := r.ResolveStack(ctx, result)
	if err != nil {
		return types.Block{}, err
	}
	return sr.ResolveBlock(ctx, result)
}

func (r *ResourceResolver) ResolveStack(ctx context.Context, ct types.ConnectionTarget) (*StackResolver, error) {
	if ct.StackName != "" {
		return r.resolveStackByName(ctx, ct.StackName)
	}
	if ct.StackId == 0 {
		ct.StackId = r.CurStackId
	}
	return r.resolveStackById(ctx, ct.StackId)
}

func (r *ResourceResolver) ResolveCurProviderType(ctx context.Context) (string, error) {
	sr, err := r.ResolveStack(ctx, types.ConnectionTarget{StackId: r.CurStackId})
	if err != nil {
		return "", err
	}
	return sr.Stack.ProviderType, nil
}

func (r *ResourceResolver) BackfillMissingBlocks(ctx context.Context, blocks []types.Block) error {
	sr, err := r.ResolveStack(ctx, types.ConnectionTarget{StackId: r.CurStackId})
	if err != nil {
		return fmt.Errorf("unable to resolve stack: %w", err)
	}

	for _, block := range blocks {
		block.StackId = r.CurStackId
		if err = sr.AddBlock(ctx, block); err != nil {
			return fmt.Errorf("unable to add block (%s) to resolver: %w", block.Name, err)
		}
	}

	return nil
}

func (r *ResourceResolver) GetCurrentEnvs(ctx context.Context) (map[int64]types.Environment, error) {
	sr, err := r.ResolveStack(ctx, types.ConnectionTarget{StackId: r.CurStackId})
	if err != nil {
		return nil, fmt.Errorf("unable to resolve stack: %w", err)
	}
	return sr.Envs(ctx)
}

func (r *ResourceResolver) resolveStackByName(ctx context.Context, stackName string) (*StackResolver, error) {
	if sr, ok := r.StacksByName[stackName]; ok {
		return sr, nil
	}
	if err := r.loadStacks(ctx); err != nil {
		return nil, err
	}
	if sr, ok := r.StacksByName[stackName]; ok {
		return sr, nil
	}
	return nil, StackDoesNotExistError{StackName: stackName}
}

func (r *ResourceResolver) resolveStackById(ctx context.Context, stackId int64) (*StackResolver, error) {
	if sr, ok := r.StacksById[stackId]; ok {
		return sr, nil
	}
	if err := r.loadStacks(ctx); err != nil {
		return nil, err
	}
	if sr, ok := r.StacksById[stackId]; ok {
		return sr, nil
	}
	return nil, StackIdDoesNotExistError{StackId: stackId}
}

func (r *ResourceResolver) loadStacks(ctx context.Context) error {
	stacks, err := r.ApiClient.Stacks().List(ctx)
	if err != nil {
		return err
	}
	for _, stack := range stacks {
		sr := &StackResolver{ApiClient: r.ApiClient, Stack: *stack}
		r.StacksById[stack.Id] = sr
		r.StacksByName[stack.Name] = sr
	}
	return nil
}

// resolveSharedBlockEnv performs resolution of shared blocks
// When a block is marked shared, it is created once for all preview envs and stored in a special environment (i.e. `previews-shared`)
// We only perform this resolution under the following circumstances:
// 1. The block is marked 'shared'
// 2. Our stack contains a `previews-shared` env
func (r *ResourceResolver) resolveSharedEnv(curBlock types.Block, curStackResolver *StackResolver) *types.Environment {
	if !curBlock.IsShared {
		// Block is not marked shared
		return nil
	}
	if curStackResolver.PreviewsSharedEnvId == 0 {
		// The stack doesn't have a `previews-shared` env
		return nil
	}

	env := curStackResolver.EnvsById[curStackResolver.PreviewsSharedEnvId]
	return &env
}

// resolveEnv determines whether an env should be set on the connection target
//
//	We only need to modify the env for preview envs, pipeline envs can't be normalized
func (r *ResourceResolver) setEnv(ct types.ConnectionTarget, configuredEnv *types.Environment, currentEnv types.Environment, sharedEnv *types.Environment) types.ConnectionTarget {
	// for pipeline environments, we don't do normalization so just return the configured env
	//   if there is no configured env, just return the current env
	if currentEnv.Type == types.EnvTypePipeline {
		if configuredEnv == nil {
			ct.EnvId = &currentEnv.Id
			ct.EnvName = currentEnv.Name
			return ct
		} else {
			ct.EnvId = &configuredEnv.Id
			ct.EnvName = configuredEnv.Name
			return ct
		}
	}

	// for preview envs, if there is no shared env, we resolve to the current env
	if sharedEnv == nil {
		ct.EnvId = &currentEnv.Id
		ct.EnvName = currentEnv.Name
		return ct
	} else {
		// for preview envs, if there is a shared env, we resolve to the shared
		ct.EnvId = &sharedEnv.Id
		ct.EnvName = sharedEnv.Name
		return ct
	}
}
