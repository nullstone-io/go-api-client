package types

type GitopsConnection struct {
	IdModel
	OrgName     string `json:"orgName"`
	StackId     int64  `json:"stackId"`
	VcsProvider string `json:"vcsProvider"`
	RepoOwner   string `json:"repoOwner"`
	RepoName    string `json:"repoName"`
	RepoUrl     string `json:"repoUrl"`

	Configs GitopsConnectionConfigs `json:"configs"`
}
