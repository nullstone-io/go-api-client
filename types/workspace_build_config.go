package types

import "github.com/google/uuid"

// WorkspaceBuildConfig is the build configuration for a single app workspace (block + env).
// It is stored beside the workspace config, not inside it: a write takes effect on the next
// build with no apply or infra run, so there is no version history to expose (unlike
// WorkspaceConfig). The server materializes it on first read from the block template's
// Build (WorkspaceTemplateConfig.Build) or, failing that, DefaultBuildConfig(module).
//
// The addressing fields mirror WorkspaceTemplate so a consumer holding only the response
// can re-address the workspace (stack/block/env) without a second lookup.
type WorkspaceBuildConfig struct {
	OrgName      string      `json:"orgName"`
	StackId      int64       `json:"stackId"`
	BlockId      int64       `json:"blockId"`
	EnvId        int64       `json:"envId"`
	WorkspaceUid uuid.UUID   `json:"workspaceUid"`
	Config       BuildConfig `json:"config"`
}
