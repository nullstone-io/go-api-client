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
	// AutoDeploy means a push to the workspace's tracked branch (or PR) builds and deploys the app.
	AutoDeploy bool `json:"autoDeploy"`
	// DefaultBranch seeds TrackedBranch for new pipeline-env workspaces of an app, unless another
	// application of the same repository already follows something in that env.
	DefaultBranch string `json:"defaultBranch"`

	Repo               string `json:"repo"`
	RepoUrl            string `json:"repoUrl"`
	GitopsConnectionId *int64 `json:"gitopsConnectionId"`
}

// WorkspaceAutomation is the per-workspace (block + env) automation record.
// It applies to every block type and every env type; a row exists for a workspace once it has
// been read or written, seeded from the block's BlockAutomationDefaults.
//
// TrackedBranch and PullRequestId are what an Application workspace follows: a branch or a pull
// request, never both. Every application of one repository follows the same one in an env, so they
// are read-only here and written per repository through WorkspaceAutomations.SetSource.
type WorkspaceAutomation struct {
	OrgName string `json:"orgName"`
	StackId int64  `json:"stackId"`
	BlockId int64  `json:"blockId"`
	EnvId   int64  `json:"envId"`

	IsEnabled  bool `json:"isEnabled"`
	AutoDeploy bool `json:"autoDeploy"`

	// TrackedBranch is the branch the workspace follows; empty when it follows a pull request or nothing.
	TrackedBranch string `json:"trackedBranch"`
	// PullRequestId is the pull request the workspace follows; nil when it follows a branch or nothing.
	PullRequestId *int64 `json:"pullRequestId"`
}
