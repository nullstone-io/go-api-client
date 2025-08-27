package find

import (
	"context"
	"encoding/json"
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

	if !ct.HasEnv() {
		// If the connection target doesn't have env id or env name, we need to find the env by name in the current stack
		curStackResolver, err := r.ResolveStack(ctx, types.ConnectionTarget{StackId: r.CurStackId})
		if err != nil {
			return result, err
		}
		matchEnv, err := curStackResolver.ResolveEnv(ctx, result, r.CurEnvId)
		if err != nil {
			return result, err
		}
		result.EnvName = matchEnv.Name
	}

	env, err := sr.ResolveEnv(ctx, result, r.CurEnvId)
	if err != nil {
		return result, err
	}
	envId := env.Id
	result.EnvId = &envId
	result.EnvName = env.Name

	block, err := sr.ResolveBlock(ctx, result)
	if err != nil {
		return result, err
	}
	result.BlockId = block.Id
	result.BlockName = block.Name

	sharedEnv := r.resolveSharedBlockEnv(ct, block, env, sr)
	if sharedEnv != nil {
		sharedEnvId := sharedEnv.Id
		result.EnvId = &sharedEnvId
		result.EnvName = sharedEnv.Name
	}

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

func (r *ResourceResolver) FindEnv(ctx context.Context, ct types.ConnectionTarget) (types.Environment, error) {
	result := ct
	sr, err := r.ResolveStack(ctx, result)
	if err != nil {
		return types.Environment{}, err
	}
	return sr.ResolveEnv(ctx, result, r.CurEnvId)
}

func (r *ResourceResolver) ResolveWorkspaceDetails(ctx context.Context, ct types.ConnectionTarget) (types.WorkspaceDetails, error) {
	wd := types.WorkspaceDetails{}

	result := ct
	sr, err := r.ResolveStack(ctx, result)
	if err != nil {
		return wd, err
	}
	wd.Stack = sr.Stack
	wd.OrgName = wd.Stack.OrgName

	block, err := sr.ResolveBlock(ctx, ct)
	if err != nil {
		return wd, err
	}
	wd.BlockRaw, _ = json.Marshal(block)
	env, err := sr.ResolveEnv(ctx, ct, r.CurEnvId)
	if err != nil {
		return wd, err
	}
	wd.Env = env
	return wd, nil
}

func (r *ResourceResolver) ResolveStack(ctx context.Context, ct types.ConnectionTarget) (*StackResolver, error) {
	// Prefer StackId over StackName -- it's possible for a user to rename the stack
	if ct.StackId != 0 {
		return r.resolveStackById(ctx, ct.StackId)
	}
	if ct.StackName != "" {
		return r.resolveStackByName(ctx, ct.StackName)
	}
	return r.resolveStackById(ctx, r.CurStackId)
}

func (r *ResourceResolver) ResolveCurProviderType(ctx context.Context) (string, error) {
	sr, err := r.ResolveStack(ctx, types.ConnectionTarget{StackId: r.CurStackId})
	if err != nil {
		return "", err
	}
	return sr.Stack.ProviderType, nil
}

func (r *ResourceResolver) BackfillMissingBlocks(ctx context.Context, blocks []types.Block) ([]types.Block, error) {
	sr, err := r.ResolveStack(ctx, types.ConnectionTarget{StackId: r.CurStackId})
	if err != nil {
		return nil, fmt.Errorf("unable to resolve stack: %w", err)
	}

	missing := make([]types.Block, 0)
	for _, block := range blocks {
		block.StackId = r.CurStackId
		if added, err := sr.AddBlock(ctx, block); err != nil {
			return nil, fmt.Errorf("unable to add block (%s) to resolver: %w", block.Name, err)
		} else if added {
			missing = append(missing, block)
		}
	}

	return missing, nil
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
// 1. The current env must be a preview environment
// 2. The block is marked 'shared'
// 3. The user did not specify an explicit environment
// 4. Our stack contains a `previews-shared` env
func (r *ResourceResolver) resolveSharedBlockEnv(original types.ConnectionTarget, curBlock types.Block, curEnv types.Environment, curStackResolver *StackResolver) *types.Environment {
	if curEnv.Type != types.EnvTypePreview {
		// Current env is not a preview env
		return nil
	}
	if !curBlock.IsShared {
		// Block is not marked shared
		return nil
	}
	if original.EnvId != nil {
		// User specified an explicit environment
		return nil
	}
	if curStackResolver.PreviewsSharedEnvId == 0 {
		// The stack doesn't have a `previews-shared` env
		return nil
	}

	env := curStackResolver.EnvsById[curStackResolver.PreviewsSharedEnvId]
	return &env
}
