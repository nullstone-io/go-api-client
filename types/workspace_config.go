package types

import (
	"github.com/jinzhu/copier"
)

type WorkspaceConfig struct {
	Source        string       `json:"source"`
	SourceVersion string       `json:"sourceVersion"`
	Variables     Variables    `json:"variables"`
	EnvVariables  EnvVariables `json:"envVariables"`
	Connections   Connections  `json:"connections"`
	Providers     Providers    `json:"providers"`
	// Capabilities
	// Deprecated
	Capabilities      CapabilityConfigs      `json:"capabilities"`
	NamedCapabilities NamedCapabilityConfigs `json:"namedCapabilities"`

	// Dependencies represents a list of workspace references that are necessary for this run
	// This is saved to the run config so that a user can quickly access a list of dependencies
	//   It *should not* be used by the nullfire engine to pull in dependencies
	//   because a cleanup run excludes dependencies that are detached, but not destroyed
	Dependencies Dependencies `json:"dependencies"`

	// This field is used to capture user specific configuration for unlaunched dependencies
	// that require some configuration in order to be launched for the first time.
	// Two examples are:
	// - An application that is connected to an unlaunched "Existing Network"
	// - An application that is connected to an unlaunched Datadog datastore via a capability
	DependencyConfigs DependencyConfigs `json:"dependencyConfigs"`
}

func (c WorkspaceConfig) Clone() (WorkspaceConfig, error) {
	config := WorkspaceConfig{}
	err := copier.CopyWithOption(&config, c, copier.Option{DeepCopy: true})
	return config, err
}

func FillWorkspaceConfigMissingEnv(c *WorkspaceConfig, env Environment) {
	envId := env.Id
	fillRef := func(conn Connection) bool {
		if conn.Reference == nil {
			return false
		}
		filled := false
		if conn.Reference.StackId == env.StackId {
			if conn.Reference.EnvId == nil {
				conn.Reference.EnvId = &envId
				filled = true
			}
			if conn.Reference.EnvName == "" {
				conn.Reference.EnvName = env.Name
				filled = true
			}
		}
		return filled
	}
	fillConns := func(conns Connections) bool {
		filled := false
		for name, conn := range conns {
			if fillRef(conn) {
				conns[name] = conn
				filled = true
			}
		}
		return filled
	}

	fillConns(c.Connections)
	for i, capability := range c.Capabilities {
		if fillConns(capability.Connections) {
			c.Capabilities[i] = capability
		}
	}
}
