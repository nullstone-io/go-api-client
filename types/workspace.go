package types

import (
	"time"
)

const (
	WorkspaceStatusNotProvisioned = "not-provisioned"
	WorkspaceStatusProvisioned    = "provisioned"
)

type Workspace struct {
	UidCreatedModel
	OrgName       string    `json:"orgName"`
	StackId       int64     `json:"stackId"`
	BlockId       int64     `json:"blockId"`
	EnvId         int64     `json:"envId"`
	Status        string    `json:"status"`
	StatusAt      time.Time `json:"statusAt"`
	ProvisionedAt time.Time `json:"provisionedAt"`

	NumQueuedRuns *int `json:"numQueuedRuns,omitempty"`

	// ActiveRun represents the currently-active run
	// This is not guaranteed to be the most recent run created
	// All "queued" runs are essentially waiting for completion of ActiveRun
	ActiveRun *Run `json:"activeRun,omitempty"`

	// LastFinishedRun represents the most recent run that finished (completed or failed)
	LastFinishedRun *Run `json:"lastFinishedRun,omitempty"`

	// LastSuccessfulRun represents the most recent run that completed successfully
	// This run *must* be in the "completed" status
	LastSuccessfulRun *Run `json:"lastSuccessfulRun,omitempty"`

	// Deprecated
	StackName string `json:"stackName"`
	// Deprecated
	BlockName string `json:"blockName"`
	// Deprecated
	EnvName string `json:"envName"`

	// HasPartialChanges is set to true when the most recent run did not fully apply changes
	// The engine should update this in three cases:
	// 1. A plan detects no changes
	// 2. The execution of TF apply succeeds
	// 3. The execution of TF apply fails
	HasPartialChanges bool `json:"hasPartialChanges"`

	// CurrentActivity identifies what is currently happening on the workspace
	// e.g., idle, provisioning, deprovisioning
	CurrentActivity string `json:"currentActivity"`
}
