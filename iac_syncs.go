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
// CommitInfo is best-effort metadata read from the user's local git repo; the server treats
// it as advisory and tolerates an empty value.
//
// YamlConfigFiles, when non-empty, switches the sync to detached mode (NUL-49): the
// supplied IaC file contents are used directly and the workflow skips its GitHub fetch.
// Map keys are repo-relative paths (e.g. ".nullstone/dev.yml"); values are the file
// contents. The server validates keys (no `..`, no absolute paths) and total payload
// size before accepting the request.
type TriggerIacSyncPayload struct {
	AutoPlan        bool              `json:"autoPlan"`
	AutoApply       bool              `json:"autoApply"`
	CommitInfo      types.CommitInfo  `json:"commitInfo"`
	YamlConfigFiles map[string]string `json:"yamlConfigFiles,omitempty"`
}

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
