package types

type EnvironmentType string

const (
	EnvTypePipeline EnvironmentType = "PipelineEnv"
	EnvTypePreview  EnvironmentType = "PreviewEnv"
	EnvTypeGlobal   EnvironmentType = "GlobalEnv"
)

type Environment struct {
	IdModel
	Type           EnvironmentType `json:"type"`
	Name           string          `json:"name"`
	OrgName        string          `json:"orgName"`
	StackId        int64           `json:"stackId"`
	ProviderConfig ProviderConfig  `json:"providerConfig"`
	PipelineOrder  *int            `json:"pipelineOrder,omitempty"`
	ContextKey     string          `json:"contextKey"`
}
