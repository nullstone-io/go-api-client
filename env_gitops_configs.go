package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

// EnvGitopsConfigs is the API resource for what one env follows from each repository
// (types.GitopsConnectionConfig), seen from the env.
//
// A preview env is set up like any other env: its WorkspaceAutomations say which blocks take part, and these
// configs say what it follows. They replace preview apps, which are no longer stored; the PreviewApps read
// endpoints remain and answer from these.
type EnvGitopsConfigs struct {
	Client *Client
}

func (e EnvGitopsConfigs) envPath(stackId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/envs/%d/gitops_configs", e.Client.Config.OrgName, stackId, envId)
}

// SetEnvGitopsTriggerInput names what a preview env follows from one repository
type SetEnvGitopsTriggerInput struct {
	RepoUrl string `json:"repoUrl"`
	types.GitopsTrigger
}

// List - GET /orgs/:orgName/stacks/:stackId/envs/:envId/gitops_configs
// Returns one config per repository the env follows, for any env type, with PullRequest loaded.
func (e EnvGitopsConfigs) List(ctx context.Context, stackId, envId int64) ([]types.GitopsConnectionConfig, error) {
	res, err := e.Client.Do(ctx, http.MethodGet, e.envPath(stackId, envId), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonVal[[]types.GitopsConnectionConfig](res)
}

// Replace - PUT /orgs/:orgName/stacks/:stackId/envs/:envId/gitops_configs
// Sets what a preview env follows from each repository in triggers; repositories that are not listed are left
// untouched. The response is the env's full list, as List returns it. A repository whose trigger changed is
// synced right away when it has a commit to sync (a branch or a pull request).
// Only a preview env is configured here; a pipeline env's configs are written with its gitops connection.
func (e EnvGitopsConfigs) Replace(ctx context.Context, stackId, envId int64, triggers []SetEnvGitopsTriggerInput) ([]types.GitopsConnectionConfig, error) {
	rawPayload, _ := json.Marshal(triggers)
	res, err := e.Client.Do(ctx, http.MethodPut, e.envPath(stackId, envId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	return response.ReadJsonVal[[]types.GitopsConnectionConfig](res)
}
