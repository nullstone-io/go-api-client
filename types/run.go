package types

import (
	"github.com/google/uuid"
	"time"
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

func IsTerminalRunStatus(runStatus string) bool {
	switch runStatus {
	default:
		return false
	case RunStatusCompleted:
	case RunStatusFailed:
	case RunStatusCancelled:
	case RunStatusDisapproved:
	}
	return true
}

type Run struct {
	UidCreatedModel
	WorkspaceUid uuid.UUID `json:"workspaceUid"`
	CommitSha    string    `json:"commitSha"`

	// IsLaunch flags whether this run is provisioning a workspace for the first time or after a destroy
	IsLaunch bool `json:"isLaunch"`
	// IsDestroy determines whether to run a destroy plan instead of an apply plan
	IsDestroy bool `json:"isDestroy"`
	// DestroyDependencies is a list of dependencies that a user wants to destroy along with the primary workspace
	// `*` - All dependencies
	// ``  - No dependencies
	// `<stack-id>/<block-id>/<env-id>,...` - a list of comma-delimited workspaces that should be destroyed
	DestroyDependencies string `json:"destroyDependencies"`

	// IsApproved determines whether the user has approved this run
	// A value of nil indicates that a decision was not made yet
	IsApproved *bool `json:"isApproved"`
	// ApprovedBy is the person who made the approval
	// A value of `auto::<username>` means this run was auto-approved
	ApprovedBy string `json:"approvedBy"`
	// ApprovedAt is the time when the person approved the run
	// or when it was auto-approved
	ApprovedAt time.Time `json:"approvedAt"`

	// IsCleanup is flagged when this run cleans up resources that could not be cleaned up in a previous run
	// Typically, this is flagged in conjunction with RunConfig.Targets
	IsCleanup bool `json:"isCleanup"`

	// Phase determines the current phase of the run (plan|apply)
	Phase string `json:"phase"`
	// Status determines the current status of the run; whether it's running, needs approval, or finished
	Status string `json:"status"`
	// StatusMessage contains error messages when the status is in a failure state
	StatusMessage string `json:"statusMessage"`

	// IsPrimary indicates that this run is the primary run created by the user
	// The processor sets this to false when creating runs for dependencies
	// If this is set to false, this run should not be picked up by a worker because another worker created it and is managing it
	IsPrimary bool `json:"isPrimary"`
	// IsCandidate indicates that this run is not guaranteed to execute
	// When the run is picked up, the engine will recompute whether to execute the run
	// Candidate runs serve as a mechanism to reserve a spot in a workspace's queue
	// This flips to false if the engine computes that this run should execute
	IsCandidate bool `json:"isCandidate"`
	// MultiRunUid is set to a non-zero uuid when this run is contained within a multi-run
	// A multi-run spans multiple workspaces in an effort to gracefully apply/destroy multiple workspaces
	MultiRunUid uuid.UUID `json:"multiRunUid"`
	// RequiredBy is a list of workspaces that forced this run to be created
	// For instance, a cluster needs a network.
	//   A launch run would have a network run  RequiredBy: cluster
	//   A destroy run would have a cluster run RequiredBy: network
	RequiredBy []WorkspaceTarget `json:"requiredBy"`
	// AuditedRunUid represents the Run that this Run is auditing for provisioning
	// If this is null, this Run is a regular Run
	// If not, this Run is considered an AuditRun
	// AuditRuns are excluded from LastFinishedRun, ActiveRun, NextRun, and Pending run
	AuditedRunUid *uuid.UUID `json:"auditedRunUid"`
	// Condition provides a mechanism for conditionally executing a run
	// This allows the engine to evict runs that do not meet the condition when the run is ready to plan/apply
	// If empty, this run will always execute
	Condition string `json:"condition"`

	WorkspaceWorkflowId *int64 `json:"workspaceWorkflowId"`

	MultiRun *MultiRun  `json:"multiRun,omitempty"`
	Config   *RunConfig `json:"config,omitempty"`
}
