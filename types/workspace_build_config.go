package types

import (
	"time"

	"github.com/google/uuid"
)

// WorkspaceBuildConfig is the build configuration for a single app workspace (block + env).
// It is stored beside the workspace config, not inside it: changes take effect on the next
// build with no apply or infra run. A row is materialized on first read from the block
// template's Build (WorkspaceTemplateConfig.Build) or, failing that, DefaultBuildConfig(module).
type WorkspaceBuildConfig struct {
	WorkspaceUid uuid.UUID `json:"workspaceUid"`
	// Version increments on every write so callers can detect changes between reads.
	Version   int64       `json:"version"`
	UpdatedAt time.Time   `json:"updatedAt"`
	UpdatedBy string      `json:"updatedBy"`
	Config    BuildConfig `json:"config"`
}
