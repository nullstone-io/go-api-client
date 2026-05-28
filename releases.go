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
// A release runs infra-update (apply) and/or deploy for one or more apps through a single
// server-side workflow. The engine builds a dependency tree rooted at the requested apps,
// applies any infra changes, and deploys the resolved version for each requested app.
type ReleaseCreatePayload struct {
	// Apps are the apps to release. At least one is required.
	Apps []ReleaseApp `json:"apps"`

	// AutomationTool refers to the tool that triggered this release (release-global; empty = manual).
	AutomationTool string `json:"automationTool"`

	// IsApproved approves the infra-update when the release runs an apply. Honored only for stack architects.
	IsApproved *bool `json:"isApproved,omitempty"`
}

// ReleaseApp identifies a single app to release and its per-app deploy inputs.
type ReleaseApp struct {
	AppId      int64  `json:"appId"`
	FromSource bool   `json:"fromSource"`
	CommitSha  string `json:"commitSha"`
	Version    string `json:"version"`
	Reference  string `json:"reference"`
	// EnvVars are additional environment variables to set on the app's infra resources for this release.
	EnvVars map[string]string `json:"envVars,omitempty"`
}

func (r Releases) basePath(stackId, envId int64) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/envs/%d/releases", r.Client.Config.OrgName, stackId, envId)
}

// Create starts an app-release intent workflow and returns it.
func (r Releases) Create(ctx context.Context, stackId, envId int64, payload ReleaseCreatePayload) (*types.IntentWorkflow, error) {
	rawPayload, _ := json.Marshal(payload)
	res, err := r.Client.Do(ctx, http.MethodPost, r.basePath(stackId, envId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.IntentWorkflow](res)
}
