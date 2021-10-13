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
	OrgName  string    `json:"orgName"`
	StackId  int64     `json:"stackId"`
	BlockId  int64     `json:"blockId"`
	EnvId    int64     `json:"envId"`
	Status   string    `json:"status"`
	StatusAt time.Time `json:"statusAt"`

	NumQueuedRuns *int `json:"numQueuedRuns,omitempty"`

	ActiveRun         *Run `json:"activeRun,omitempty"`
	LastFinishedRun   *Run `json:"lastFinishedRun,omitempty"`
	LastSuccessfulRun *Run `json:"lastSuccessfulRun,omitempty"`

	// Deprecated
	StackName string `json:"stackName"`
	// Deprecated
	BlockName string `json:"blockName"`
	// Deprecated
	EnvName string `json:"envName"`
}
