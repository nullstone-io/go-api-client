package types

import (
	"fmt"
	"strings"
)

type CapabilityConfig struct {
	Id             int64       `json:"id"`
	Name           string      `json:"name"`
	Source         string      `json:"source"`
	SourceVersion  string      `json:"sourceVersion"`
	Variables      Variables   `json:"variables"`
	Connections    Connections `json:"connections"`
	NeedsDestroyed bool        `json:"needsDestroyed"`
	Namespace      string      `json:"namespace"`
}

func (c CapabilityConfig) TfModuleAddr() string {
	return fmt.Sprintf("module.cap_%d", c.Id)
}

func (c CapabilityConfig) TfModuleName() string {
	return fmt.Sprintf("cap_%d", c.Id)
}

func (c CapabilityConfig) EnvPrefix() string {
	if c.Namespace == "" {
		return ""
	}
	return fmt.Sprintf("%s_", strings.ToUpper(c.Namespace))
}
