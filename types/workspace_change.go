package types

import (
	"github.com/google/uuid"
	"time"
)

type ChangeType string

const (
	// ChangeTypeModuleVersion - there can only be one of these; no identifier needed
	ChangeTypeModuleVersion ChangeType = "module_version"
	// ChangeTypeVariable - use variable name (key) as identifier
	ChangeTypeVariable ChangeType = "variable"
	// ChangeTypeEnvVariable - use env variable key as identifier
	ChangeTypeEnvVariable ChangeType = "env_variable"
	// ChangeTypeCapability - use the capability id, it is unique within a workspace
	ChangeTypeCapability ChangeType = "capability"
	// ChangeTypeConnection - use the connection key as identifier
	ChangeTypeConnection ChangeType = "connection"
)

type ChangeAction string

const (
	ChangeActionAdd    ChangeAction = "add"
	ChangeActionUpdate ChangeAction = "update"
	ChangeActionDelete ChangeAction = "delete"
)

const (
	ChangeIdentifierModuleVersion = "module_version"
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
