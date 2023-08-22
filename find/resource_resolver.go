package find

import (
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

func (r *ResourceResolver) Resolve(ct types.ConnectionTarget) (types.ConnectionTarget, error) {
	result := ct

	sr, err := r.ResolveStack(result)
	if err != nil {
		return result, err
	}
	result.StackId = sr.Stack.Id
	result.StackName = sr.Stack.Name

	env, err := sr.ResolveEnv(result, r.CurEnvId)
	if err != nil {
		return result, err
	}
	envId := env.Id
	result.EnvId = &envId
	result.EnvName = env.Name

	block, err := sr.ResolveBlock(result)
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

func (r *ResourceResolver) FindBlock(ct types.ConnectionTarget) (types.Block, error) {
	result := ct
	sr, err := r.ResolveStack(result)
	if err != nil {
		return types.Block{}, err
	}
	return sr.ResolveBlock(result)
}

func (r *ResourceResolver) ResolveStack(ct types.ConnectionTarget) (*StackResolver, error) {
	if ct.StackName != "" {
		return r.resolveStackByName(ct.StackName)
	}
	if ct.StackId == 0 {
		ct.StackId = r.CurStackId
	}
	return r.resolveStackById(ct.StackId)
}

func (r *ResourceResolver) ResolveCurProviderType() (string, error) {
	sr, err := r.ResolveStack(types.ConnectionTarget{StackId: r.CurStackId})
	if err != nil {
		return "", err
	}
	return sr.Stack.ProviderType, nil
}

func (r *ResourceResolver) resolveStackByName(stackName string) (*StackResolver, error) {
	if sr, ok := r.StacksByName[stackName]; ok {
		return sr, nil
	}
	if err := r.loadStacks(); err != nil {
		return nil, err
	}
	if sr, ok := r.StacksByName[stackName]; ok {
		return sr, nil
	}
	return nil, StackDoesNotExistError{StackName: stackName}
}

func (r *ResourceResolver) resolveStackById(stackId int64) (*StackResolver, error) {
	if sr, ok := r.StacksById[stackId]; ok {
		return sr, nil
	}
	if err := r.loadStacks(); err != nil {
		return nil, err
	}
	if sr, ok := r.StacksById[stackId]; ok {
		return sr, nil
	}
	return nil, StackIdDoesNotExistError{StackId: stackId}
}

func (r *ResourceResolver) loadStacks() error {
	stacks, err := r.ApiClient.Stacks().List()
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
