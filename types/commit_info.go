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
	Repo        string `json:"repo"`
	RepoUrl     string `json:"repoUrl"`
	BranchName  string `json:"branchName"`
	CommitSha   string `json:"commitSha"`
	CommitUrl   string `json:"commitUrl"`
	VcsUsername string `json:"vcsUsername"`
	PRNumber    int    `json:"prNumber"`
	PRId        int64  `json:"prId"`
}
