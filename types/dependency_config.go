package types

type DependencyConfig struct {
	Reference   WorkspaceTarget `json:"reference" pg:"reference"`
	Variables   Variables       `json:"variables" pg:"variables"`
	Connections Connections     `json:"connections" pg:"connections"`
}
