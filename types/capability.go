package types

type Capability struct {
	IdModel
	OrgName             string                      `json:"orgName"`
	AppId               int64                       `json:"appId"`
	Name                string                      `json:"name"`
	ModuleSource        string                      `json:"moduleSource"`
	ModuleSourceVersion string                      `json:"moduleSourceVersion"`
	Connections         map[string]ConnectionTarget `json:"connections"`
	Namespace           string                      `json:"namespace"`
}
