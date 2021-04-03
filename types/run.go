package types

import (
	"github.com/google/uuid"
)

const (
	RunPhasePlan  = "plan"
	RunPhaseApply = "apply"
)

const (
	RunStatusQueued        = "queued"
	RunStatusResolving     = "resolving"
	RunStatusInitializing  = "initializing"
	RunStatusAwaiting      = "awaiting-dependencies"
	RunStatusRunning       = "running"
	RunStatusCompleted     = "completed"
	RunStatusNeedsApproval = "needs-approval"
	RunStatusFailed        = "failed"
	RunStatusCancelled     = "cancelled"
	RunStatusDisapproved   = "disapproved"
)

type Run struct {
	UidCreatedModel
	WorkspaceUid  uuid.UUID `json:"workspaceUid"`
	IsDestroy     bool      `json:"isDestroy"`
	IsApproved    *bool     `json:"isApproved"`
	Phase         string    `json:"phase"`
	Status        string    `json:"status"`
	StatusMessage string    `json:"statusMessage"`
	MultiRunUid   uuid.UUID `json:"multiRunUid"`

	MultiRun *MultiRun  `json:"multiRun,omitempty"`
	Config   *RunConfig `json:"config,omitempty"`
	Apply    *RunApply  `json:"apply,omitempty"`
}
