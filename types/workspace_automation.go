package types

// BlockAutomationDefaults is the block-level default for workspace automation.
// Every workspace of the block (any env type) seeds its WorkspaceAutomation from this row
// the first time it is read, so editing a default changes only workspaces created later.
//
// AutoDeploy only applies to Application blocks. Repo and RepoUrl mirror the
// block's repository as known when the defaults were last written; they are informational and
// are not writable through UpdateDefaults (the block owns its repository).
type BlockAutomationDefaults struct {
	OrgName string `json:"orgName"`
	StackId int64  `json:"stackId"`
	BlockId int64  `json:"blockId"`

	// IsEnabled means "this block participates in an env": it is provisioned on env launch and
	// IaC sync, and it is selected by default when a preview env is created.
	IsEnabled bool `json:"isEnabled"`
	// AutoDeploy means the app builds and deploys when its repository's trigger fires in an env.
	AutoDeploy bool `json:"autoDeploy"`

	Repo               string `json:"repo"`
	RepoUrl            string `json:"repoUrl"`
	GitopsConnectionId *int64 `json:"gitopsConnectionId"`
}

// WorkspaceAutomation is the per-workspace (block + env) automation record.
// It applies to every block type and every env type; a row exists for a workspace once it has
// been read or written, seeded from the block's BlockAutomationDefaults.
//
// What fires an app's auto-deploy is not here: it is the env's GitOps connection config for the app's
// repository (its trigger). Preview envs follow their preview app's pull request or branch.
type WorkspaceAutomation struct {
	OrgName string `json:"orgName"`
	StackId int64  `json:"stackId"`
	BlockId int64  `json:"blockId"`
	EnvId   int64  `json:"envId"`

	IsEnabled  bool `json:"isEnabled"`
	AutoDeploy bool `json:"autoDeploy"`
}
