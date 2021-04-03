package api

import (
	"github.com/google/uuid"
)

type MultiRun struct {
	UidCreatedModel
	TargetWorkspaceUid uuid.UUID `json:"targetWorkspaceUid"`
	IsDestroy          bool      `json:"isDestroy"`

	TargetWorkspace *Workspace `json:"targetWorkspace,omitempty"`
	Runs            []*Run     `json:"runs,omitempty"`
}
