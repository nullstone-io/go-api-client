package api

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

type PreviewApps struct {
	Client *Client
}

func (p PreviewApps) basePath(stackId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/envs/%d/preview_apps", p.Client.Config.OrgName, stackId, envId)
}

func (p PreviewApps) stackPath(stackId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/preview_apps", p.Client.Config.OrgName, stackId)
}

func (p PreviewApps) stackByNamePath(stackName string) string {
	return fmt.Sprintf("/orgs/%s/stacks_by_name/%s/preview_apps", p.Client.Config.OrgName, stackName)
}

// List - GET /orgs/{orgName}/stacks/{stackId}/envs/{envId}/preview_apps
func (p PreviewApps) List(ctx context.Context, stackId, envId int64) ([]types.PreviewApp, error) {
	res, err := p.Client.Do(ctx, http.MethodGet, p.basePath(stackId, envId), nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return response.ReadJsonVal[[]types.PreviewApp](res)
}

type FindPreviewAppsInput struct {
	// Repo matches against either the repo full name (e.g. "acme/widgets") or the full repo URL
	Repo string
	// PullRequest matches against either the pull request number (e.g. GitHub PR #123) or the
	// Nullstone pull request id. When nil, preview apps are not filtered by pull request.
	PullRequest *int64
}

func (input FindPreviewAppsInput) query() url.Values {
	q := url.Values{}
	if input.Repo != "" {
		q.Set("repo", input.Repo)
	}
	if input.PullRequest != nil {
		q.Set("pull_request", strconv.FormatInt(*input.PullRequest, 10))
	}
	return q
}

// Find - GET /orgs/{orgName}/stacks/{stackId}/preview_apps
// Searches preview apps across every environment in the stack.
// This is used to locate a preview environment when only the repo and pull request are known.
func (p PreviewApps) Find(ctx context.Context, stackId int64, input FindPreviewAppsInput) ([]types.PreviewApp, error) {
	res, err := p.Client.Do(ctx, http.MethodGet, p.stackPath(stackId), input.query(), nil, nil)
	if err != nil {
		return nil, err
	}

	return response.ReadJsonVal[[]types.PreviewApp](res)
}

// FindByStackName - GET /orgs/{orgName}/stacks_by_name/{stackName}/preview_apps
// Same as Find for callers that know the stack by name instead of id (e.g. CI reacting to a PR trigger).
func (p PreviewApps) FindByStackName(ctx context.Context, stackName string, input FindPreviewAppsInput) ([]types.PreviewApp, error) {
	res, err := p.Client.Do(ctx, http.MethodGet, p.stackByNamePath(stackName), input.query(), nil, nil)
	if err != nil {
		return nil, err
	}

	return response.ReadJsonVal[[]types.PreviewApp](res)
}
