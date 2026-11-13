package api_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	api "gopkg.in/nullstone-io/go-api-client.v0"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

const envGitopsConfigsBody = `[
  {
    "gitopsConnectionId": 3,
    "envId": 9,
    "isEnabled": true,
    "isAutoApplyEnabled": true,
    "isPreview": true,
    "triggerType": "pull_request",
    "gitBranch": "",
    "pullRequestId": 547,
    "tagFilter": "",
    "createdBy": "brad",
    "updatedBy": "sam",
    "updatedAt": "2026-11-13T14:30:00Z",
    "pullRequest": {"id": 547, "repo": "nullstone/widgets", "number": 68, "state": "open"}
  },
  {
    "gitopsConnectionId": 4,
    "envId": 9,
    "isEnabled": true,
    "isAutoApplyEnabled": true,
    "isPreview": true,
    "triggerType": "tag",
    "gitBranch": "",
    "pullRequestId": null,
    "tagFilter": "v*",
    "createdBy": "",
    "updatedBy": "",
    "updatedAt": "2026-11-13T14:30:00Z"
  }
]`

func TestEnvGitopsConfigs_List(t *testing.T) {
	client, rec := recordingClient(t, http.StatusOK, envGitopsConfigsBody)

	got, err := client.EnvGitopsConfigs().List(context.Background(), 42, 9)

	require.NoError(t, err)
	assert.Equal(t, http.MethodGet, rec.method)
	assert.Equal(t, "/orgs/nullstone/stacks/42/envs/9/gitops_configs", rec.path)
	require.Len(t, got, 2)

	pullRequestId := int64(547)
	first := got[0]
	assert.True(t, first.IsPreview)
	assert.Equal(t, types.GitopsTrigger{TriggerType: types.GitopsTriggerTypePullRequest, PullRequestId: &pullRequestId}, first.Trigger())
	assert.Equal(t, "brad", first.CreatedBy)
	assert.Equal(t, "sam", first.UpdatedBy)
	assert.Equal(t, time.Date(2026, 11, 13, 14, 30, 0, 0, time.UTC), first.UpdatedAt)
	require.NotNil(t, first.PullRequest)
	assert.Equal(t, 68, first.PullRequest.Number)

	second := got[1]
	assert.Equal(t, types.GitopsTrigger{TriggerType: types.GitopsTriggerTypeTag, TagFilter: "v*"}, second.Trigger())
	assert.Nil(t, second.PullRequest)
	assert.Equal(t, "", second.CreatedBy, "a config written before authors were recorded")
}

func TestEnvGitopsConfigs_Replace(t *testing.T) {
	client, rec := recordingClient(t, http.StatusOK, envGitopsConfigsBody)
	pullRequestId := int64(547)
	input := []api.SetEnvGitopsTriggerInput{
		{RepoUrl: "https://github.com/nullstone/widgets", GitopsTrigger: types.GitopsTrigger{PullRequestId: &pullRequestId}},
		{RepoUrl: "https://github.com/nullstone/gadgets", GitopsTrigger: types.GitopsTrigger{TriggerType: types.GitopsTriggerTypeTag, TagFilter: "v*"}},
	}

	got, err := client.EnvGitopsConfigs().Replace(context.Background(), 42, 9, input)

	require.NoError(t, err)
	assert.Equal(t, http.MethodPut, rec.method)
	assert.Equal(t, "/orgs/nullstone/stacks/42/envs/9/gitops_configs", rec.path)
	assert.JSONEq(t, `[
  {"repoUrl": "https://github.com/nullstone/widgets", "triggerType": "", "gitBranch": "", "pullRequestId": 547, "tagFilter": ""},
  {"repoUrl": "https://github.com/nullstone/gadgets", "triggerType": "tag", "gitBranch": "", "pullRequestId": null, "tagFilter": "v*"}
]`, string(rec.body))
	require.Len(t, got, 2)
	assert.Equal(t, int64(9), got[0].EnvId)
}
