package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

// PlatformDataOverlayInput is the body of PUT .../platform-data/:kind/overlay
//
// Actor is the deploy initiator to credit for the overlay; the server only
// trusts it when the caller is the org runner (same rule as StateMetadataInput).
type PlatformDataOverlayInput struct {
	DeployId int64                        `json:"deploy_id"`
	Actor    string                       `json:"actor"`
	Data     types.PlatformDataEnvOverlay `json:"data"`
}

type WorkspacePlatformData struct {
	Client *Client
}

func (w WorkspacePlatformData) basePath(stackId int64, workspaceUid uuid.UUID, kind string) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/workspaces/%s/platform-data/%s", w.Client.Config.OrgName, stackId, workspaceUid, kind)
}

func (w WorkspacePlatformData) overlayPath(stackId int64, workspaceUid uuid.UUID, kind string) string {
	return fmt.Sprintf("%s/overlay", w.basePath(stackId, workspaceUid, kind))
}

// Get - GET /orgs/:orgName/stacks/:stackId/workspaces/:workspaceUid/platform-data/:kind
// Returns nil, nil when the workspace has no state version or no record for the kind (404).
func (w WorkspacePlatformData) Get(ctx context.Context, stackId int64, workspaceUid uuid.UUID, kind string) (*types.PlatformData, error) {
	res, err := w.Client.Do(ctx, http.MethodGet, w.basePath(stackId, workspaceUid, kind), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	// response.ReadJsonPtr already swallows a 404 into (nil, nil); the explicit
	// check keeps that contract visible here since callers rely on nil meaning
	// "no record for this kind yet". It checks the status code directly so the
	// body is only read once (response.Verify consumes it on error).
	if res.StatusCode == http.StatusNotFound {
		res.Body.Close()
		return nil, nil
	}
	return response.ReadJsonPtr[types.PlatformData](res)
}

// PutOverlay - PUT /orgs/:orgName/stacks/:stackId/workspaces/:workspaceUid/platform-data/:kind/overlay
// Merges a deploy overlay into the platform data record for the kind.
// Requires org runner permissions.
func (w WorkspacePlatformData) PutOverlay(ctx context.Context, stackId int64, workspaceUid uuid.UUID, kind string, input PlatformDataOverlayInput) error {
	raw, err := json.Marshal(input)
	if err != nil {
		return fmt.Errorf("error marshaling platform data overlay: %w", err)
	}
	res, err := w.Client.Do(ctx, http.MethodPut, w.overlayPath(stackId, workspaceUid, kind), nil, nil, json.RawMessage(raw))
	if err != nil {
		return err
	}
	return response.Verify(res)
}
