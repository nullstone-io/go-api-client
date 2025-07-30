package types

type ModuleConfig struct {
	Source         string      `json:"source"`
	SourceVersion  string      `json:"sourceVersion"`
	SourceToolName string      `json:"sourceToolName"`
	Variables      Variables   `json:"variables"`
	Connections    Connections `json:"connections"`
	Providers      Providers   `json:"providers"`
}
