package types

type RepoBranch struct {
	Name      string     `json:"name"`
	Protected bool       `json:"protected"`
	IsDefault bool       `json:"isDefault"`
	Commit    CommitInfo `json:"commit"`
}
