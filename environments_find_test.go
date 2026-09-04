package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

func TestFindEnvironmentsInput_query(t *testing.T) {
	t.Run("zero value sends no params", func(t *testing.T) {
		assert.Empty(t, FindEnvironmentsInput{}.query().Encode())
	})

	t.Run("encodes every filter, tags sorted and repeated", func(t *testing.T) {
		isProd := false
		q := FindEnvironmentsInput{
			Types:  []types.EnvironmentType{types.EnvTypePreview, types.EnvTypePipeline},
			Status: types.EnvStatusArchived,
			IsProd: &isProd,
			Search: "pr-*",
			Tags:   map[string]string{"tier": "gold", "claim": ""},
		}.query()

		assert.Equal(t, "PreviewEnv,PipelineEnv", q.Get("type"))
		assert.Equal(t, "archived", q.Get("status"))
		assert.Equal(t, "false", q.Get("is_prod"))
		assert.Equal(t, "pr-*", q.Get("search"))
		assert.Equal(t, []string{"claim=", "tier=gold"}, q["tag"], "empty value is sent as KEY= so the API can tell it apart from an absent filter")
	})
}
