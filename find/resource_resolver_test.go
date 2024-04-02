package find

import (
	"context"
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
	env4 := types.Environment{
		IdModel:   types.IdModel{Id: 14},
		Type:      types.EnvTypePreview,
		Name:      "f-1103-add-something",
		OrgName:   "nullstone",
		StackId:   stack2.Id,
		Reference: "purple-snail",
	}
	env5 := types.Environment{
		IdModel:   types.IdModel{Id: 15},
		Type:      types.EnvTypePreviewsShared,
		Name:      "previews-shared",
		OrgName:   "nullstone",
		StackId:   stack2.Id,
		Reference: "teal-bear",
	}
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
	block4 := types.Block{
		IdModel:  types.IdModel{Id: 104},
		Type:     "block",
		OrgName:  "nullstone",
		StackId:  stack2.Id,
		Name:     "block4",
		IsShared: true,
	}
	block5 := types.Block{
		IdModel:  types.IdModel{Id: 105},
		Type:     "block",
		OrgName:  "nullstone",
		StackId:  stack2.Id,
		Name:     "block5",
		IsShared: false,
	}

	stacks := []types.Stack{stack1, stack2}
	envs := []types.Environment{env1, env2, env4, env5}
	blocks := []types.Block{block1, block2, block4, block5}
	router := mux.NewRouter()
	mocks.ListStacks(router, stacks)
	mocks.ListEnvironments(router, envs)
	mocks.ListBlocks(router, blocks)
	apiClient := mocks.Client(t, "nullstone", router)

	rr := NewResourceResolver(apiClient, stack1.Id, env1.Id)
	rr2 := NewResourceResolver(apiClient, stack2.Id, env4.Id)

	t.Run("stack does not exist", func(t *testing.T) {
		ct := types.ConnectionTarget{
			StackId: stack3Id,
			BlockId: block3Id,
			EnvId:   &env3Id,
		}
		_, err := rr.Resolve(context.Background(), ct)
		assert.ErrorIs(t, err, StackIdDoesNotExistError{StackId: stack3Id})
	})
	t.Run("env does not exist", func(t *testing.T) {
		ct := types.ConnectionTarget{
			StackId: stack1.Id,
			BlockId: block2.Id,
			EnvId:   &env3Id,
		}
		_, err := rr.Resolve(context.Background(), ct)
		assert.ErrorIs(t, err, EnvIdDoesNotExistError{StackName: "primary", EnvId: env3Id})
	})
	t.Run("block does not exist", func(t *testing.T) {
		ct := types.ConnectionTarget{
			StackId: stack1.Id,
			BlockId: block3Id,
			EnvId:   &env1.Id,
		}
		_, err := rr.Resolve(context.Background(), ct)
		assert.ErrorIs(t, err, BlockIdDoesNotExistError{StackName: "primary", BlockId: block3Id})
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
		got, err := rr.Resolve(context.Background(), ct)
		assert.NoError(t, err, "unexpected error")
		assert.Equal(t, want, got)
	})
	t.Run("skip previews-shared because it does not exist", func(t *testing.T) {
		ct := types.ConnectionTarget{
			StackId: stack1.Id,
			BlockId: block2.Id,
		}
		want := ct
		want.StackName = stack1.Name
		want.EnvId = &env1.Id
		want.EnvName = env1.Name
		want.BlockName = block2.Name
		got, err := rr.Resolve(context.Background(), ct)
		assert.NoError(t, err, "unexpected error")
		assert.Equal(t, want, got)
	})
	t.Run("block is not shared, reset env to current env", func(t *testing.T) {
		ct := types.ConnectionTarget{
			StackId: stack2.Id,
			BlockId: block5.Id,
			EnvId:   &env5.Id,
		}
		want := ct
		want.StackName = stack2.Name
		want.EnvId = &env4.Id
		want.EnvName = env4.Name
		want.BlockName = block5.Name
		got, err := rr2.Resolve(context.Background(), ct)
		assert.NoError(t, err, "unexpected error")
		assert.Equal(t, want, got)
	})
	t.Run("load previews-shared", func(t *testing.T) {
		ct := types.ConnectionTarget{
			StackId: stack2.Id,
			BlockId: block4.Id,
		}
		want := ct
		want.StackName = stack2.Name
		want.EnvId = &env5.Id
		want.EnvName = env5.Name
		want.BlockName = block4.Name
		got, err := rr2.Resolve(context.Background(), ct)
		assert.NoError(t, err, "unexpected error")
		assert.Equal(t, want, got)
	})
}

func TestResourceResolver_BackfillMissingBlocks(t *testing.T) {
	stack1 := types.Stack{
		IdModel:      types.IdModel{Id: 1},
		OrgName:      "nullstone",
		Name:         "primary",
		ProviderType: "aws",
	}
	env1 := types.Environment{
		IdModel:   types.IdModel{Id: 11},
		Type:      types.EnvTypePipeline,
		Name:      "dev",
		OrgName:   "nullstone",
		StackId:   stack1.Id,
		Reference: "purple-parrot",
	}
	block1 := types.Block{
		IdModel:  types.IdModel{Id: 101},
		Type:     "Block",
		OrgName:  "nullstone",
		StackId:  stack1.Id,
		Name:     "block1",
		IsShared: false,
	}
	block2 := types.Block{
		IdModel:  types.IdModel{Id: 102},
		Type:     "Block",
		OrgName:  "nullstone",
		StackId:  stack1.Id,
		Name:     "block2",
		IsShared: false,
	}

	stacks := []types.Stack{stack1}
	envs := []types.Environment{env1}
	blocks := []types.Block{block1, block2}
	router := mux.NewRouter()
	mocks.ListStacks(router, stacks)
	mocks.ListEnvironments(router, envs)
	mocks.ListBlocks(router, blocks)
	apiClient := mocks.Client(t, "nullstone", router)

	rr := NewResourceResolver(apiClient, stack1.Id, env1.Id)

	t.Run("adds all blocks", func(t *testing.T) {
		newBlocks := []types.Block{
			{
				Type:     "Block",
				OrgName:  "nullstone",
				StackId:  0,
				Name:     "block3",
				IsShared: false,
			},
			{
				Type:     "Block",
				OrgName:  "nullstone",
				StackId:  0,
				Name:     "block1",
				IsShared: true,
			},
			{
				IdModel:  types.IdModel{Id: 102},
				Type:     "Block",
				OrgName:  "nullstone",
				StackId:  0,
				Name:     "block2",
				IsShared: true,
			},
		}
		ctx := context.Background()
		err := rr.BackfillMissingBlocks(ctx, newBlocks)
		assert.NoError(t, err)

		sr, err := rr.ResolveStack(ctx, types.ConnectionTarget{StackId: stack1.Id})
		assert.NoError(t, err)

		assert.Equal(t, 3, len(sr.BlocksByName))
		assert.Equal(t, 2, len(sr.BlocksById))

		// check to make sure "block3" was added and StackId set
		assert.Equal(t, sr.BlocksByName["block3"].StackId, stack1.Id)

		// check to make sure "block1" was updated because the name was the same
		assert.Equal(t, sr.BlocksByName["block1"].IsShared, true)

		// check to make sure "block2" was updated because the id was the same
		assert.Equal(t, sr.BlocksById[102].IsShared, true)
	})
}
