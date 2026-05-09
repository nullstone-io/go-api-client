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
type TriggerIacSyncPayload struct {
	AutoPlan   bool             `json:"autoPlan"`
	AutoApply  bool             `json:"autoApply"`
	CommitInfo types.CommitInfo `json:"commitInfo"`

	// TODO(detached-mode, NUL-49): add optional YamlConfigFiles map[string]string here so
	// callers can submit IaC contents directly instead of forcing the workflow to fetch
	// from GitHub. Mirror the field on the server. See iac-sync-detached-mode.md.
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
