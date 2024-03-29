package types

type AppEnv struct {
	IdModel
	AppId     int64  `json:"appId"`
	EnvId     int64  `json:"envId"`
	Version   string `json:"version"`
	CommitSha string `json:"commitSha"`

	App *Application `json:"app,omitempty"`
	Env *Environment `json:"env,omitempty"`
}
