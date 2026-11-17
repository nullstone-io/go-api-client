package types

// BuildConfigChange is one setting of a build config that an IaC sync changed.
// A build config takes effect on the next build, so these are a record of what the sync wrote;
// they are never pending and never cause an infra run (unlike WorkspaceChange).
type BuildConfigChange struct {
	// Key is the setting's path under `build:` in an IaC file (e.g. builder, nullstone.dockerfile)
	Key string `json:"key"`
	// Current is the value before the sync
	Current any `json:"current"`
	// Desired is the value the sync wrote
	Desired any `json:"desired"`
}
