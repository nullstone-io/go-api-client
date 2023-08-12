package find

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"gopkg.in/nullstone-io/go-api-client.v0/mocks"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"testing"
)

func TestResourceResolver(t *testing.T) {
	stack1 := types.Stack{
		IdModel:      types.IdModel{Id: 1},
		OrgName:      "nullstone",
		Name:         "primary",
		ProviderType: "aws",
	}
	stack2 := types.Stack{
		IdModel:      types.IdModel{Id: 2},
		OrgName:      "nullstone",
		Name:         "secondary",
		ProviderType: "aws",
	}
	stack3Id := int64(3)
	env1 := types.Environment{
		IdModel:   types.IdModel{Id: 11},
		Type:      types.EnvTypePipeline,
		Name:      "dev",
		OrgName:   "nullstone",
		StackId:   stack1.Id,
		Reference: "purple-parrot",
	}
	env2 := types.Environment{
		IdModel:   types.IdModel{Id: 12},
		Type:      types.EnvTypePipeline,
		Name:      "prod",
		OrgName:   "nullstone",
		StackId:   stack2.Id,
		Reference: "orange-iguana",
	}
	env3Id := int64(13)
	block1 := types.Block{
		IdModel:  types.IdModel{Id: 101},
		Type:     "block",
		OrgName:  "nullstone",
		StackId:  stack1.Id,
		Name:     "block1",
		IsShared: false,
	}
	block2 := types.Block{
		IdModel:  types.IdModel{Id: 102},
		Type:     "block",
		OrgName:  "nullstone",
		StackId:  stack1.Id,
		Name:     "block2",
		IsShared: false,
	}
	block3Id := int64(103)

	stacks := []types.Stack{stack1, stack2}
	envs := []types.Environment{env1, env2}
	blocks := []types.Block{block1, block2}
	router := mux.NewRouter()
	mocks.ListStacks(router, stacks)
	mocks.ListEnvironments(router, envs)
	mocks.ListBlocks(router, blocks)
	apiClient := mocks.Client(t, "nullstone", router)

	rr := NewResourceResolver(apiClient, stack1.Id, env1.Id)

	t.Run("stack does not exist", func(t *testing.T) {
		ct := types.ConnectionTarget{
			StackId: stack3Id,
			BlockId: block3Id,
			EnvId:   &env3Id,
		}
		want := ct
		got, err := rr.Resolve(ct)
		assert.ErrorIs(t, err, StackIdDoesNotExistError{StackId: stack3Id})
		assert.Equal(t, want, got)
	})
	t.Run("env does not exist", func(t *testing.T) {
		ct := types.ConnectionTarget{
			StackId: stack1.Id,
			BlockId: block3Id,
			EnvId:   &env3Id,
		}
		want := ct
		want.StackName = stack1.Name
		got, err := rr.Resolve(ct)
		assert.ErrorIs(t, err, EnvIdDoesNotExistError{StackName: "primary", EnvId: env3Id})
		assert.Equal(t, want, got)
	})
	t.Run("block does not exist", func(t *testing.T) {
		ct := types.ConnectionTarget{
			StackId: stack1.Id,
			BlockId: block3Id,
			EnvId:   &env1.Id,
		}
		want := ct
		want.StackName = stack1.Name
		want.EnvName = env1.Name
		got, err := rr.Resolve(ct)
		assert.ErrorIs(t, err, BlockIdDoesNotExistError{StackName: "primary", BlockId: block3Id})
		assert.Equal(t, want, got)
	})
	t.Run("load successfully", func(t *testing.T) {
		ct := types.ConnectionTarget{
			StackId: stack1.Id,
			BlockId: block1.Id,
			EnvId:   &env1.Id,
		}
		want := ct
		want.StackName = stack1.Name
		want.EnvName = env1.Name
		want.BlockName = block1.Name
		got, err := rr.Resolve(ct)
		assert.NoError(t, err, "unexpected error")
		assert.Equal(t, want, got)
	})
	t.Run("load again", func(t *testing.T) {
		ct := types.ConnectionTarget{
			StackId: stack1.Id,
			BlockId: block1.Id,
			EnvId:   &env1.Id,
		}
		want := ct
		want.StackName = stack1.Name
		want.EnvName = env1.Name
		want.BlockName = block1.Name
		got, err := rr.Resolve(ct)
		assert.NoError(t, err, "unexpected error")
		assert.Equal(t, want, got)
	})
}
