package types

type CommitInfo struct {
	RepoOwner   string `json:"repoOwner"`
	RepoName    string `json:"repoName"`
	RepoUrl     string `json:"repoUrl"`
	CommitSha   string `json:"commitSha"`
	CommitUrl   string `json:"commitUrl"`
	VcsUsername string `json:"vcsUsername"`
}
