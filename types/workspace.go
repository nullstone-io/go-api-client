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

	Module            *Module           `json:"module,omitempty"`
	ActiveRun         *Run              `json:"activeRun,omitempty"`
	LastSuccessfulRun *Run              `json:"lastSuccessfulRun,omitempty"`
	Dependencies      []WorkspaceTarget `json:"dependencies,omitempty"`
}
