package mocks

import (
	"context"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

type EnvironmentStore struct {
	Envs []types.Environment
}

func (s EnvironmentStore) ListEnvs(ctx context.Context, stackId int64) ([]types.Environment, error) {
	result := make([]types.Environment, 0)
	for _, env := range s.Envs {
		if env.StackId == stackId {
			result = append(result, env)
		}
	}
	return result, nil
}

func (s EnvironmentStore) GetEnvById(ctx context.Context, stackId int64, envId int64) (*types.Environment, error) {
	for _, env := range s.Envs {
		if env.StackId == stackId && env.Id == envId {
			return &env, nil
		}
	}
	return nil, nil
}

func (s EnvironmentStore) GetEnvByName(ctx context.Context, stackId int64, envName string) (*types.Environment, error) {
	for _, env := range s.Envs {
		if env.StackId == stackId && env.Name == envName {
			return &env, nil
		}
	}
	return nil, nil
}
