package types

type AppRunInput struct {
	Id        int64  `json:"id"`
	CommitSha string `json:"commitSha"`
}

type CreateEnvRunInput struct {
	IsDestroy bool           `json:"isDestroy"`
	Apps      *[]AppRunInput `json:"apps"`
}
