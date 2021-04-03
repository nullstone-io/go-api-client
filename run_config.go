package api

import "github.com/google/uuid"

type RunConfig struct {
	WorkspaceUid  uuid.UUID   `json:"workspaceUid"`
	Source        string      `json:"source"`
	SourceVersion string      `json:"sourceVersion"`
	Variables     Variables   `json:"variables"`
	Connections   Connections `json:"connections"`
	Providers     Providers   `json:"providers"`
}
