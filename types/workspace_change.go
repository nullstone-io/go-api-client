package types

import (
	"github.com/google/uuid"
	"time"
)

type ChangeType string

const (
	// ModuleVersionChange - there can only be one of these; no identifier needed
	ModuleVersionChange ChangeType = "module_version"
	// VariableChange - use variable name (key) as identifier
	VariableChange ChangeType = "variable"
	// EnvVariableChange - use env variable key as identifier
	EnvVariableChange ChangeType = "env_variable"
	// CapabilityChange - use the capability id, it is unique within a workspace
	CapabilityChange ChangeType = "capability"
)

type ChangeAction string

const (
	ChangeActionAdd    ChangeAction = "add"
	ChangeActionUpdate ChangeAction = "update"
	ChangeActionDelete ChangeAction = "delete"
)

type WorkspaceChange struct {
	tableName    struct{}     `pg:"workspace_changes"`
	Id           int64        `json:"id" pg:"id,pk"`
	WorkspaceUid uuid.UUID    `json:"workspaceUid" pg:"workspace_uid"`
	ChangeType   ChangeType   `json:"changeType" pg:"change_type"`
	Identifier   string       `json:"identifier" pg:"identifier"`
	Action       ChangeAction `json:"action" pg:"action"`
	Version      int64        `json:"version" pg:"version"`
	CreatedAt    time.Time    `json:"createdAt" pg:"created_at"`
	CreatedBy    string       `json:"createdBy" pg:"created_by"`
	Current      any          `json:"current,omitempty" pg:"-"`
	Desired      any          `json:"desired" pg:"desired"`
}
