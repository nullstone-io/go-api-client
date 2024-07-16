package types

type PullRequest struct {
	IdModel
	Repo       string  `json:"repo"`
	RepoUrl    string  `json:"repoUrl"`
	Number     int     `json:"number"`
	State      string  `json:"state"`
	Title      string  `json:"title"`
	Body       *string `json:"body"`
	HeadBranch string  `json:"headBranch"`
	HeadSha    string  `json:"headSha"`
	BaseBranch string  `json:"baseBranch"`
	BaseSha    string  `json:"baseSha"`
}
