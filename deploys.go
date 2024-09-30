package api

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	AutomationToolCircleCI       = "circleci"
	AutomationToolGithubActions  = "github-actions"
	AutomationToolGitlab         = "gitlab"
	AutomationToolBitbucket      = "bitbucket"
	AutomationToolJenkins        = "jenkins"
	AutomationToolTravis         = "travis"
	AutomationToolAzurePipelines = "azure-pipeline"
	AutomationToolAppveyor       = "appveyor"
	AutomationToolTeamCity       = "team-city"
	AutomationToolCodeship       = "codeship"
	AutomationToolSemaphore      = "semaphore"
)

type Deploys struct {
	Client *Client
}

// DeployCreatePayload is the payload for creating a deploy
//
//		If `FromSource` is specified, we deploy from the latest commit on the configured branch for this environment
//		Otherwise, version is required
//		commitSha and reference are optional and get populated on the deploy no matter whether we are deploying
//	   fromSource or by version
type DeployCreatePayload struct {
	FromSource     bool   `json:"fromSource"`
	CommitSha      string `json:"commitSha"`
	Version        string `json:"version"`
	Reference      string `json:"reference"`
	AutomationTool string `json:"automationTool"`
}

// DeployCreateResult contains the result of Deploys Create
// The result can be one of types.Deploy or types.IntentWorkflow
type DeployCreateResult struct {
	Deploy         *types.Deploy
	IntentWorkflow *types.IntentWorkflow
}

func (d Deploys) basePath(stackId, appId, envId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/apps/%d/envs/%d/deploys", d.Client.Config.OrgName, stackId, appId, envId)
}

func (d Deploys) path(stackId, appId, envId, deployId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/apps/%d/envs/%d/deploys/%d", d.Client.Config.OrgName, stackId, appId, envId, deployId)
}

func (d Deploys) Create(ctx context.Context, stackId, appId, envId int64, payload DeployCreatePayload) (*DeployCreateResult, error) {
	rawPayload, _ := json.Marshal(payload)
	res, err := d.Client.Do(ctx, http.MethodPost, d.basePath(stackId, appId, envId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	if err := response.Verify(res); err != nil {
		if response.IsNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	defer res.Body.Close()
	result := &DeployCreateResult{}
	if raw, err := io.ReadAll(res.Body); err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	} else {
		// Try to parse into IntentWorkflow; if it doesn't match, parse into Deploy
		if err := json.Unmarshal(raw, &result.IntentWorkflow); err != nil || result.IntentWorkflow.Intent == "" {
			result.IntentWorkflow = nil
			if err := json.Unmarshal(raw, &result.Deploy); err != nil {
				return result, fmt.Errorf("unknown response body: %w", err)
			}
		}
	}
	return result, nil
}

func (d Deploys) Get(ctx context.Context, stackId, appId, envId, deployId int64) (*types.Deploy, error) {
	res, err := d.Client.Do(ctx, http.MethodGet, d.path(stackId, appId, envId, deployId), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.Deploy](res)
}

func (d Deploys) GetLatest(ctx context.Context, stackId, appId, envId int64, since *time.Time) (*types.Deploy, error) {
	query := url.Values{}
	if since != nil {
		query["since"] = []string{since.Format(time.RFC3339)}
	}
	res, err := d.Client.Do(ctx, http.MethodGet, d.basePath(stackId, appId, envId)+"/latest", query, nil, nil)
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.Deploy](res)
}
