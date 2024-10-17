package types

import (
	"github.com/google/uuid"
	"reflect"
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

func (wc *WorkspaceChange) Merge(latest *WorkspaceChange) *WorkspaceChange {
	// 9 permutations based on action: (3 x 3)
	// wc add		latest add
	// wc add		latest update
	// wc add		latest delete
	// wc update	latest add
	// wc update	latest update
	// wc update	latest delete
	// wc delete	latest add
	// wc delete	latest update
	// wc delete	latest delete

	switch wc.Action {
	case ChangeActionAdd:
		switch latest.Action {
		case ChangeActionAdd:
			// The latest add wins
			return latest
		case ChangeActionUpdate:
			// Use "latest", but with Add and set Current=nil
			// user+version pulled from "latest"
			result := *latest
			result.Action = ChangeActionAdd
			result.Current = nil
			return &result
		case ChangeActionDelete:
			// undo the original add
			return nil
		}
	case ChangeActionUpdate:
		switch latest.Action {
		case ChangeActionAdd:
			// It doesn't make sense to add something after an update
			// Ignore the add
			return wc
		case ChangeActionUpdate:
			// Applying an update over an update
			if reflect.DeepEqual(wc.Current, latest.Desired) {
				// The update is effectively ignored
				return nil
			}
			// Use Current from the first, Desired from the second
			// Use user+version from second
			result := *latest
			result.Current = wc.Current
			return &result
		case ChangeActionDelete:
			// The update is effectively ignored
			result := *latest
			result.Current = wc.Current
			return &result
		}
	case ChangeActionDelete:
		switch latest.Action {
		case ChangeActionAdd:
			if reflect.DeepEqual(wc.Current, latest.Desired) {
				// undo the original
				return nil
			}
			result := *latest
			result.Action = ChangeActionUpdate
			result.Current = wc.Current
			return &result
		case ChangeActionUpdate:
			if reflect.DeepEqual(wc.Current, latest.Desired) {
				// undo the original
				return nil
			}
			result := *latest
			result.Action = ChangeActionUpdate
			result.Current = wc.Current
			return &result
		case ChangeActionDelete:
			// Can't triple-stamp a double-stamp lloyd!
			// Ignore the second delete
			return wc
		}
	}
	return nil
}
