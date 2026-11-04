package api_test

import (
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	api "gopkg.in/nullstone-io/go-api-client.v0"
	"gopkg.in/nullstone-io/go-api-client.v0/mocks"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

const blockAutomationDefaultsBody = `{
  "orgName": "nullstone",
  "stackId": 42,
  "blockId": 7,
  "isEnabled": true,
  "autoLaunch": false,
  "autoDeploy": true,
  "defaultBranch": "main",
  "repo": "nullstone/widgets",
  "repoUrl": "https://github.com/nullstone/widgets",
  "gitopsConnectionId": 3
}`

const workspaceAutomationBody = `{
  "orgName": "nullstone",
  "stackId": 42,
  "blockId": 7,
  "envId": 9,
  "isEnabled": true,
  "autoLaunch": true,
  "autoDeploy": false,
  "trackedBranch": "release",
  "gitopsConnectionId": null
}`

type recordedRequest struct {
	method, path, contentType string
	body                      []byte
}

func recordingClient(t *testing.T, status int, responseBody string) (*api.Client, *recordedRequest) {
	rec := &recordedRequest{}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec.method, rec.path = r.Method, r.URL.Path
		rec.contentType = r.Header.Get("Content-Type")
		rec.body, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = w.Write([]byte(responseBody))
	})
	return mocks.Client(t, "nullstone", handler), rec
}

func TestWorkspaceAutomations_GetDefaults(t *testing.T) {
	tests := []struct {
		name    string
		status  int
		body    string
		wantNil bool
		wantErr bool
	}{
		{name: "200 decodes the defaults", status: http.StatusOK, body: blockAutomationDefaultsBody},
		{name: "404 returns nil", status: http.StatusNotFound, body: `{"message": "block not found"}`, wantNil: true},
		{name: "500 returns an api error", status: http.StatusInternalServerError, body: `{"message": "boom"}`, wantErr: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, rec := recordingClient(t, test.status, test.body)

			got, err := client.WorkspaceAutomations().GetDefaults(context.Background(), 42, 7)

			assert.Equal(t, http.MethodGet, rec.method)
			assert.Equal(t, "/orgs/nullstone/stacks/42/blocks/7/automation", rec.path)
			if test.wantErr {
				var ae response.ApiError
				require.ErrorAs(t, err, &ae)
				assert.Nil(t, got)
				return
			}
			require.NoError(t, err)
			if test.wantNil {
				assert.Nil(t, got)
				return
			}
			require.NotNil(t, got)
			gitopsId := int64(3)
			assert.Equal(t, types.BlockAutomationDefaults{
				OrgName: "nullstone", StackId: 42, BlockId: 7,
				IsEnabled: true, AutoLaunch: false, AutoDeploy: true, DefaultBranch: "main",
				Repo: "nullstone/widgets", RepoUrl: "https://github.com/nullstone/widgets", GitopsConnectionId: &gitopsId,
			}, *got)
		})
	}
}

func TestWorkspaceAutomations_UpdateDefaults(t *testing.T) {
	client, rec := recordingClient(t, http.StatusOK, blockAutomationDefaultsBody)
	input := api.UpdateBlockAutomationDefaultsInput{IsEnabled: true, AutoLaunch: false, AutoDeploy: true, DefaultBranch: "main"}

	got, err := client.WorkspaceAutomations().UpdateDefaults(context.Background(), 42, 7, input)

	require.NoError(t, err)
	assert.Equal(t, http.MethodPut, rec.method)
	assert.Equal(t, "/orgs/nullstone/stacks/42/blocks/7/automation", rec.path)
	assert.Equal(t, "application/json", rec.contentType)
	assert.JSONEq(t, `{"isEnabled": true, "autoLaunch": false, "autoDeploy": true, "defaultBranch": "main"}`, string(rec.body))
	require.NotNil(t, got)
	assert.Equal(t, "main", got.DefaultBranch)
}

func TestWorkspaceAutomations_Get(t *testing.T) {
	tests := []struct {
		name    string
		status  int
		body    string
		wantNil bool
	}{
		{name: "200 decodes the workspace automation", status: http.StatusOK, body: workspaceAutomationBody},
		{name: "404 returns nil", status: http.StatusNotFound, body: `{"message": "env not found"}`, wantNil: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, rec := recordingClient(t, test.status, test.body)

			got, err := client.WorkspaceAutomations().Get(context.Background(), 42, 7, 9)

			require.NoError(t, err)
			assert.Equal(t, http.MethodGet, rec.method)
			assert.Equal(t, "/orgs/nullstone/stacks/42/blocks/7/envs/9/automation", rec.path)
			if test.wantNil {
				assert.Nil(t, got)
				return
			}
			require.NotNil(t, got)
			assert.Equal(t, types.WorkspaceAutomation{
				OrgName: "nullstone", StackId: 42, BlockId: 7, EnvId: 9,
				IsEnabled: true, AutoLaunch: true, AutoDeploy: false, TrackedBranch: "release",
			}, *got)
		})
	}
}

func TestWorkspaceAutomations_Update(t *testing.T) {
	client, rec := recordingClient(t, http.StatusOK, workspaceAutomationBody)
	input := api.UpdateWorkspaceAutomationInput{IsEnabled: true, AutoLaunch: true, AutoDeploy: false, TrackedBranch: "release"}

	got, err := client.WorkspaceAutomations().Update(context.Background(), 42, 7, 9, input)

	require.NoError(t, err)
	assert.Equal(t, http.MethodPut, rec.method)
	assert.Equal(t, "/orgs/nullstone/stacks/42/blocks/7/envs/9/automation", rec.path)
	// blockId is omitted when zero so the single-workspace body carries only the flags
	assert.JSONEq(t, `{"isEnabled": true, "autoLaunch": true, "autoDeploy": false, "trackedBranch": "release"}`, string(rec.body))
	require.NotNil(t, got)
	assert.Equal(t, "release", got.TrackedBranch)
}

func TestWorkspaceAutomations_ListByEnv(t *testing.T) {
	client, rec := recordingClient(t, http.StatusOK, `[`+workspaceAutomationBody+`]`)

	got, err := client.WorkspaceAutomations().ListByEnv(context.Background(), 42, 9)

	require.NoError(t, err)
	assert.Equal(t, http.MethodGet, rec.method)
	assert.Equal(t, "/orgs/nullstone/stacks/42/envs/9/automations", rec.path)
	require.Len(t, got, 1)
	assert.Equal(t, int64(7), got[0].BlockId)
}

func TestWorkspaceAutomations_ReplaceByEnv(t *testing.T) {
	client, rec := recordingClient(t, http.StatusOK, `[`+workspaceAutomationBody+`]`)
	input := []api.UpdateWorkspaceAutomationInput{
		{BlockId: 7, IsEnabled: true, AutoLaunch: true, AutoDeploy: false, TrackedBranch: "release"},
		{BlockId: 8, IsEnabled: false, AutoLaunch: true, AutoDeploy: false},
	}

	got, err := client.WorkspaceAutomations().ReplaceByEnv(context.Background(), 42, 9, input)

	require.NoError(t, err)
	assert.Equal(t, http.MethodPut, rec.method)
	assert.Equal(t, "/orgs/nullstone/stacks/42/envs/9/automations", rec.path)
	assert.JSONEq(t, `[
  {"blockId": 7, "isEnabled": true, "autoLaunch": true, "autoDeploy": false, "trackedBranch": "release"},
  {"blockId": 8, "isEnabled": false, "autoLaunch": true, "autoDeploy": false, "trackedBranch": ""}
]`, string(rec.body))
	require.Len(t, got, 1)
	assert.Equal(t, int64(9), got[0].EnvId)
}
