package find

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"gopkg.in/nullstone-io/go-api-client.v0/mocks"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"testing"
)

func TestStackResolver_ResolveEnv(t *testing.T) {
	stack1Id := int64(1)
	env1 := types.Environment{
		IdModel:   types.IdModel{Id: 11},
		Type:      types.EnvTypePipeline,
		Name:      "dev",
		OrgName:   "nullstone",
		StackId:   stack1Id,
		Reference: "purple-parrot",
	}
	env2 := types.Environment{
		IdModel:   types.IdModel{Id: 12},
		Type:      types.EnvTypePipeline,
		Name:      "prod",
		OrgName:   "nullstone",
		StackId:   stack1Id,
		Reference: "orange-iguana",
	}
	env3 := types.Environment{
		IdModel:   types.IdModel{Id: 13},
		Type:      types.EnvTypePipeline,
		Name:      "missing",
		OrgName:   "nullstone",
		StackId:   stack1Id,
		Reference: "green-monkey",
	}

	envs := mocks.EnvironmentStore{Envs: []types.Environment{env1, env2}}
	router := mux.NewRouter()
	apiClient := mocks.Client(t, "nullstone", router)

	sr := StackResolver{
		ApiClient: apiClient,
		EnvGetter: envs,
		Stack: types.Stack{
			IdModel:      types.IdModel{Id: stack1Id},
			Reference:    "red-jaguar",
			Name:         "primary",
			OrgName:      "nullstone",
			Description:  "Primary Stack",
			ProviderType: "aws",
		},
		PreviewsSharedEnvId: 0,
		EnvsById:            map[int64]types.Environment{},
		EnvsByName:          map[string]types.Environment{},
		BlocksById:          map[int64]types.Block{},
		BlocksByName:        map[string]types.Block{},
	}

	t.Run("does not exist", func(t *testing.T) {
		ct := types.ConnectionTarget{
			StackId: stack1Id,
			EnvId:   &env3.Id,
		}
		got, gotErr := sr.ResolveEnv(context.Background(), ct, env1.Id)
		assert.ErrorIs(t, gotErr, EnvIdDoesNotExistError{StackName: "primary", EnvId: env3.Id})
		assert.Equal(t, types.Environment{}, got, "env should be empty")
	})
	t.Run("needs loaded first", func(t *testing.T) {
		ct := types.ConnectionTarget{
			StackId: stack1Id,
			EnvId:   &env1.Id,
		}
		got, err := sr.ResolveEnv(context.Background(), ct, env1.Id)
		assert.NoError(t, err)
		assert.Equal(t, env1, got, "should resolve env1")
	})
	t.Run("already loaded", func(t *testing.T) {
		ct := types.ConnectionTarget{
			StackId: stack1Id,
			EnvId:   &env1.Id,
		}
		got, err := sr.ResolveEnv(context.Background(), ct, env1.Id)
		assert.NoError(t, err)
		assert.Equal(t, env1, got, "should resolve env1")
	})
	t.Run("loads current env", func(t *testing.T) {
		ct := types.ConnectionTarget{
			StackId: stack1Id,
		}
		got, err := sr.ResolveEnv(context.Background(), ct, env1.Id)
		assert.NoError(t, err)
		assert.Equal(t, env1, got, "should resolve env1")
	})
}

func TestStackResolver_ResolveBlock(t *testing.T) {
	stack1Id := int64(1)
	block1 := types.Block{
		IdModel:  types.IdModel{Id: 101},
		Type:     "block",
		OrgName:  "nullstone",
		StackId:  stack1Id,
		Name:     "block1",
		IsShared: false,
	}
	block2 := types.Block{
		IdModel:  types.IdModel{Id: 102},
		Type:     "block",
		OrgName:  "nullstone",
		StackId:  stack1Id,
		Name:     "block2",
		IsShared: false,
	}
	block3 := types.Block{
		IdModel:  types.IdModel{Id: 103},
		Type:     "block",
		OrgName:  "nullstone",
		StackId:  stack1Id,
		Name:     "block3",
		IsShared: false,
	}
	blocks := []types.Block{block1, block2}

	envs := mocks.EnvironmentStore{Envs: []types.Environment{}}
	router := mux.NewRouter()
	mocks.ListBlocks(router, blocks)
	apiClient := mocks.Client(t, "nullstone", router)

	sr := StackResolver{
		ApiClient: apiClient,
		EnvGetter: envs,
		Stack: types.Stack{
			IdModel:      types.IdModel{Id: stack1Id},
			Reference:    "red-jaguar",
			Name:         "primary",
			OrgName:      "nullstone",
			Description:  "Primary Stack",
			ProviderType: "aws",
		},
		PreviewsSharedEnvId: 0,
		EnvsById:            map[int64]types.Environment{},
		EnvsByName:          map[string]types.Environment{},
		BlocksById:          map[int64]types.Block{},
		BlocksByName:        map[string]types.Block{},
	}

	t.Run("does not exist", func(t *testing.T) {
		ct := types.ConnectionTarget{
			StackId: stack1Id,
			BlockId: block3.Id,
		}
		got, gotErr := sr.ResolveBlock(context.Background(), ct)
		assert.ErrorIs(t, gotErr, BlockIdDoesNotExistError{StackName: "primary", BlockId: block3.Id})
		assert.Equal(t, types.Block{}, got, "block should be empty")
	})
	t.Run("needs loaded first", func(t *testing.T) {
		ct := types.ConnectionTarget{
			StackId: stack1Id,
			BlockId: block1.Id,
		}
		got, err := sr.ResolveBlock(context.Background(), ct)
		assert.NoError(t, err)
		assert.Equal(t, block1, got, "should resolve block1")
	})
	t.Run("already loaded", func(t *testing.T) {
		ct := types.ConnectionTarget{
			StackId: stack1Id,
			BlockId: block1.Id,
		}
		got, err := sr.ResolveBlock(context.Background(), ct)
		assert.NoError(t, err)
		assert.Equal(t, block1, got, "should resolve block1")
	})
}
