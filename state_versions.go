package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
)

type StateVersions struct {
	Client *Client
}

// StateMetadataInput mirrors arcana's models.StateVersionMetadata. Empty
// fields are skipped on the server side (partial update) — so only the
// fields the caller has resolved need to be populated.
//
// Actor here is the *run initiator* (the human or external trigger that
// kicked off the workspace workflow). The JWT identity making the PATCH
// call (typically the org runner) is not what we want to credit on the
// rail, so the caller must send the initiator explicitly.
type StateMetadataInput struct {
	Actor        string `json:"actor,omitempty"`
	ActorPicture string `json:"actorPicture,omitempty"`
	SourceKind   string `json:"sourceKind,omitempty"`
	TriggerEvent string `json:"triggerEvent,omitempty"`
	GitSha       string `json:"gitSha,omitempty"`
	GitBranch    string `json:"gitBranch,omitempty"`
	Summary      string `json:"summary,omitempty"`
}

// IsEmpty reports whether every field is unset — convenient gate for callers
// who want to skip the network round-trip when there's nothing to attach.
func (m StateMetadataInput) IsEmpty() bool {
	return m == StateMetadataInput{}
}

func (s StateVersions) runMetadataPath(stackId int64, workspaceUid, runUid uuid.UUID) string {
	return fmt.Sprintf(
		"orgs/%s/stacks/%d/workspaces/%s/runs/%s/state-metadata",
		s.Client.Config.OrgName, stackId, workspaceUid, runUid,
	)
}

// PatchRunMetadata - PATCH /orgs/:orgName/stacks/:stackId/workspaces/:workspaceUid/runs/:runUid/state-metadata
//
// One nullfire run may produce N state versions (split/partial applies). The
// server-side handler fans the partial update across every row keyed by the
// run UID. A 404 means the state server hasn't observed the state push yet — callers
// who want to retry should use response.IsNotFoundError to detect that case.
func (s StateVersions) PatchRunMetadata(ctx context.Context, stackId int64, workspaceUid, runUid uuid.UUID, input StateMetadataInput) error {
	raw, _ := json.Marshal(input)
	res, err := s.Client.Do(ctx, http.MethodPatch, s.runMetadataPath(stackId, workspaceUid, runUid), nil, nil, json.RawMessage(raw))
	if err != nil {
		return err
	}
	return response.Verify(res)
}
