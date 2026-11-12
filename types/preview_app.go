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
	// TriggerType is what this app follows in its preview env: pull_request, branch, tag or release.
	// The server derives it from BranchName/PullRequestId when either is set.
	TriggerType string `json:"triggerType"`
	// BranchName configures this workspace to sync on push events to this branch
	// BranchName is set only when TriggerType is branch
	BranchName *string `json:"branchName"`
	// PullRequestId configures this workspace to sync on synchronized events to this PR
	// PullRequestId is set only when TriggerType is pull_request
	PullRequestId *int64 `json:"pullRequestId"`
	// TagFilter is a newline-separated list of tag filter patterns; it is only meaningful when TriggerType is tag.
	// An empty filter matches every tag.
	TagFilter string `json:"tagFilter"`

	Config *PreviewPRConfig `json:"config"`
	// Automation is the workspace automation record for this app in the preview env.
	// It is the successor of Config: the flags come from workspace automation (seeded from the
	// legacy config until the backfill) while Config stays until the legacy endpoints are removed.
	Automation  *WorkspaceAutomation `json:"automation,omitempty"`
	PullRequest *PullRequest         `json:"pullRequest"`
}
