package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/nullstone-io/go-api-client.v0/response"
	"gopkg.in/nullstone-io/go-api-client.v0/types"
)

// UpdateTemplateMetadataInput is a partial update of a block template's
// governance metadata. Every field is a pointer so a caller can update a single
// field without clearing the others.
type UpdateTemplateMetadataInput struct {
	// DataClassification updates the sensitivity level: nil leaves it untouched,
	// an empty string clears it (unclassified), any other value sets it.
	DataClassification *string `json:"dataClassification,omitempty"`
}

// ApplyTo merges the provided fields onto existing metadata, leaving untouched
// any field whose pointer is nil.
func (i UpdateTemplateMetadataInput) ApplyTo(existing types.WorkspaceMetadata) types.WorkspaceMetadata {
	if i.DataClassification != nil {
		existing.DataClassification = types.ClassificationLevel(*i.DataClassification)
	}
	return existing
}

// WorkspaceMetadata is the API resource for a block's governance metadata. It
// mirrors WorkspaceModule, but targets the block-level (env-invariant) template
// metadata rather than a per-env module.
type WorkspaceMetadata struct {
	Client *Client
}

func (wm WorkspaceMetadata) basePath(stackId, blockId, envId int64) string {
	return fmt.Sprintf("/orgs/%s/stacks/%d/blocks/%d/envs/%d/metadata", wm.Client.Config.OrgName, stackId, blockId, envId)
}

// Update - PUT /orgs/:orgName/stacks/:stackId/blocks/:blockId/metadata
// Each field of the input is optional; only the provided fields are updated.
func (wm WorkspaceMetadata) Update(ctx context.Context, stackId, blockId, envId int64, input UpdateTemplateMetadataInput) (*types.WorkspaceTemplate, error) {
	raw, _ := json.Marshal(input)
	res, err := wm.Client.Do(ctx, http.MethodPut, wm.basePath(stackId, blockId, envId), nil, nil, json.RawMessage(raw))
	if err != nil {
		return nil, err
	}
	return response.ReadJsonPtr[types.WorkspaceTemplate](res)
}
