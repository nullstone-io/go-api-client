package types

const (
	RepoInfoVcsProviderGithub = "github"
)

type RepoInfo struct {
	VcsProvider string `json:"vcsProvider"`
	Owner       string `json:"owner"`
	Name        string `json:"name"`
	FullName    string `json:"fullName"`
	Url         string `json:"url"`
}

type CommitInfo struct {
}

type RefInfo struct {
	BranchName  string `json:"branchName"`
	RefType     string `json:"refType"`
	RefSha      string `json:"commitSha"`
	RefUrl      string `json:"commitUrl"`
	VcsUsername string `json:"vcsUsername"`
}
