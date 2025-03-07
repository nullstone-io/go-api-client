package types

import "time"

type WorkspaceChangeset struct {
	LatestUpdateAt *time.Time `json:"latestUpdateAt"`

	// Unapplied are workspace changes that have not been applied to a DesiredConfig
	Unapplied []WorkspaceChange `json:"unapplied"`
	// Outstanding are workspace changes that were applied to a DesiredConfig, but do not have a successful Run
	Outstanding []WorkspaceChange `json:"outstanding"`
	// InProgress are workspace changes that were applied to a DesiredConfig, but do not have a finished Run
	InProgress []WorkspaceChange `json:"inProgress"`
}
