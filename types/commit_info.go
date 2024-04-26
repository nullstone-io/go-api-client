package types

const (
	CommitInfoVcsProviderGithub = "github"
)

type CommitInfo struct {
	VcsProvider string `json:"vcsProvider"`
	RepoOwner   string `json:"repoOwner"`
	RepoName    string `json:"repoName"`
	Repo        string `json:"repo"`
	RepoUrl     string `json:"repoUrl"`
	BranchName  string `json:"branchName"`
	CommitSha   string `json:"commitSha"`
	CommitUrl   string `json:"commitUrl"`
	VcsUsername string `json:"vcsUsername"`
}
