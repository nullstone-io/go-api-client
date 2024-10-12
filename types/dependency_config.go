package types

type DependencyConfig struct {
	Reference   WorkspaceTarget `json:"reference"`
	Variables   Variables       `json:"variables"`
	Connections Connections     `json:"connections"`
}
