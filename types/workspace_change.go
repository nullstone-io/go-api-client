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
	Id           int64        `json:"id"`
	WorkspaceUid uuid.UUID    `json:"workspaceUid"`
	ChangeType   ChangeType   `json:"changeType"`
	Identifier   string       `json:"identifier"`
	Action       ChangeAction `json:"action"`
	Version      int64        `json:"version"`
	CreatedAt    time.Time    `json:"createdAt"`
	CreatedBy    string       `json:"createdBy"`
	Current      any          `json:"current,omitempty"`
	Desired      any          `json:"desired"`
}
