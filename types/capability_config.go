package types

type CapabilityConfig struct {
	Id             int64       `json:"id"`
	Name           string      `json:"name"`
	Source         string      `json:"source"`
	SourceVersion  string      `json:"sourceVersion"`
	Variables      Variables   `json:"variables"`
	Connections    Connections `json:"connections"`
	NeedsDestroyed bool        `json:"needsDestroyed"`
}
