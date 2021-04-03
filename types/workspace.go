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
	OrgName   string    `json:"orgName"`
	StackName string    `json:"stackName"`
	BlockName string    `json:"blockName"`
	EnvName   string    `json:"envName"`
	Status    string    `json:"status"`
	StatusAt  time.Time `json:"statusAt"`

	ActiveRun         *Run              `json:"activeRun,omitempty" pg:"-"`
	LastSuccessfulRun *Run              `json:"lastSuccessfulRun,omitempty" pg:"-"`
	Dependencies      []WorkspaceTarget `json:"dependencies,omitempty" pg:"-"`
}
