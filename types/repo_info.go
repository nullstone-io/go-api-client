package types

type RepoInfo struct {
	Repo `json:",inline"`
	
	IsPrivate     bool   `json:"isPrivate"`
	Language      string `json:"language"`
	DefaultBranch string `json:"defaultBranch"`
}
