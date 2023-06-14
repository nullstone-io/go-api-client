package types

type CreateEnvRunInput struct {
	IsDestroy bool     `json:"isDestroy"`
	AppIds    *[]int64 `json:"appIds"`
	CommitSha *string  `json:"commitSha"`
}
