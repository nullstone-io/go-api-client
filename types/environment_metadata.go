package types

// EnvironmentMetadata is an extensible container for platform-defined
// descriptive metadata about an environment. Description is its first member;
// future metadata (owner, lifecycle policy, …) lands here too, without a
// migration — the whole struct is persisted as a single jsonb column.
//
// It is deliberately a *closed* set of fields, mirroring WorkspaceMetadata.
// Open-ended, user-defined keys belong in Environment.Tags instead.
type EnvironmentMetadata struct {
	// Description is free-form prose describing what the environment is for.
	// Empty = no description.
	Description string `json:"description,omitempty"`
}
