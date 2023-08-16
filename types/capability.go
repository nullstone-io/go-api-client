package types

type Capability struct {
	IdModel
	OrgName             string                      `json:"orgName"`
	AppId               int64                       `json:"appId"`
	EnvId               *int64                      `json:"envId"`
	Name                string                      `json:"name"`
	ModuleSource        string                      `json:"moduleSource,omitempty"`
	ModuleSourceVersion string                      `json:"moduleSourceVersion,omitempty"`
	Connections         map[string]ConnectionTarget `json:"connections,omitempty"`
	Namespace           string                      `json:"namespace"`
	Status              string                      `json:"status,omitempty"`
}
