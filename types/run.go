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
	WorkspaceUid uuid.UUID `json:"workspaceUid"`

	// IsDestroy determines whether to run a destroy plan instead of an apply plan
	IsDestroy bool `json:"isDestroy"`

	// IsApproved determines whether the user has approved this run
	// A value of nil indicates that a decision was not made yet
	IsApproved *bool `json:"isApproved"`

	// IsCleanup is flagged when this run cleans up resources that could not be cleaned up in a previous run
	// Typically, this is flagged in conjunction with RunConfig.Targets
	IsCleanup bool `json:"isCleanup"`

	// Phase determines the current phase of the run (plan|apply)
	Phase string `json:"phase"`
	// Status determines the current status of the run; whether it's running, needs approval, or finished
	Status string `json:"status"`
	// StatusMessage contains error messages when the status is in a failure state
	StatusMessage string `json:"statusMessage"`

	// MultiRunUid is set to a non-zero uuid when this run is contained within a multi-run
	// A multi-run spans multiple workspaces in an effort to gracefully apply/destroy multiple workspaces
	MultiRunUid uuid.UUID `json:"multiRunUid"`

	MultiRun *MultiRun  `json:"multiRun,omitempty"`
	Config   *RunConfig `json:"config,omitempty"`
	Apply    *RunApply  `json:"apply,omitempty"`
}
