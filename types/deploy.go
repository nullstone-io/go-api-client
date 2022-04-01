package types

type Deploy struct {
	IdModel
	OrgName string `json:"orgName"`
	StackId string `json:"stackId"`
	AppId   int64  `json:"appId"`
	EnvId   int64  `json:"envId"`
	Version string `json:"version"`

	App *Application `json:"app,omitempty"`
	Env *Environment `json:"env,omitempty"`
}
