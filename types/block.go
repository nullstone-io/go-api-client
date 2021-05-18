package types

type Block struct {
	IdModel
	Type                string                     `json:"type"`
	OrgName             string                     `json:"orgName"`
	StackId             int64                      `json:"stackId"`
	Reference           string                     `json:"reference"`
	Name                string                     `json:"name"`
	ModuleSource        string                     `json:"moduleSource"`
	ModuleSourceVersion string                     `json:"moduleSourceVersion"`
	ParentBlocks        map[string]string          `json:"parentBlocks"`
	Connections         map[string]BlockConnection `json:"connections"`

	// Deprecated
	StackName           string                     `json:"stackName"`
}
