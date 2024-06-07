package types

const (
	CommitInfoVcsProviderGithub = "github"
)

const (
	CommitInfoTypeBranch = "branch"
	CommitInfoTypePr     = "pr"
	CommitInfoTypePush   = "push"
)

type CommitInfo struct {
	Type        string `json:"type"`
	VcsProvider string `json:"vcsProvider"`
	RepoOwner   string `json:"repoOwner"`
	RepoName    string `json:"repoName"`
	// Repo is `RepoOwner/RepoName`
	Repo       string `json:"repo"`
	RepoUrl    string `json:"repoUrl"`
	BranchName string `json:"branchName"`
	CommitSha  string `json:"commitSha"`
	// CommitUrl is the HTML URL to browse this commit
	CommitUrl     string `json:"commitUrl"`
	CommitMessage string `json:"commitMessage"`
	// CommitUsername is the user in VCS that created the commit
	// This is not guaranteed to be the same as the AuthorUsername
	// When using the GitHub UI to merge, the CommitUsername is actually `web-flow`
	CommitUsername string `json:"commitUsername"`
	// AuthorUsername is the user in VCS that authored the commit
	// This refers to the user that originally made the code changes
	AuthorUsername string `json:"authorUsername"`
	// VcsUsername
	// Deprecated
	VcsUsername string `json:"vcsUsername"`
	PRNumber    int    `json:"prNumber"`
	PRId        int64  `json:"prId"`
	PRTitle     string `json:"prTitle"`
}
