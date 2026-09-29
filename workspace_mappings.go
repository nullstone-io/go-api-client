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

type WorkspaceMappings struct {
	Client *Client
}

func (w WorkspaceMappings) path(stackId int64, workspaceUid uuid.UUID) string {
	return fmt.Sprintf("orgs/%s/stacks/%d/workspaces/%s/mapping", w.Client.Config.OrgName, stackId, workspaceUid)
}

// Put - PUT /orgs/:orgName/stacks/:stackId/workspaces/:workspaceUid/mapping
// Registers which Nullstone workspace (stack/block/env) the state workspace stores state for.
// The mapping is inserted when absent; when the workspace already has a mapping, it is returned untouched.
//
// Errors are returned as is, including response.NotFoundError (404):
// the state workspace is created on first backend access, so it does not exist before the first `terraform init`.
func (w WorkspaceMappings) Put(ctx context.Context, stackId int64, workspaceUid uuid.UUID, input types.WorkspaceMappingInput) (*types.WorkspaceMapping, error) {
	raw, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("error marshaling workspace mapping: %w", err)
	}
	res, err := w.Client.Do(ctx, http.MethodPut, w.path(stackId, workspaceUid), nil, nil, json.RawMessage(raw))
	if err != nil {
		return nil, err
	}
	// response.ReadJsonPtr swallows a 404 into (nil, nil); verify first so a 404 reaches the caller as an error
	if err := response.Verify(res); err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.WorkspaceMapping](res)
}
