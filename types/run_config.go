package types

import (
	"github.com/google/uuid"
)

type RunConfig struct {
	WorkspaceUid  uuid.UUID         `json:"workspaceUid"`
	Source        string            `json:"source"`
	SourceVersion string            `json:"sourceVersion"`
	Variables     Variables         `json:"variables"`
	Connections   Connections       `json:"connections"`
	Capabilities  CapabilityConfigs `json:"capabilities"`
	Providers     Providers         `json:"providers"`
	Targets       RunTargets        `json:"targets"`
}
