package types

type ApplicationEnvironment struct {
	OrgName string `json:"orgName"`
	AppName string `json:"appName"`
	EnvName string `json:"envName"`
	Version string `json:"version"`
}
