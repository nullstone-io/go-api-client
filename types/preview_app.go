package types

// PreviewApp is one app of a preview env and what it follows, in the legacy preview app shape.
//
// Preview apps are no longer stored. The API derives this from what a preview env stores now: an app is listed
// when its WorkspaceAutomation in the env is enabled, and it follows what the env's GitopsConnectionConfig for
// the app's repository follows. It is read-only.
type PreviewApp struct {
	OrgName    string `json:"orgName"`
	StackId    int64  `json:"stackId"`
	EnvId      int64  `json:"envId"`
	AppId      int64  `json:"appId"`
	ContextKey string `json:"contextKey"`
	Repo       string `json:"repo"`
	RepoUrl    string `json:"repoUrl"`
	// TriggerType is what the env follows from the app's repository: pull_request, branch, tag or release.
	// It is empty when the env follows nothing in that repository.
	TriggerType string `json:"triggerType"`
	// BranchName is set only when TriggerType is branch
	BranchName *string `json:"branchName"`
	// PullRequestId is set only when TriggerType is pull_request
	PullRequestId *int64 `json:"pullRequestId"`
	// TagFilter is a newline-separated list of tag filter patterns; it is only meaningful when TriggerType is tag.
	// An empty filter matches every tag.
	TagFilter string `json:"tagFilter"`

	// Config is the app's legacy preview config, nil when it has none
	Config *PreviewPRConfig `json:"config"`
	// Automation is the workspace automation record for this app in the preview env
	Automation  *WorkspaceAutomation `json:"automation,omitempty"`
	PullRequest *PullRequest         `json:"pullRequest"`
}
