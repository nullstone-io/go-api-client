package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

type IacSyncs struct {
	Client *Client
}

// TriggerIacSyncPayload is the body of POST /orgs/{orgName}/stacks/{stackId}/envs/{envId}/iac_syncs.
//
// AutoPlan queues an infra-update Run on workspaces with detected IaC changes (skipping
// not-provisioned workspaces). AutoApply implies AutoPlan and auto-approves the Runs.
//
// CommitInfo.Repository (owner and name) is required: it is the repository the sync runs
// for, which block and event ownership are checked against, and the server rejects the
// request (400) without it. The CLI reads it from the git origin remote or --repo. The rest
// of CommitInfo is advisory metadata read from the local git repo and may be empty.
//
// YamlConfigFiles, when non-empty, switches the sync to detached mode (NUL-49): the
// supplied IaC file contents are used directly and the workflow skips its GitHub fetch.
// Map keys are repo-relative paths (e.g. ".nullstone/dev.yml"); values are the file
// contents. The server validates keys (no `..`, no absolute paths) and total payload
// size before accepting the request.
//
// AutomationTool names the CI tool running the sync (AutomationToolGithubActions, ...), the
// same vocabulary as DeployCreatePayload.AutomationTool; empty means a person ran it. The
// server records an automation-tool sync with trigger event "automation" and the tool as
// the trigger source, and a sync with no tool as a manual ("user") sync. An environment
// whose governance disables manual IaC sync rejects the latter with 403 and error type
// ErrorTypeManualIacSyncDisabled (see types.EnvironmentGovernance).
type TriggerIacSyncPayload struct {
	AutoPlan        bool              `json:"autoPlan"`
	AutoApply       bool              `json:"autoApply"`
	CommitInfo      types.CommitInfo  `json:"commitInfo"`
	YamlConfigFiles map[string]string `json:"yamlConfigFiles,omitempty"`
	AutomationTool  string            `json:"automationTool,omitempty"`
}

// ErrorTypeManualIacSyncDisabled is the `type` of the 403 error returned when a manual
// (non-CI) IaC sync targets an environment with EnvironmentGovernance.DisableManualIacSync
// set. Clients match on it to print the server's message as the whole explanation.
const ErrorTypeManualIacSyncDisabled = "problems/manual-iac-sync-disabled"

func (s IacSyncs) basePath(stackId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/envs/%d/iac_syncs", s.Client.Config.OrgName, stackId, envId)
}

// Trigger - POST /orgs/:orgName/stacks/:stackId/envs/:envId/iac_syncs
func (s IacSyncs) Trigger(ctx context.Context, stackId, envId int64, payload TriggerIacSyncPayload) (*types.IntentWorkflow, error) {
	rawPayload, _ := json.Marshal(payload)
	res, err := s.Client.Do(ctx, http.MethodPost, s.basePath(stackId, envId), nil, nil, json.RawMessage(rawPayload))
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.IntentWorkflow](res)
}
