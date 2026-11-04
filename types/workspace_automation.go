package types

// BlockAutomationDefaults is the block-level default for workspace automation.
// Every workspace of the block (any env type) seeds its WorkspaceAutomation from this row
// the first time it is read, so editing a default changes only workspaces created later.
//
// AutoDeploy and DefaultBranch only apply to Application blocks. Repo and RepoUrl mirror the
// block's repository as known when the defaults were last written; they are informational and
// are not writable through UpdateDefaults (the block owns its repository).
type BlockAutomationDefaults struct {
	OrgName string `json:"orgName"`
	StackId int64  `json:"stackId"`
	BlockId int64  `json:"blockId"`

	// IsEnabled means "this block participates in an env": it is provisioned on env launch and
	// IaC sync, and it is selected by default when a preview env is created.
	IsEnabled bool `json:"isEnabled"`
	// AutoLaunch means the env launches automatically after it is created (preview envs today).
	AutoLaunch bool `json:"autoLaunch"`
	// AutoDeploy means a push to the workspace's tracked branch (or PR) builds and deploys the app.
	AutoDeploy bool `json:"autoDeploy"`
	// DefaultBranch seeds TrackedBranch for new pipeline-env workspaces of an app.
	DefaultBranch string `json:"defaultBranch"`

	Repo               string `json:"repo"`
	RepoUrl            string `json:"repoUrl"`
	GitopsConnectionId *int64 `json:"gitopsConnectionId"`
}

// WorkspaceAutomation is the per-workspace (block + env) automation record.
// It applies to every block type and every env type; a row exists for a workspace once it has
// been read or written, seeded from the block's BlockAutomationDefaults.
//
// TrackedBranch is set only for Application workspaces in pipeline envs. Preview env workspaces
// leave it empty because the preview app record owns the pull request or branch being tracked.
type WorkspaceAutomation struct {
	OrgName string `json:"orgName"`
	StackId int64  `json:"stackId"`
	BlockId int64  `json:"blockId"`
	EnvId   int64  `json:"envId"`

	IsEnabled     bool   `json:"isEnabled"`
	AutoLaunch    bool   `json:"autoLaunch"`
	AutoDeploy    bool   `json:"autoDeploy"`
	TrackedBranch string `json:"trackedBranch"`

	GitopsConnectionId *int64 `json:"gitopsConnectionId"`
}
