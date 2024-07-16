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

	Config      *PreviewPRConfig `json:"config"`
	PullRequest *PullRequest     `json:"pullRequest"`
}

func (pa PreviewApp) IsAutoLaunchEnabled() bool {
	return pa.Config != nil &&
		pa.Config.IsEnabled &&
		pa.Config.Status == PreviewPRConfigStatusActive &&
		pa.Config.AutoLaunch
}

func (pa PreviewApp) IsAutoDeployEnabled() bool {
	return pa.Config != nil &&
		pa.Config.IsEnabled &&
		pa.Config.Status == PreviewPRConfigStatusActive &&
		pa.Config.AutoDeploy
}
