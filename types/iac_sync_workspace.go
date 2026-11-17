package types

import "github.com/google/uuid"

// IacSyncWorkspace represents the progress of a single workspace update within an IacSync.
type IacSyncWorkspace struct {
	Uid                 uuid.UUID         `json:"uid"`
	IacSyncUid          uuid.UUID         `json:"iacSyncUid"`
	WorkspaceUid        uuid.UUID         `json:"workspaceUid"`
	PriorConfigVersion  int64             `json:"priorConfigVersion"`
	ConfigVersion       int64             `json:"configVersion"`
	WorkspaceWorkflowId *int64            `json:"workspaceWorkflowId"`
	IacSync             IacSync           `json:"iacSync"`
	WorkspaceChanges    []WorkspaceChange `json:"workspaceChanges"`
	// BuildChanges lists the build config settings this sync wrote for the workspace.
	// They are informational: a build config applies to the next build with no infra run.
	BuildChanges []BuildConfigChange `json:"buildChanges"`
}
