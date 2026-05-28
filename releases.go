package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

type Releases struct {
	Client *Client
}

// ReleaseCreatePayload is the payload for creating a release.
//
// A release runs infra-update (apply) and/or deploy through a single server-side workflow.
// The version/commit fields mirror DeployCreatePayload; the engine resolves the app version
// and decides the path (apply+deploy when infra changed, deploy-only otherwise).
type ReleaseCreatePayload struct {
	FromSource     bool   `json:"fromSource"`
	CommitSha      string `json:"commitSha"`
	Version        string `json:"version"`
	Reference      string `json:"reference"`
	AutomationTool string `json:"automationTool"`

	// EnvVars are additional environment variables to set on the app's infra resources for this release.
	EnvVars map[string]string `json:"envVars,omitempty"`

	// IsApproved approves the infra-update when the release runs an apply. Honored only for stack architects.
	IsApproved *bool `json:"isApproved,omitempty"`
}

func (r Releases) basePath(stackId, appId, envId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/apps/%d/envs/%d/releases", r.Client.Config.OrgName, stackId, appId, envId)
}

// Create starts an app-release intent workflow and returns it.
// The endpoint always starts a workflow; a "nothing to release" outcome surfaces as
// the returned IntentWorkflow terminating with status types.IntentWorkflowStatusNoOp.
func (r Releases) Create(ctx context.Context, stackId, appId, envId int64, payload ReleaseCreatePayload) (*types.IntentWorkflow, error) {
	rawPayload, _ := json.Marshal(payload)
	res, err := r.Client.Do(ctx, http.MethodPost, r.basePath(stackId, appId, envId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.IntentWorkflow](res)
}
