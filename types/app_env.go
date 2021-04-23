package types

type AppEnv struct {
	IdModel
	AppId   int    `json:"appId"`
	EnvId   int    `json:"envId"`
	Version string `json:"version"`

	App *Application `json:"app,omitempty"`
	Env *Environment `json:"env,omitempty"`
}
