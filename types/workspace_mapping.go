package types

import (
	"time"

	"github.com/google/uuid"
)

// WorkspaceMapping ties the workspace that stores Terraform state to the Nullstone workspace
// (stack/block/env) it stores state for. Platform data is scoped to a stack through it.
// A mapping is written once and never modified.
type WorkspaceMapping struct {
	// ArcanaWorkspaceUid is the uid of the workspace in the state store
	ArcanaWorkspaceUid uuid.UUID `json:"arcanaWorkspaceUid"`
	// NullfireWorkspaceUid is the uid of the Nullstone workspace (types.Workspace.Uid)
	NullfireWorkspaceUid uuid.UUID `json:"nullfireWorkspaceUid"`
	OrgName              string    `json:"orgName"`
	StackId              int64     `json:"stackId"`
	BlockId              int64     `json:"blockId"`
	EnvId                int64     `json:"envId"`
	CreatedAt            time.Time `json:"createdAt"`
}

// WorkspaceMappingInput is the body of PUT .../workspaces/:workspaceUid/mapping
// The stack id comes from the route.
type WorkspaceMappingInput struct {
	BlockId      int64     `json:"blockId"`
	EnvId        int64     `json:"envId"`
	WorkspaceUid uuid.UUID `json:"workspaceUid"`
}
