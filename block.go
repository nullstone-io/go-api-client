package api

type Block struct {
	IdModel
	Name                string            `json:"name"`
	OrgName             string            `json:"orgName"`
	StackName           string            `json:"stackName"`
	Layer               string            `json:"layer"`
	ModuleSource        string            `json:"moduleSource"`
	ModuleSourceVersion string            `json:"moduleSourceVersion"`
	ParentBlocks        map[string]string `json:"parentBlocks"`

	Application *Application `json:"application,omitempty"`
}
