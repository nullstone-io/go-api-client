package types

type CreateRunInput struct {
	IsDestroy         bool              `json:"isDestroy"`
	IsApproved        *bool             `json:"isApproved"`
	Source            string            `json:"source"`
	SourceVersion     string            `json:"sourceVersion"`
	Variables         Variables         `json:"variables"`
	EnvVariables      EnvVariables      `json:"envVariables"`
	Connections       Connections       `json:"connections"`
	Capabilities      CapabilityConfigs `json:"capabilities"`
	Providers         Providers         `json:"providers"`
	DependencyConfigs DependencyConfigs `json:"dependencyConfigs"`
}
