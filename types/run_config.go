package types

import (
	"github.com/google/uuid"
)

type RunConfig struct {
	WorkspaceUid uuid.UUID `json:"workspaceUid"`
	// Targets represents a list of fully-qualified resource addresses
	// During this run, the engine will attempt to update *only* these targets
	//
	// Targets can be specified with suffix wildcards (e.g. module.cap_123.*)
	// These suffix wildcards *must* end with '.*'
	Targets RunTargets `json:"targets"`
	// PrevDependencies represents a list of workspace references that were dependencies of the previous run
	PrevDependencies Dependencies `json:"prevDependencies"`

	WorkspaceConfig `json:",inline"`
}
