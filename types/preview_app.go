package types

type PreviewApp struct {
	IdModel
	OrgName    string `json:"orgName"`
	StackId    int64  `json:"stackId"`
	EnvId      int64  `json:"envId"`
	AppId      int64  `json:"appId"`
	ContextKey string `json:"contextKey"`
	Repo       string `json:"repo"`
	RepoUrl    string `json:"repoUrl"`
	// BranchName configures this workspace to sync on push events to this branch
	// BranchName is nil when PullRequestId is not nil
	BranchName *string `json:"branchName"`
	// PullRequestId configures this workspace to sync on synchronized events to this PR
	// PullRequestId is nil when BranchName is not nil
	PullRequestId *int64 `json:"pullRequestId"`

	Config *PreviewPRConfig `json:"config"`
	// Automation is the workspace automation record for this app in the preview env.
	// It is the successor of Config: the flags come from workspace automation (seeded from the
	// legacy config until the backfill) while Config stays until the legacy endpoints are removed.
	Automation  *WorkspaceAutomation `json:"automation,omitempty"`
	PullRequest *PullRequest         `json:"pullRequest"`
}
