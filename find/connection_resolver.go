package find

import (
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

// ConnectionResolver provides a mechanism to resolve the resulting workspace of a types.ConnectionTarget
type ConnectionResolver struct {
	ApiClient    *api.Client
	OrgName      string
	CurStackId   int64
	CurEnvId     int64
	StacksById   map[int64]*StackResolver
	StacksByName map[string]*StackResolver
}

// NewPreloadedConnectionResolver initializes a new ConnectionResolver by:
// 1. preloading all stacks accessible in the org
// 2. preloading all envs in the requested stack
// 3. preloading all blocks in the requested stack
func NewPreloadedConnectionResolver(apiClient *api.Client, orgName string, curStackId, curEnvId int64) (ConnectionResolver, error) {
	resources := ConnectionResolver{
		ApiClient:  apiClient,
		OrgName:    orgName,
		CurStackId: curStackId,
		CurEnvId:   curEnvId,
	}

	stacks, err := apiClient.Stacks().List()
	if err != nil {
		return resources, fmt.Errorf("unable to fetch stacks for org (org=%s): %w", orgName, err)
	}

	resources.StacksById = map[int64]*StackResolver{}
	resources.StacksByName = map[string]*StackResolver{}
	for _, stack := range stacks {
		sr := &StackResolver{
			Stack:        *stack,
			EnvsByName:   nil,
			BlocksByName: nil,
		}
		if stack.Id == curStackId {
			// Load envs for current stack
			if err := sr.LoadEnvs(apiClient, orgName); err != nil {
				return resources, err
			}
			// Load blocks for current stack
			if err := sr.LoadBlocks(apiClient, orgName); err != nil {
				return resources, err
			}
		}
		if stack.Name == "global" {
			// Preload the global stack/env
			if err := sr.LoadEnvs(apiClient, orgName); err != nil {
				return resources, err
			}
		}
		resources.StacksById[stack.Id] = sr
		resources.StacksByName[stack.Name] = sr
	}
	return resources, nil
}

func (r ConnectionResolver) Resolve(ct types.ConnectionTarget) (types.ConnectionTarget, error) {
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
	if block.IsShared && sr.PreviewsSharedEnvId != 0 {
		envId := sr.PreviewsSharedEnvId
		result.EnvId = &envId
	}

	return result, nil
}

func (r ConnectionResolver) ResolveStack(ct types.ConnectionTarget) (*StackResolver, error) {
	if ct.StackName != "" {
		sr, ok := r.StacksByName[ct.StackName]
		if !ok {
			return nil, StackDoesNotExistError{StackName: ct.StackName}
		}
		return sr, nil
	}
	if ct.StackId == 0 {
		ct.StackId = r.CurStackId
	}
	sr, ok := r.StacksById[ct.StackId]
	if !ok {
		return nil, StackIdDoesNotExistError{StackId: ct.StackId}
	}
	return sr, nil
}
