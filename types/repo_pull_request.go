package types

type RepoPullRequest struct {
	Id         int64                  `json:"id"`
	Repo       string                 `json:"repo"`
	RepoUrl    string                 `json:"repoUrl"`
	Number     int                    `json:"number"`
	State      string                 `json:"state"`
	Title      string                 `json:"title"`
	Body       *string                `json:"body"`
	HeadBranch string                 `json:"headBranch"`
	HeadSha    string                 `json:"headSha"`
	BaseBranch string                 `json:"baseBranch"`
	BaseSha    string                 `json:"baseSha"`
	Labels     []RepoPullRequestLabel `json:"labels"`
}

type RepoPullRequestLabel struct {
	ID          int64  `json:"id,omitempty"`
	URL         string `json:"url,omitempty"`
	Name        string `json:"name,omitempty"`
	Color       string `json:"color,omitempty"`
	Description string `json:"description,omitempty"`
	Default     bool   `json:"default,omitempty"`
	NodeID      string `json:"node_id,omitempty"`
}
