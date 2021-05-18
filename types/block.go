package types

type Block struct {
	IdModel
	Type                string                     `json:"type"`
	OrgName             string                     `json:"orgName"`
	StackId             int64                      `json:"stackId"`
	StackName           string                     `json:"stackName"`
	Reference           string                     `json:"reference"`
	Name                string                     `json:"name"`
	ModuleSource        string                     `json:"moduleSource"`
	ModuleSourceVersion string                     `json:"moduleSourceVersion"`
	Connections         map[string]BlockConnection `json:"connections"`

	// Deprecated
	ParentBlocks map[string]string `json:"parentBlocks"`
}
