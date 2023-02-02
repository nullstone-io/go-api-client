package types

type CreateRunInput struct {
	// Create a run that destroys this workspace
	IsDestroy bool `json:"isDestroy"`

	// DestroyDependencies allows the user to identify which dependencies to destroy along with the block
	// IsDestroy must be enabled for this field to have an effect
	// `*` indicates attempt to destroy the workspace and its dependencies
	// ``  indicates attempt to destroy only the specified workspace
	// `<stack-id>/<block-id>/<env-id>,...` indicates a comma-delimited list of dependencies to destroy with the workspace
	DestroyDependencies string `json:"destroyDependencies"`

	IsApproved *bool `json:"isApproved"`
}
