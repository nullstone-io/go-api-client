package types

type CapabilityConfig struct {
	Id            int64       `json:"id"`
	Source        string      `json:"source" pg:"source"`
	SourceVersion string      `json:"sourceVersion" pg:"source_version"`
	Variables     Variables   `json:"variables" pg:"variables"`
	Connections   Connections `json:"connections" pg:"connections"`
}
